function loadXMLDoc() {
    var xmlhttp;
    xmlhttp=new XMLHttpRequest();
    xmlhttp.onreadystatechange=function()
    {
	if (xmlhttp.readyState==4 && xmlhttp.status==200)
	{
	    document.getElementById("fromserver").innerHTML=xmlhttp.responseText;
	}
    }
    xmlhttp.open("POST","post",true);
    xmlhttp.send();
}