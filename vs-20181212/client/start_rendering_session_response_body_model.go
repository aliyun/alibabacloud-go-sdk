// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartRenderingSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHostname(v string) *StartRenderingSessionResponseBody
	GetHostname() *string
	SetIsRepeatedRequest(v bool) *StartRenderingSessionResponseBody
	GetIsRepeatedRequest() *bool
	SetLocation(v *StartRenderingSessionResponseBodyLocation) *StartRenderingSessionResponseBody
	GetLocation() *StartRenderingSessionResponseBodyLocation
	SetPortMappings(v []*StartRenderingSessionResponseBodyPortMappings) *StartRenderingSessionResponseBody
	GetPortMappings() []*StartRenderingSessionResponseBodyPortMappings
	SetRenderingInstanceId(v string) *StartRenderingSessionResponseBody
	GetRenderingInstanceId() *string
	SetRequestId(v string) *StartRenderingSessionResponseBody
	GetRequestId() *string
	SetSessionId(v string) *StartRenderingSessionResponseBody
	GetSessionId() *string
	SetStateInfo(v *StartRenderingSessionResponseBodyStateInfo) *StartRenderingSessionResponseBody
	GetStateInfo() *StartRenderingSessionResponseBodyStateInfo
}

type StartRenderingSessionResponseBody struct {
	// The instance hostname. The default value is the access elastic IP address (EIP).
	//
	// example:
	//
	// cn-xxx.ecr.aliyuncs.com
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// Indicates whether the request is a repeated request.
	//
	// example:
	//
	// false
	IsRepeatedRequest *bool `json:"IsRepeatedRequest,omitempty" xml:"IsRepeatedRequest,omitempty"`
	// The location information of the cloud application service instance.
	Location *StartRenderingSessionResponseBodyLocation `json:"Location,omitempty" xml:"Location,omitempty" type:"Struct"`
	// The port mapping information.
	PortMappings []*StartRenderingSessionResponseBodyPortMappings `json:"PortMappings,omitempty" xml:"PortMappings,omitempty" type:"Repeated"`
	// The cloud application service instance ID.
	//
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The session ID.
	//
	// example:
	//
	// session-i205217481741918129226
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The session state information.
	StateInfo *StartRenderingSessionResponseBodyStateInfo `json:"StateInfo,omitempty" xml:"StateInfo,omitempty" type:"Struct"`
}

func (s StartRenderingSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartRenderingSessionResponseBody) GoString() string {
	return s.String()
}

func (s *StartRenderingSessionResponseBody) GetHostname() *string {
	return s.Hostname
}

func (s *StartRenderingSessionResponseBody) GetIsRepeatedRequest() *bool {
	return s.IsRepeatedRequest
}

func (s *StartRenderingSessionResponseBody) GetLocation() *StartRenderingSessionResponseBodyLocation {
	return s.Location
}

func (s *StartRenderingSessionResponseBody) GetPortMappings() []*StartRenderingSessionResponseBodyPortMappings {
	return s.PortMappings
}

func (s *StartRenderingSessionResponseBody) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *StartRenderingSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartRenderingSessionResponseBody) GetSessionId() *string {
	return s.SessionId
}

func (s *StartRenderingSessionResponseBody) GetStateInfo() *StartRenderingSessionResponseBodyStateInfo {
	return s.StateInfo
}

func (s *StartRenderingSessionResponseBody) SetHostname(v string) *StartRenderingSessionResponseBody {
	s.Hostname = &v
	return s
}

func (s *StartRenderingSessionResponseBody) SetIsRepeatedRequest(v bool) *StartRenderingSessionResponseBody {
	s.IsRepeatedRequest = &v
	return s
}

func (s *StartRenderingSessionResponseBody) SetLocation(v *StartRenderingSessionResponseBodyLocation) *StartRenderingSessionResponseBody {
	s.Location = v
	return s
}

func (s *StartRenderingSessionResponseBody) SetPortMappings(v []*StartRenderingSessionResponseBodyPortMappings) *StartRenderingSessionResponseBody {
	s.PortMappings = v
	return s
}

func (s *StartRenderingSessionResponseBody) SetRenderingInstanceId(v string) *StartRenderingSessionResponseBody {
	s.RenderingInstanceId = &v
	return s
}

func (s *StartRenderingSessionResponseBody) SetRequestId(v string) *StartRenderingSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartRenderingSessionResponseBody) SetSessionId(v string) *StartRenderingSessionResponseBody {
	s.SessionId = &v
	return s
}

func (s *StartRenderingSessionResponseBody) SetStateInfo(v *StartRenderingSessionResponseBodyStateInfo) *StartRenderingSessionResponseBody {
	s.StateInfo = v
	return s
}

func (s *StartRenderingSessionResponseBody) Validate() error {
	if s.Location != nil {
		if err := s.Location.Validate(); err != nil {
			return err
		}
	}
	if s.PortMappings != nil {
		for _, item := range s.PortMappings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.StateInfo != nil {
		if err := s.StateInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StartRenderingSessionResponseBodyLocation struct {
	// The province code of the cloud application service instance.
	//
	// example:
	//
	// 610000
	ProvinceCode *string `json:"ProvinceCode,omitempty" xml:"ProvinceCode,omitempty"`
}

func (s StartRenderingSessionResponseBodyLocation) String() string {
	return dara.Prettify(s)
}

func (s StartRenderingSessionResponseBodyLocation) GoString() string {
	return s.String()
}

func (s *StartRenderingSessionResponseBodyLocation) GetProvinceCode() *string {
	return s.ProvinceCode
}

func (s *StartRenderingSessionResponseBodyLocation) SetProvinceCode(v string) *StartRenderingSessionResponseBodyLocation {
	s.ProvinceCode = &v
	return s
}

func (s *StartRenderingSessionResponseBodyLocation) Validate() error {
	return dara.Validate(s)
}

type StartRenderingSessionResponseBodyPortMappings struct {
	// The public port or port range, such as 22. For a port range, separate the start port and end port with a forward slash (/), for example, 10/20.
	//
	// example:
	//
	// 10013/10020
	ExternalPort *string `json:"ExternalPort,omitempty" xml:"ExternalPort,omitempty"`
	// The internal port or port range. The internal port corresponds to the public port on a one-to-one basis. For a port range, separate the start port and end port with a forward slash (/), for example, 10/20.
	//
	// example:
	//
	// 49008/49015
	InternalPort *string `json:"InternalPort,omitempty" xml:"InternalPort,omitempty"`
}

func (s StartRenderingSessionResponseBodyPortMappings) String() string {
	return dara.Prettify(s)
}

func (s StartRenderingSessionResponseBodyPortMappings) GoString() string {
	return s.String()
}

func (s *StartRenderingSessionResponseBodyPortMappings) GetExternalPort() *string {
	return s.ExternalPort
}

func (s *StartRenderingSessionResponseBodyPortMappings) GetInternalPort() *string {
	return s.InternalPort
}

func (s *StartRenderingSessionResponseBodyPortMappings) SetExternalPort(v string) *StartRenderingSessionResponseBodyPortMappings {
	s.ExternalPort = &v
	return s
}

func (s *StartRenderingSessionResponseBodyPortMappings) SetInternalPort(v string) *StartRenderingSessionResponseBodyPortMappings {
	s.InternalPort = &v
	return s
}

func (s *StartRenderingSessionResponseBodyPortMappings) Validate() error {
	return dara.Validate(s)
}

type StartRenderingSessionResponseBodyStateInfo struct {
	// The state description.
	//
	// example:
	//
	// Session starting
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The session state.
	//
	// example:
	//
	// SessionStarting
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The time when the state was last updated. The value is a UTC time in the YYYY-MM-DDTHH:mmZ format based on the ISO 8601 standard.
	//
	// example:
	//
	// 2021-05-06T06:37Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s StartRenderingSessionResponseBodyStateInfo) String() string {
	return dara.Prettify(s)
}

func (s StartRenderingSessionResponseBodyStateInfo) GoString() string {
	return s.String()
}

func (s *StartRenderingSessionResponseBodyStateInfo) GetComment() *string {
	return s.Comment
}

func (s *StartRenderingSessionResponseBodyStateInfo) GetState() *string {
	return s.State
}

func (s *StartRenderingSessionResponseBodyStateInfo) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *StartRenderingSessionResponseBodyStateInfo) SetComment(v string) *StartRenderingSessionResponseBodyStateInfo {
	s.Comment = &v
	return s
}

func (s *StartRenderingSessionResponseBodyStateInfo) SetState(v string) *StartRenderingSessionResponseBodyStateInfo {
	s.State = &v
	return s
}

func (s *StartRenderingSessionResponseBodyStateInfo) SetUpdateTime(v string) *StartRenderingSessionResponseBodyStateInfo {
	s.UpdateTime = &v
	return s
}

func (s *StartRenderingSessionResponseBodyStateInfo) Validate() error {
	return dara.Validate(s)
}
