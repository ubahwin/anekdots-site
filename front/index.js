function sendRequest(method, url, body = null) {
    const headers = {
        'Content-Type': 'application/json',
    };

    return fetch(url, {
        method: method,
        body: JSON.stringify(body),
        headers: headers
    }).then(async (response) => ({
        status: response.status,
        data: await response.json(),
    }));
}

function newlineToBr(str) {
    return str.replace(/(?:\r\n|\r|\\n)/g, '<br>');
}

sendRequest('POST', "http://localhost:9999/api/anekdot")
    .then(data => {
        console.log(data);
        document.querySelector(".category-input").textContent = data.data.category_name;
        document.querySelector(".id-input").textContent = "№" + data.data.id;
        document.querySelector(".anekdot-input").innerHTML = newlineToBr(data.data.text);
    });

