// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRetcodeShareUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPid(v string) *GetRetcodeShareUrlRequest
	GetPid() *string
}

type GetRetcodeShareUrlRequest struct {
	// The project ID (PID) of the application.
	//
	// To obtain the application PID, log on to the Application Real-Time Monitoring Service (ARMS) console. In the left-side navigation pane, choose **Browser Monitoring**>**Browser Monitoring**. Then, click the name of the application. The URL in the address bar contains the application PID, in the xxx format. As the browser is encoded, you must modify the PID. For example, if the PID in the URL is xxx%4074xxx, you must replace %40 with an at sign (@) to obtain xxx@74xxx.
	//
	// This parameter is required.
	//
	// example:
	//
	// iioe7jcnuk@582846f37******
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
}

func (s GetRetcodeShareUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRetcodeShareUrlRequest) GoString() string {
	return s.String()
}

func (s *GetRetcodeShareUrlRequest) GetPid() *string {
	return s.Pid
}

func (s *GetRetcodeShareUrlRequest) SetPid(v string) *GetRetcodeShareUrlRequest {
	s.Pid = &v
	return s
}

func (s *GetRetcodeShareUrlRequest) Validate() error {
	return dara.Validate(s)
}
