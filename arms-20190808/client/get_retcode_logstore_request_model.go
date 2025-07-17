// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRetcodeLogstoreRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPid(v string) *GetRetcodeLogstoreRequest
	GetPid() *string
	SetRegionId(v string) *GetRetcodeLogstoreRequest
	GetRegionId() *string
}

type GetRetcodeLogstoreRequest struct {
	// The process identifier (PID) of the application. To obtain the PID of the application, perform the following steps: Log on to the Application Real-Time Monitoring Service (ARMS) console. In the left-side navigation pane, choose **Browser Monitoring*	- > **Browser Monitoring**. On the Browser Monitoring page, click the name of the application. The URL in the address bar contains the PID of the application. The PID is in the pid=xxx format. The PID is usually percent encoded as xxx%40xxx. You must modify this value to remove the percent encoding. For example, if the PID in the URL is xxx%4074xxx, you must replace %40 with the at sign (@) to obtain xxx@74xxx.
	//
	// This parameter is required.
	//
	// example:
	//
	// atc889zkcf@d8deedfa9bf****
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetRetcodeLogstoreRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRetcodeLogstoreRequest) GoString() string {
	return s.String()
}

func (s *GetRetcodeLogstoreRequest) GetPid() *string {
	return s.Pid
}

func (s *GetRetcodeLogstoreRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetRetcodeLogstoreRequest) SetPid(v string) *GetRetcodeLogstoreRequest {
	s.Pid = &v
	return s
}

func (s *GetRetcodeLogstoreRequest) SetRegionId(v string) *GetRetcodeLogstoreRequest {
	s.RegionId = &v
	return s
}

func (s *GetRetcodeLogstoreRequest) Validate() error {
	return dara.Validate(s)
}
