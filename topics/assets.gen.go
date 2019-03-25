// Code generated by go-bindata.
// sources:
// assets/templates/args.html
// assets/templates/collection/item.html
// assets/templates/collection/list.html
// assets/templates/collection/main.html
// assets/templates/signals.html
// assets/templates/env.html
// assets/templates/mem.html
// assets/templates/metrics.html
// assets/templates/proc.html
// assets/templates/scopes.html
// assets/templates/version.html
// DO NOT EDIT!

package topics

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)
type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _assetsTemplatesArgsHtml = []byte(`{{ define "content" }}

<p>
    The set of command-line arguments used when starting this process.
</p>

<table>
    <thead>
    <tr>
        <th>Index</th>
        <th>Value</th>
    </tr>
    </thead>

    <tbody>
        {{ range $index, $value := . }}
            <tr>
                <td>{{$index}}</td>
                <td>{{$value}}</td>
            </tr>
        {{ end }}
    </tbody>
</table>

<br>
<button class="btn btn-istio" onclick="configreload()">Force config reload</button>

{{ template "last-refresh" .}}

<script>
    "use strict"

    function configreload() {
        var url = window.location.protocol + "//" + window.location.host + "/argj/reloadconfig";

        var ajax = new XMLHttpRequest();
        ajax.onload = onload;
        ajax.onerror = onerror;
        ajax.open("PUT", url, true);
        ajax.send();

        function onload() {
            console.log(url + " -> " + ajax.status)
        }

        function onerror(e) {
            console.error(e);
        }
    }
</script>

{{ end }}
`)

func assetsTemplatesArgsHtmlBytes() ([]byte, error) {
	return _assetsTemplatesArgsHtml, nil
}

func assetsTemplatesArgsHtml() (*asset, error) {
	bytes, err := assetsTemplatesArgsHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/args.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTemplatesCollectionItemHtml = []byte(`{{ define "content" }}

{{ with $context := . }}
    {{ if ne $context.Error "" }}
        <b>{{$context.Error}}</b>
    {{else}}
        <p> Item {{ $context.Collection }}/{{ $context.Key }}</p>
        <div class="language-yaml highlighter-rouge">
            <div class="highlight">
                <pre class="highlight"><code>{{ $context.Value }}</code></pre>
            </div>
        </div>
    {{end}}
{{end}}
{{ template "last-refresh" .}}

{{ end }}
`)

func assetsTemplatesCollectionItemHtmlBytes() ([]byte, error) {
	return _assetsTemplatesCollectionItemHtml, nil
}

func assetsTemplatesCollectionItemHtml() (*asset, error) {
	bytes, err := assetsTemplatesCollectionItemHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/collection/item.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTemplatesCollectionListHtml = []byte(`{{ define "content" }}

{{ with $context := . }}

    {{ if ne $context.Error "" }}
        <b>{{$context.Error}}</b>
    {{else}}
        <p> Collection {{ $context.Collection }} </p>

        <table>
            <thead>
            <tr>
                <th>Index</th>
                <th>Item</th>
            </tr>
            </thead>

            <tbody>
                 {{ range $index, $key := $context.Keys }}
                <tr>
                    <td>{{$index}}</td>
                    <td><a href="{{$context.Collection}}/{{$key}}">{{$key}}</a></td>
                </tr>
                {{ end }}
            </tbody>
        </table>
    {{end}}

{{end}}

{{ template "last-refresh" .}}

{{ end }}
`)

func assetsTemplatesCollectionListHtmlBytes() ([]byte, error) {
	return _assetsTemplatesCollectionListHtml, nil
}

func assetsTemplatesCollectionListHtml() (*asset, error) {
	bytes, err := assetsTemplatesCollectionListHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/collection/list.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTemplatesCollectionMainHtml = []byte(`{{ define "content" }}

{{ with $context := .}}
    <p>Collections</p>

    <table>
        <tbody>
        {{ range $context.Collections }}
            <tr>
                <td><a href="{{.}}">{{.}}</a></td>
            </tr>
        {{ end }}
        </tbody>
    </table>
{{end}}

{{ template "last-refresh" .}}

{{ end }}
`)

func assetsTemplatesCollectionMainHtmlBytes() ([]byte, error) {
	return _assetsTemplatesCollectionMainHtml, nil
}

func assetsTemplatesCollectionMainHtml() (*asset, error) {
	bytes, err := assetsTemplatesCollectionMainHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/collection/main.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTemplatesCommandsHtml = []byte(`{{ define "content" }}

<p>
    Send commands to the running process.
</p>

<br>
<button class="btn btn-istio" onclick="sendConfigReload()">Reload Config (SIGUSR1)</button>

{{ template "last-refresh" .}}

<script>
    "use strict"

    function sendConfigReload() {
        var url = window.location.protocol + "//" + window.location.host + "/commandj/reloadconfig";

        var ajax = new XMLHttpRequest();
        ajax.onload = onload;
        ajax.onerror = onerror;
        ajax.open("PUT", url, true);
        ajax.send();

        function onload() {
            console.log(url + " -> " + ajax.status)
        }

        function onerror(e) {
            console.error(e);
        }
    }
</script>

{{ end }}
`)

func assetsTemplatesCommandsHtmlBytes() ([]byte, error) {
	return _assetsTemplatesCommandsHtml, nil
}

func assetsTemplatesCommandsHtml() (*asset, error) {
	bytes, err := assetsTemplatesCommandsHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/signals.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTemplatesEnvHtml = []byte(`{{ define "content" }}

<p>
    The set of environment variables defined for this process.
</p>

<table>
    <thead>
    <tr>
        <th>Name</th>
        <th>Value</th>
    </tr>
    </thead>

    <tbody>
        {{ range . }}
            <tr>
                <td>{{.Name}}</td>
                <td>{{.Value}}</td>
            </tr>
        {{ end }}
    </tbody>
</table>

{{ template "last-refresh" .}}

{{ end }}
`)

func assetsTemplatesEnvHtmlBytes() ([]byte, error) {
	return _assetsTemplatesEnvHtml, nil
}

func assetsTemplatesEnvHtml() (*asset, error) {
	bytes, err := assetsTemplatesEnvHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/env.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTemplatesMemHtml = []byte(`{{ define "content" }}

<p>
    This information is gathered from the Go runtime and represents the ongoing memory consumption
    of this process. Please refer to the <a href="https://golang.org/pkg/runtime/#MemStats">Go documentation on the MemStats type</a>
    for more information about all of these values.
</p>

<table>
    <thead>
    <tr>
        <th>Counter</th>
        <th>Value</th>
        <th>Description</th>
    </tr>
    </thead>
    <tbody>
    <tr>
        <td>HeapInuse</td>
        <td id="HeapInuse">{{.HeapInuse}} bytes</td>
        <td>Bytes in in-use spans.</td>
    </tr>

    <tr>
        <td>Total Alloc</td>
        <td id="TotalAlloc">{{.TotalAlloc}} bytes</td>
        <td>Cumulative bytes allocated for heap objects.</td>
    </tr>

    <tr>
        <td>Sys</td>
        <td id="Sys">{{.Sys}} bytes</td>
        <td>Total bytes of memory obtained from the OS.</td>
    </tr>

    <tr>
        <td>Lookups</td>
        <td id="Lookups">{{.Lookups}} lookups</td>
        <td>Number of pointer lookups performed by the runtime.</td>
    </tr>

    <tr>
        <td>Mallocs</td>
        <td id="Mallocs">{{.Mallocs}} objects</td>
        <td>Cumulative count of heap objects allocated.</td>
    </tr>

    <tr>
        <td>Frees</td>
        <td id="Frees">{{.Frees}} objects</td>
        <td>Cumulative count of heap objects freed.</td>
    </tr>

    <tr>
        <td>Live</td>
        <td id="Live">0 objects</td>
        <td>Count of live heap objects.</td>
    </tr>

    <tr>
        <td>HeapAlloc</td>
        <td id="HeapAlloc">{{.HeapAlloc}} bytes</td>
        <td>Allocated heap objects.</td>
    </tr>

    <tr>
        <td>HeapSys</td>
        <td id="HeapSys">{{.HeapSys}} bytes</td>
        <td>Heap memory obtained from the OS.</td>
    </tr>

    <tr>
        <td>HeapIdle</td>
        <td id="HeapIdle">{{.HeapIdle}} bytes</td>
        <td>Bytes in idle (unused) spans.</td>
    </tr>

    <tr>
        <td>HeapReleased</td>
        <td id="HeapReleased">{{.HeapReleased}} bytes</td>
        <td>Physical memory returned to the OS.</td>
    </tr>

    <tr>
        <td>HeapObjects</td>
        <td id="HeapObjects">{{.HeapObjects}} objects</td>
        <td>Number of allocated heap objects.</td>
    </tr>

    <tr>
        <td>StackInuse</td>
        <td id="StackInuse">{{.StackInuse}} bytes</td>
        <td>Bytes in stack spans.</td>
    </tr>

    <tr>
        <td>StackSys</td>
        <td id="StackSys">{{.StackSys}} bytes</td>
        <td>Stack memory obtained from the OS.</td>
    </tr>

    <tr>
        <td>NextGC</td>
        <td id="NextGC">{{.NextGC}} bytes</td>
        <td>Target heap size of the next GC cycle.</td>
    </tr>

    <tr>
        <td>LastGC</td>
        <td id="LastGC"></td>
        <td>The time the last garbage collection finished.</td>
    </tr>

    <script>
        // we do this so there's a useful date in the table, which avoids things shifting around during initial paint
        var d = new Date().toLocaleString();
        document.getElementById("LastGC").innerText = d;
    </script>

    <tr>
        <td>PauseTotalNs</td>
        <td id="PauseTotalNs">{{.PauseTotalNs}} ns</td>
        <td>Cumulative time spent in GC stop-the-world pauses.</td>
    </tr>

    <tr>
        <td>NumGC</td>
        <td id="NumGC">{{.NumGC}} GC cycles</td>
        <td>Completed GC cycles.</td>
    </tr>

    <tr>
        <td>NumForcedGC</td>
        <td id="NumForcedGC">{{.NumForcedGC}} GC cycles</td>
        <td>GC cycles that were forced by the application calling the GC function.</td>
    </tr>

    <tr>
        <td>GCCPUFraction</td>
        <td id="GCCPUFraction"></td>
        <td>Fraction of this program's available CPU time used by the GC since the program started.</td>
    </tr>
    </tbody>
</table>

<br>
<button class="btn btn-istio" onclick="forceCollection()">Force Garbage Collection</button>

{{ template "last-refresh" .}}

<script>
    "use strict"

    function refreshMemStats() {
        var url = window.location.protocol + "//" + window.location.host + "/memj/";

        var ajax = new XMLHttpRequest();
        ajax.onload = onload;
        ajax.onerror = onerror;
        ajax.open("GET", url, true);
        ajax.send();

        function onload() {
            if (this.status == 200) { // request succeeded
                var ms = JSON.parse(this.responseText);
                document.getElementById("HeapInuse").innerText = ms.HeapInuse.toLocaleString() + " bytes";
                document.getElementById("TotalAlloc").innerText = ms.TotalAlloc.toLocaleString() + " bytes";
                document.getElementById("Sys").innerText = ms.Sys.toLocaleString() + " bytes";
                document.getElementById("Lookups").innerText = ms.Lookups.toLocaleString() + " lookups";
                document.getElementById("Mallocs").innerText = ms.Mallocs.toLocaleString() + " objects";
                document.getElementById("Frees").innerText = ms.Frees.toLocaleString() + " objects";
                document.getElementById("Live").innerText = (ms.Mallocs - ms.Frees).toLocaleString() + " objects";
                document.getElementById("HeapAlloc").innerText = ms.HeapAlloc.toLocaleString() + " bytes";
                document.getElementById("HeapSys").innerText = ms.HeapSys.toLocaleString() + " bytes";
                document.getElementById("HeapIdle").innerText = ms.HeapIdle.toLocaleString() + " bytes";
                document.getElementById("HeapReleased").innerText = ms.HeapReleased.toLocaleString() + " bytes";
                document.getElementById("HeapObjects").innerText = ms.HeapObjects.toLocaleString() + " objects";
                document.getElementById("StackInuse").innerText = ms.StackInuse.toLocaleString() + " bytes";
                document.getElementById("StackSys").innerText = ms.StackSys.toLocaleString() + " bytes";
                document.getElementById("NextGC").innerText = ms.NextGC.toLocaleString() + " bytes";
                document.getElementById("PauseTotalNs").innerText = ms.PauseTotalNs.toLocaleString() + " ns";
                document.getElementById("NumGC").innerText = ms.NumGC.toLocaleString() + " GC cycles";
                document.getElementById("NumForcedGC").innerText = ms.NumForcedGC.toLocaleString() + " GC cycles";

                var d = new Date(ms.LastGC / 1000000).toLocaleString();
                document.getElementById("LastGC").innerText = d;

                var frac = ms.GCCPUFraction;
                if (frac < 0) {
                    frac = 0.0;
                }
                var percent = (frac * 100).toFixed(2);
                document.getElementById("GCCPUFraction").innerText = percent + "%";

                updateRefreshTime();
            }
        }

        function onerror(e) {
            console.error(e);
        }
    }

    function forceCollection() {
        var url = window.location.protocol + "//" + window.location.host + "/memj/forcecollection";

        var ajax = new XMLHttpRequest();
        ajax.onload = onload;
        ajax.onerror = onerror;
        ajax.open("PUT", url, true);
        ajax.send();

        function onload() {
        }

        function onerror(e) {
            console.error(e);
        }
    }

    refreshMemStats();
    window.setInterval(refreshMemStats, 1000);

</script>

{{ end }}
`)

func assetsTemplatesMemHtmlBytes() ([]byte, error) {
	return _assetsTemplatesMemHtml, nil
}

func assetsTemplatesMemHtml() (*asset, error) {
	bytes, err := assetsTemplatesMemHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/mem.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTemplatesMetricsHtml = []byte(`{{ define "content" }}

<p>
    The set of metrics published by this process.
</p>

<style>
    .metric-card {
        margin-bottom: 4px;
    }

    .metric-header {
        padding: 0.5rem;
        cursor: pointer;
        background-color: #5f6364;
    }

    .metric-body {
        padding: 0;
    }

    .metric-table {
        margin: 0;
    }
</style>

{{ range . }}
    {{ if .Metrics }}
        <div class="card metric-card">
            <div class="card-header metric-header" onclick="$('#{{.Name}}').collapse('toggle')" role="button" aria-expanded="false" aria-controls="{{.Name}}">
                {{.Name}} [{{.Type}}]</a>
            </div>

            <div id="{{.Name}}" class="collapse">
                <div class="card-body metric-body">
                    <table class="metric-table">
                        {{ if eq .Type "GAUGE" "COUNTER" "UNTYPED" }}
                            <thead>
                                <tr>
                                    <th>Labels</th>
                                    <th>Value</th>
                                </tr>
                            </thead>
                            <tbody>
                                {{ range .Metrics }}
                                    <tr>
                                        <td>
                                            {{ range $k, $v := .Labels }}
                                                {{$k}} : {{$v}}<br>
                                            {{ end }}
                                        </td>
                                        <td>
                                            {{.Value}}
                                        </td>
                                    </tr>
                                {{ end }}
                            </tbody>
                        {{ else }}
                            <thead>
                                <tr>
                                    <th>Labels</th>
                                    <th>Count</th>
                                    <th>Sum</th>
                                </tr>
                            </thead>
                            <tbody>
                                {{ range .Metrics }}
                                    <tr>
                                        <td>
                                            {{ range $k, $v := .Labels }}
                                                {{$k}} : {{$v}}<br>
                                            {{ end }}
                                        </td>
                                        <td>
                                            {{.Count}}
                                        </td>
                                        <td>
                                            {{.Sum}}
                                        </td>
                                    </tr>
                                {{ end }}
                            </tbody>
                        {{ end }}
                    </table>
                </div>
            </div>
        </div>
    {{ end }}
{{ end }}

{{ template "last-refresh" .}}

<script>
    "use strict"

    function refreshMetrics() {
        var url = window.location.protocol + "//" + window.location.host + "/metricj/";

        var ajax = new XMLHttpRequest();
        ajax.onload = onload;
        ajax.onerror = onerror;
        ajax.open("GET", url, true);
        ajax.send();

        function onload() {
            if (this.status == 200) { // request succeeded
                var families = JSON.parse(this.responseText);

                for (var i = 0; i < families.length; i++) {
                    var name = families[i].name;
                    var div = document.getElementById(name);

                    if (div) {
                        var tbody = div.getElementsByTagName("TBODY")[0];
                        for (var j = 0; j < tbody.children.length; j++) {
                            var tr = tbody.children[j];
                            var labels_td = tr.children[0];

                            var labels = "";
                            if (families[i].metrics[j].labels) {
                                for (var key in families[i].metrics[j].labels) {
                                    if (labels.length > 0) {
                                        labels = labels + "<br>";
                                    }

                                    labels = labels + key + " : " + families[i].metrics[j].labels[key];
                                }
                            }

                            labels_td.innerHTML = labels;

                            if (families[i].metrics[j].value != undefined) {
                                var value_td = tr.children[1];
                                value_td.innerText = families[i].metrics[j].value;
                            } else {
                                var count_td = tr.children[1];
                                var sum_td = tr.children[2];
                                count_td.innerText = families[i].metrics[j].count;
                                sum_td.innerText = families[i].metrics[j].sum;
                            }
                        }
                    }
                }


                updateRefreshTime();
            }
        }

        function onerror(e) {
            console.error(e);
        }
    }

    refreshMetrics();
    window.setInterval(refreshMetrics, 1000);

</script>

{{ end }}
`)

func assetsTemplatesMetricsHtmlBytes() ([]byte, error) {
	return _assetsTemplatesMetricsHtml, nil
}

func assetsTemplatesMetricsHtml() (*asset, error) {
	bytes, err := assetsTemplatesMetricsHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/metrics.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTemplatesProcHtml = []byte(`{{ define "content" }}

<p>
    Information about this process.
</p>

<table>
    <thead>
    <tr>
        <th>Name</th>
        <th>Value</th>
    </tr>
    </thead>

    <tbody>
        <tr>
            <td># Threads</td>
            <td id="Threads">{{.Threads}}</td>
        </tr>

        <tr>
            <td># Goroutines</td>
            <td id="Goroutines">{{.Goroutines}}</td>
        </tr>

        <tr>
            <td>Effective Group Id</td>
            <td>{{.Egid}}</td>
        </tr>

        <tr>
            <td>Effective User Id</td>
            <td>{{.Euid}}</td>
        </tr>

        <tr>
            <td>Group Id</td>
            <td>{{.Gid}}</td>
        </tr>

        <tr>
            <td>Groups</td>
            <td>{{.Groups}}</td>
        </tr>

        <tr>
            <td>Hostname</td>
            <td>{{.Hostname}}</td>
        </tr>

        <tr>
            <td>Parent Process Id</td>
            <td>{{.Ppid}}</td>
        </tr>

        <tr>
            <td>Process Id</td>
            <td>{{.Pid}}</td>
        </tr>

        <tr>
            <td>Temporary Directory</td>
            <td>{{.TempDir}}</td>
        </tr>

        <tr>
            <td>User Id</td>
            <td>{{.UID}}</td>
        </tr>

        <tr>
            <td>Working Directory</td>
            <td>{{.Wd}}</td>
        </tr>
    </tbody>
</table>

{{ template "last-refresh" .}}

<script>
    "use strict"

    function refreshProcStats() {
        var url = window.location.protocol + "//" + window.location.host + "/procj/";

        var ajax = new XMLHttpRequest();
        ajax.onload = onload;
        ajax.onerror = onerror;
        ajax.open("GET", url, true);
        ajax.send();

        function onload() {
            if (this.status == 200) { // request succeeded
                var pi = JSON.parse(this.responseText);
                document.getElementById("Threads").innerText = pi.threads;
                document.getElementById("Goroutines").innerText = pi.goroutines;

                updateRefreshTime();
            }
        }

        function onerror(e) {
            console.error(e);
        }
    }

    refreshProcStats();
    window.setInterval(refreshProcStats, 1000);

</script>

{{ end }}
`)

func assetsTemplatesProcHtmlBytes() ([]byte, error) {
	return _assetsTemplatesProcHtml, nil
}

func assetsTemplatesProcHtml() (*asset, error) {
	bytes, err := assetsTemplatesProcHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/proc.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTemplatesScopesHtml = []byte(`{{ define "content" }}
<p>
        Logging for this process is organized in scopes. Each scope has different
        output controls which determine how much and what kind of logging is produced
        by the scope.
</p>

<table>
    <thead>
        <tr>
            <th>Scope</th>
            <th>Description</th>
            <th>Output Level</th>
            <th>Stack Trace Level</th>
            <th>Log Callers?</th>
        </tr>
    </thead>

    <tbody>

        {{ range $index, $value := . }}
            <tr id="{{$value.Name}}">
                <td>{{$value.Name}}</td>
                <td>{{$value.Description}}</td>
                <td class="text-center">
                    <div class="dropdown">
                        <button id="outputLevel" class="btn btn-istio dropdown-toggle" type="button" data-toggle="dropdown">
                            {{$value.OutputLevel}}
                        </button>
                        <div class="dropdown-menu">
                            <a class="dropdown-item" onclick="selectOutputLevel(this, 'none')">none</a>
                            <a class="dropdown-item" onclick="selectOutputLevel(this, 'error')">error</a>
                            <a class="dropdown-item" onclick="selectOutputLevel(this, 'warn')">warn</a>
                            <a class="dropdown-item" onclick="selectOutputLevel(this, 'info')">info</a>
                            <a class="dropdown-item" onclick="selectOutputLevel(this, 'debug')">debug</a>
                        </div>
                    </div>
                </td>

                <td class="text-center">
                    <div class="dropdown">
                        <button id="stackTraceLevel" class="btn btn-istio dropdown-toggle" type="button" data-toggle="dropdown">
                        {{$value.StackTraceLevel}}
                        </button>
                        <div class="dropdown-menu">
                            <a class="dropdown-item" onclick="selectStackTraceLevel(this, 'none')">none</a>
                            <a class="dropdown-item" onclick="selectStackTraceLevel(this, 'error')">error</a>
                            <a class="dropdown-item" onclick="selectStackTraceLevel(this, 'warn')">warn</a>
                            <a class="dropdown-item" onclick="selectStackTraceLevel(this, 'info')">info</a>
                            <a class="dropdown-item" onclick="selectStackTraceLevel(this, 'debug')">debug</a>
                        </div>
                    </div>
                </td>

                <td class="text-center">
                    {{ if $value.LogCallers}}
                        <input id="logCallers" onclick="toggleLogCallers(this)" type="checkbox" checked="checked">
                    {{ else }}
                        <input id="logCallers" onclick="toggleLogCallers(this)" type="checkbox">
                    {{ end }}
                </td>
            </tr>
        {{ end }}

    </tbody>
</table>

{{ template "last-refresh" .}}

<script>
    "use strict"

    function refreshScopes() {
        var url = window.location.protocol + "//" + window.location.host + "/scopej/";

        var ajax = new XMLHttpRequest();
        ajax.onload = onload;
        ajax.onerror = onerror;
        ajax.open("GET", url, true);
        ajax.send();

        function onload() {
            if (this.status == 200) { // request succeeded
                var si = JSON.parse(this.responseText);
                for (var i = 0; i < si.length; i++) {
                    var info = si[i]

                    var tr = document.getElementById(info.name);
                    tr.querySelector("#outputLevel").innerText = info.output_level;
                    tr.querySelector("#stackTraceLevel").innerText = info.stack_trace_level;
                    tr.querySelector("#logCallers").checked = info.log_callers;
                }

                updateRefreshTime();
            }
        }

        function onerror(e) {
            console.error(e);
        }
    }

    function selectOutputLevel(element, level) {
        var scope = element.parentElement.parentElement.parentElement.parentElement.id;

        var url = window.location.protocol + "//" + window.location.host + "/scopej/" + scope;
        var ajax = new XMLHttpRequest();
        ajax.onload = onload;
        ajax.onerror = onerror;
        ajax.open("GET", url, true);
        ajax.send();

        function onload() {
            if (this.status == 200) { // request succeeded
                var si = JSON.parse(this.responseText);
                si.output_level = level;

                var url = window.location.protocol + "//" + window.location.host + "/scopej/" + scope;
                var ajax = new XMLHttpRequest();
                ajax.onload = onload2;
                ajax.onerror = onerror;
                ajax.open("PUT", url, true);
                ajax.send(JSON.stringify(si));
            }

            function onload2() {
                refreshScopes();
            }
        }

        function onerror(e) {
            console.error(e);
        }
    }

    function selectStackTraceLevel(element, level) {
        var scope = element.parentElement.parentElement.parentElement.parentElement.id;

        var url = window.location.protocol + "//" + window.location.host + "/scopej/" + scope;
        var ajax = new XMLHttpRequest();
        ajax.onload = onload;
        ajax.onerror = onerror;
        ajax.open("GET", url, true);
        ajax.send();

        function onload() {
            if (this.status == 200) { // request succeeded
                var si = JSON.parse(this.responseText);
                si.stack_trace_level = level;

                var url = window.location.protocol + "//" + window.location.host + "/scopej/" + scope;
                var ajax = new XMLHttpRequest();
                ajax.onload = onload2;
                ajax.onerror = onerror;
                ajax.open("PUT", url, true);
                ajax.send(JSON.stringify(si));
            }

            function onload2() {
                refreshScopes();
            }
        }

        function onerror(e) {
            console.error(e);
        }
    }

    function toggleLogCallers(checkbox) {
        var scope = checkbox.parentElement.parentElement.id;
        var logCallers = checkbox.checked;

        var url = window.location.protocol + "//" + window.location.host + "/scopej/" + scope;
        var ajax = new XMLHttpRequest();
        ajax.onload = onload;
        ajax.onerror = onerror;
        ajax.open("GET", url, true);
        ajax.send();

        function onload() {
            if (this.status == 200) { // request succeeded
                var si = JSON.parse(this.responseText);
                si.log_callers = logCallers;

                var url = window.location.protocol + "//" + window.location.host + "/scopej/" + scope;
                var ajax = new XMLHttpRequest();
                ajax.onload = onload2;
                ajax.onerror = onerror;
                ajax.open("PUT", url, true);
                ajax.send(JSON.stringify(si));
            }

            function onload2() {
            }
        }

        function onerror(e) {
            console.error(e);
        }
    }

    refreshScopes();
    window.setInterval(refreshScopes, 1000);
</script>

{{ end }}
`)

func assetsTemplatesScopesHtmlBytes() ([]byte, error) {
	return _assetsTemplatesScopesHtml, nil
}

func assetsTemplatesScopesHtml() (*asset, error) {
	bytes, err := assetsTemplatesScopesHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/scopes.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTemplatesVersionHtml = []byte(`{{ define "content" }}

<p>
    Version information about this binary and runtime.
</p>

<table>
    <thead>
    <tr>
        <th>Name</th>
        <th>Value</th>
    </tr>
    </thead>

    <tbody>
        <tr>
            <td>Version</td>
            <td>{{.Version}}</td>
        </tr>

        <tr>
            <td>Git Revision</td>
            <td>{{.GitRevision}}</td>
        </tr>

        <tr>
            <td>User</td>
            <td>{{.User}}</td>
        </tr>

        <tr>
            <td>Host</td>
            <td>{{.Host}}</td>
        </tr>

        <tr>
            <td>GolangVersion</td>
            <td>{{.GolangVersion}}</td>
        </tr>

        <tr>
            <td>DockerHub</td>
            <td>{{.DockerHub}}</td>
        </tr>

        <tr>
            <td>Build Status</td>
            <td>{{.BuildStatus}}</td>
        </tr>
    </tbody>
</table>

{{ template "last-refresh" .}}

{{ end }}
`)

func assetsTemplatesVersionHtmlBytes() ([]byte, error) {
	return _assetsTemplatesVersionHtml, nil
}

func assetsTemplatesVersionHtml() (*asset, error) {
	bytes, err := assetsTemplatesVersionHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/version.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"assets/templates/args.html": assetsTemplatesArgsHtml,
	"assets/templates/collection/item.html": assetsTemplatesCollectionItemHtml,
	"assets/templates/collection/list.html": assetsTemplatesCollectionListHtml,
	"assets/templates/collection/main.html": assetsTemplatesCollectionMainHtml,
	"assets/templates/signals.html": assetsTemplatesCommandsHtml,
	"assets/templates/env.html": assetsTemplatesEnvHtml,
	"assets/templates/mem.html": assetsTemplatesMemHtml,
	"assets/templates/metrics.html": assetsTemplatesMetricsHtml,
	"assets/templates/proc.html": assetsTemplatesProcHtml,
	"assets/templates/scopes.html": assetsTemplatesScopesHtml,
	"assets/templates/version.html": assetsTemplatesVersionHtml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"assets": &bintree{nil, map[string]*bintree{
		"templates": &bintree{nil, map[string]*bintree{
			"args.html": &bintree{assetsTemplatesArgsHtml, map[string]*bintree{}},
			"collection": &bintree{nil, map[string]*bintree{
				"item.html": &bintree{assetsTemplatesCollectionItemHtml, map[string]*bintree{}},
				"list.html": &bintree{assetsTemplatesCollectionListHtml, map[string]*bintree{}},
				"main.html": &bintree{assetsTemplatesCollectionMainHtml, map[string]*bintree{}},
			}},
			"signals.html": &bintree{assetsTemplatesCommandsHtml, map[string]*bintree{}},
			"env.html": &bintree{assetsTemplatesEnvHtml, map[string]*bintree{}},
			"mem.html": &bintree{assetsTemplatesMemHtml, map[string]*bintree{}},
			"metrics.html": &bintree{assetsTemplatesMetricsHtml, map[string]*bintree{}},
			"proc.html": &bintree{assetsTemplatesProcHtml, map[string]*bintree{}},
			"scopes.html": &bintree{assetsTemplatesScopesHtml, map[string]*bintree{}},
			"version.html": &bintree{assetsTemplatesVersionHtml, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

