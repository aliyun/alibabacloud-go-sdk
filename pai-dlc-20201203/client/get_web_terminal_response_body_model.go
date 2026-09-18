// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWebTerminalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetWebTerminalResponseBody
	GetRequestId() *string
	SetWebTerminalUrl(v string) *GetWebTerminalResponseBody
	GetWebTerminalUrl() *string
}

type GetWebTerminalResponseBody struct {
	// The request ID for this call, used for diagnostics and troubleshooting.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The WebSocket link for accessing the container. You need to build a WebSocket client. For the detailed communication format, refer to the following code:
	//
	//   ```
	//
	//   ws = new WebSocket(
	//
	//     `wss://xxxxx`,
	//
	//   );
	//
	//   ws.onopen = function open() {
	//
	//     console.warn(\\"connected\\");
	//
	//     term.write(\\"\\");
	//
	//   };
	//
	//   ws.onclose = function close() {
	//
	//     console.warn(\\"disconnected\\");
	//
	//     term.write(\\"Connection closed\\");
	//
	//   };
	//
	//   // Receive response from the backend
	//
	//   ws.onmessage = function incoming(event) {
	//
	//     const msg = JSON.parse(event.data);
	//
	//     console.warn(msg);
	//
	//     if (msg.operation === \\"stdout\\") {
	//
	//       term.write(msg.data);
	//
	//     } else {
	//
	//       console.warn(\\"invalid msg operation: \\" + msg);
	//
	//     }
	//
	//   };
	//
	//   // Console input
	//
	//   term.onData(data => {
	//
	//     const msg = { operation: \\"stdin\\", data: data };
	//
	//     ws.send(JSON.stringify(msg));
	//
	//   });
	//
	//   term.onResize(size => {
	//
	//     const msg = { operation: \\"resize\\", cols: size.cols, rows: size.rows };
	//
	//     ws.send(JSON.stringify(msg));
	//
	//   });
	//
	//   fitAddon.fit();
	//
	// };
	//
	// ```
	//
	// example:
	//
	// wss://*****
	WebTerminalUrl *string `json:"WebTerminalUrl,omitempty" xml:"WebTerminalUrl,omitempty"`
}

func (s GetWebTerminalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWebTerminalResponseBody) GoString() string {
	return s.String()
}

func (s *GetWebTerminalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWebTerminalResponseBody) GetWebTerminalUrl() *string {
	return s.WebTerminalUrl
}

func (s *GetWebTerminalResponseBody) SetRequestId(v string) *GetWebTerminalResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWebTerminalResponseBody) SetWebTerminalUrl(v string) *GetWebTerminalResponseBody {
	s.WebTerminalUrl = &v
	return s
}

func (s *GetWebTerminalResponseBody) Validate() error {
	return dara.Validate(s)
}
