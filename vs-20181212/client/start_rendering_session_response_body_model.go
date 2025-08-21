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
	// example:
	//
	// cn-xxx.ecr.aliyuncs.com
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// example:
	//
	// false
	IsRepeatedRequest   *bool                                            `json:"IsRepeatedRequest,omitempty" xml:"IsRepeatedRequest,omitempty"`
	Location            *StartRenderingSessionResponseBodyLocation       `json:"Location,omitempty" xml:"Location,omitempty" type:"Struct"`
	PortMappings        []*StartRenderingSessionResponseBodyPortMappings `json:"PortMappings,omitempty" xml:"PortMappings,omitempty" type:"Repeated"`
	RenderingInstanceId *string                                          `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// session-i205217481741918129226
	SessionId *string                                     `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
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
	return dara.Validate(s)
}

type StartRenderingSessionResponseBodyLocation struct {
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
	// example:
	//
	// 10013/10020
	ExternalPort *string `json:"ExternalPort,omitempty" xml:"ExternalPort,omitempty"`
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
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// SessionStarting
	State *string `json:"State,omitempty" xml:"State,omitempty"`
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
