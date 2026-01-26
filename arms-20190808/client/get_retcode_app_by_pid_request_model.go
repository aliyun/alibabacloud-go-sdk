// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRetcodeAppByPidRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPid(v string) *GetRetcodeAppByPidRequest
	GetPid() *string
	SetRegionId(v string) *GetRetcodeAppByPidRequest
	GetRegionId() *string
	SetTags(v []*GetRetcodeAppByPidRequestTags) *GetRetcodeAppByPidRequest
	GetTags() []*GetRetcodeAppByPidRequestTags
}

type GetRetcodeAppByPidRequest struct {
	// The PID of the application. To obtain the PID of the application, perform the following steps: Log on to the Application Real-Time Monitoring Service (ARMS) console. In the left-side navigation pane, choose **Browser Monitoring*	- > **Browser Monitoring**. On the Browser Monitoring page, click the name of the application. The URL in the address bar contains the PID of the application. The PID is in the pid=xxx format. The PID is usually percent encoded as xxx%40xxx. You must modify this value to remove the percent encoding. For example, if the PID in the URL is xxx%4074xxx, you must replace %40 with the at sign (@) to obtain xxx@74xxx.
	//
	// This parameter is required.
	//
	// example:
	//
	// b590lhguqs@9781be0f44dXXXX
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of tags.
	Tags []*GetRetcodeAppByPidRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s GetRetcodeAppByPidRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRetcodeAppByPidRequest) GoString() string {
	return s.String()
}

func (s *GetRetcodeAppByPidRequest) GetPid() *string {
	return s.Pid
}

func (s *GetRetcodeAppByPidRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetRetcodeAppByPidRequest) GetTags() []*GetRetcodeAppByPidRequestTags {
	return s.Tags
}

func (s *GetRetcodeAppByPidRequest) SetPid(v string) *GetRetcodeAppByPidRequest {
	s.Pid = &v
	return s
}

func (s *GetRetcodeAppByPidRequest) SetRegionId(v string) *GetRetcodeAppByPidRequest {
	s.RegionId = &v
	return s
}

func (s *GetRetcodeAppByPidRequest) SetTags(v []*GetRetcodeAppByPidRequestTags) *GetRetcodeAppByPidRequest {
	s.Tags = v
	return s
}

func (s *GetRetcodeAppByPidRequest) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetRetcodeAppByPidRequestTags struct {
	// The tag key.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetRetcodeAppByPidRequestTags) String() string {
	return dara.Prettify(s)
}

func (s GetRetcodeAppByPidRequestTags) GoString() string {
	return s.String()
}

func (s *GetRetcodeAppByPidRequestTags) GetKey() *string {
	return s.Key
}

func (s *GetRetcodeAppByPidRequestTags) GetValue() *string {
	return s.Value
}

func (s *GetRetcodeAppByPidRequestTags) SetKey(v string) *GetRetcodeAppByPidRequestTags {
	s.Key = &v
	return s
}

func (s *GetRetcodeAppByPidRequestTags) SetValue(v string) *GetRetcodeAppByPidRequestTags {
	s.Value = &v
	return s
}

func (s *GetRetcodeAppByPidRequestTags) Validate() error {
	return dara.Validate(s)
}
