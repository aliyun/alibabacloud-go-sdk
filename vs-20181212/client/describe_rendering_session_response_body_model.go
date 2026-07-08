// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRenderingSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAdditionalIngresses(v []*DescribeRenderingSessionResponseBodyAdditionalIngresses) *DescribeRenderingSessionResponseBody
	GetAdditionalIngresses() []*DescribeRenderingSessionResponseBodyAdditionalIngresses
	SetAppId(v string) *DescribeRenderingSessionResponseBody
	GetAppId() *string
	SetClientId(v string) *DescribeRenderingSessionResponseBody
	GetClientId() *string
	SetHostname(v string) *DescribeRenderingSessionResponseBody
	GetHostname() *string
	SetIsp(v string) *DescribeRenderingSessionResponseBody
	GetIsp() *string
	SetLocation(v *DescribeRenderingSessionResponseBodyLocation) *DescribeRenderingSessionResponseBody
	GetLocation() *DescribeRenderingSessionResponseBodyLocation
	SetPatchId(v string) *DescribeRenderingSessionResponseBody
	GetPatchId() *string
	SetPortMappings(v []*DescribeRenderingSessionResponseBodyPortMappings) *DescribeRenderingSessionResponseBody
	GetPortMappings() []*DescribeRenderingSessionResponseBodyPortMappings
	SetRenderingInstanceId(v string) *DescribeRenderingSessionResponseBody
	GetRenderingInstanceId() *string
	SetRequestId(v string) *DescribeRenderingSessionResponseBody
	GetRequestId() *string
	SetSessionId(v string) *DescribeRenderingSessionResponseBody
	GetSessionId() *string
	SetStartTime(v string) *DescribeRenderingSessionResponseBody
	GetStartTime() *string
	SetStateInfo(v *DescribeRenderingSessionResponseBodyStateInfo) *DescribeRenderingSessionResponseBody
	GetStateInfo() *DescribeRenderingSessionResponseBodyStateInfo
}

type DescribeRenderingSessionResponseBody struct {
	// Additional optional ingress network information
	AdditionalIngresses []*DescribeRenderingSessionResponseBodyAdditionalIngresses `json:"AdditionalIngresses,omitempty" xml:"AdditionalIngresses,omitempty" type:"Repeated"`
	// Cloud application ID
	//
	// example:
	//
	// cap-b06b26edfhytbn b94a75ae1a79efc90eb
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// End client ID
	//
	// example:
	//
	// c91263a0-f9ac-45bd-bbe9-6e293ad32d91
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// Instance hostname. Defaults to the EIP address.
	//
	// example:
	//
	// 111.45.29.96
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// Carrier code. Valid values:
	//
	// 1. cmcc
	//
	// 2. unicom
	//
	// 3. telecom
	//
	// example:
	//
	// telecom
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// Cloud application service instance location
	Location *DescribeRenderingSessionResponseBodyLocation `json:"Location,omitempty" xml:"Location,omitempty" type:"Struct"`
	// Cloud application patch package ID. An empty value means the original version.
	//
	// example:
	//
	// patch-03fa76e8e13a49b6a966b063d9d309b4
	PatchId *string `json:"PatchId,omitempty" xml:"PatchId,omitempty"`
	// Port mapping information
	PortMappings []*DescribeRenderingSessionResponseBodyPortMappings `json:"PortMappings,omitempty" xml:"PortMappings,omitempty" type:"Repeated"`
	// Cloud application service instance ID
	//
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
	// Request ID
	//
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Session ID
	//
	// example:
	//
	// session-i205217481741918129226
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// Start time
	//
	// example:
	//
	// 2025-05-18T02:20:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Session state information
	StateInfo *DescribeRenderingSessionResponseBodyStateInfo `json:"StateInfo,omitempty" xml:"StateInfo,omitempty" type:"Struct"`
}

func (s DescribeRenderingSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenderingSessionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRenderingSessionResponseBody) GetAdditionalIngresses() []*DescribeRenderingSessionResponseBodyAdditionalIngresses {
	return s.AdditionalIngresses
}

func (s *DescribeRenderingSessionResponseBody) GetAppId() *string {
	return s.AppId
}

func (s *DescribeRenderingSessionResponseBody) GetClientId() *string {
	return s.ClientId
}

func (s *DescribeRenderingSessionResponseBody) GetHostname() *string {
	return s.Hostname
}

func (s *DescribeRenderingSessionResponseBody) GetIsp() *string {
	return s.Isp
}

func (s *DescribeRenderingSessionResponseBody) GetLocation() *DescribeRenderingSessionResponseBodyLocation {
	return s.Location
}

func (s *DescribeRenderingSessionResponseBody) GetPatchId() *string {
	return s.PatchId
}

func (s *DescribeRenderingSessionResponseBody) GetPortMappings() []*DescribeRenderingSessionResponseBodyPortMappings {
	return s.PortMappings
}

func (s *DescribeRenderingSessionResponseBody) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *DescribeRenderingSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRenderingSessionResponseBody) GetSessionId() *string {
	return s.SessionId
}

func (s *DescribeRenderingSessionResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeRenderingSessionResponseBody) GetStateInfo() *DescribeRenderingSessionResponseBodyStateInfo {
	return s.StateInfo
}

func (s *DescribeRenderingSessionResponseBody) SetAdditionalIngresses(v []*DescribeRenderingSessionResponseBodyAdditionalIngresses) *DescribeRenderingSessionResponseBody {
	s.AdditionalIngresses = v
	return s
}

func (s *DescribeRenderingSessionResponseBody) SetAppId(v string) *DescribeRenderingSessionResponseBody {
	s.AppId = &v
	return s
}

func (s *DescribeRenderingSessionResponseBody) SetClientId(v string) *DescribeRenderingSessionResponseBody {
	s.ClientId = &v
	return s
}

func (s *DescribeRenderingSessionResponseBody) SetHostname(v string) *DescribeRenderingSessionResponseBody {
	s.Hostname = &v
	return s
}

func (s *DescribeRenderingSessionResponseBody) SetIsp(v string) *DescribeRenderingSessionResponseBody {
	s.Isp = &v
	return s
}

func (s *DescribeRenderingSessionResponseBody) SetLocation(v *DescribeRenderingSessionResponseBodyLocation) *DescribeRenderingSessionResponseBody {
	s.Location = v
	return s
}

func (s *DescribeRenderingSessionResponseBody) SetPatchId(v string) *DescribeRenderingSessionResponseBody {
	s.PatchId = &v
	return s
}

func (s *DescribeRenderingSessionResponseBody) SetPortMappings(v []*DescribeRenderingSessionResponseBodyPortMappings) *DescribeRenderingSessionResponseBody {
	s.PortMappings = v
	return s
}

func (s *DescribeRenderingSessionResponseBody) SetRenderingInstanceId(v string) *DescribeRenderingSessionResponseBody {
	s.RenderingInstanceId = &v
	return s
}

func (s *DescribeRenderingSessionResponseBody) SetRequestId(v string) *DescribeRenderingSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRenderingSessionResponseBody) SetSessionId(v string) *DescribeRenderingSessionResponseBody {
	s.SessionId = &v
	return s
}

func (s *DescribeRenderingSessionResponseBody) SetStartTime(v string) *DescribeRenderingSessionResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeRenderingSessionResponseBody) SetStateInfo(v *DescribeRenderingSessionResponseBodyStateInfo) *DescribeRenderingSessionResponseBody {
	s.StateInfo = v
	return s
}

func (s *DescribeRenderingSessionResponseBody) Validate() error {
	if s.AdditionalIngresses != nil {
		for _, item := range s.AdditionalIngresses {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type DescribeRenderingSessionResponseBodyAdditionalIngresses struct {
	// Domain name or IP address of the cloud application service instance
	//
	// example:
	//
	// 111.45.29.96
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// Carrier code. Valid values:
	//
	// 1. cmcc
	//
	// 2. unicom
	//
	// 3. telecom
	//
	// example:
	//
	// unicom
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// List of port mappings
	PortMappings []*DescribeRenderingSessionResponseBodyAdditionalIngressesPortMappings `json:"PortMappings,omitempty" xml:"PortMappings,omitempty" type:"Repeated"`
}

func (s DescribeRenderingSessionResponseBodyAdditionalIngresses) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenderingSessionResponseBodyAdditionalIngresses) GoString() string {
	return s.String()
}

func (s *DescribeRenderingSessionResponseBodyAdditionalIngresses) GetHostname() *string {
	return s.Hostname
}

func (s *DescribeRenderingSessionResponseBodyAdditionalIngresses) GetIsp() *string {
	return s.Isp
}

func (s *DescribeRenderingSessionResponseBodyAdditionalIngresses) GetPortMappings() []*DescribeRenderingSessionResponseBodyAdditionalIngressesPortMappings {
	return s.PortMappings
}

func (s *DescribeRenderingSessionResponseBodyAdditionalIngresses) SetHostname(v string) *DescribeRenderingSessionResponseBodyAdditionalIngresses {
	s.Hostname = &v
	return s
}

func (s *DescribeRenderingSessionResponseBodyAdditionalIngresses) SetIsp(v string) *DescribeRenderingSessionResponseBodyAdditionalIngresses {
	s.Isp = &v
	return s
}

func (s *DescribeRenderingSessionResponseBodyAdditionalIngresses) SetPortMappings(v []*DescribeRenderingSessionResponseBodyAdditionalIngressesPortMappings) *DescribeRenderingSessionResponseBodyAdditionalIngresses {
	s.PortMappings = v
	return s
}

func (s *DescribeRenderingSessionResponseBodyAdditionalIngresses) Validate() error {
	if s.PortMappings != nil {
		for _, item := range s.PortMappings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRenderingSessionResponseBodyAdditionalIngressesPortMappings struct {
	// Public port or port range, such as 22. For a port range, use a forward slash (/) to separate the start and end ports. Example: 10/20.
	//
	// example:
	//
	// 11060/11079
	ExternalPort *string `json:"ExternalPort,omitempty" xml:"ExternalPort,omitempty"`
	// Private port or port range. Each private port maps one-to-one with a public port. For a port range, use a forward slash (/) to separate the start and end ports. Example: 10/20.
	//
	// example:
	//
	// 11060/11079
	InternalPort *string `json:"InternalPort,omitempty" xml:"InternalPort,omitempty"`
}

func (s DescribeRenderingSessionResponseBodyAdditionalIngressesPortMappings) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenderingSessionResponseBodyAdditionalIngressesPortMappings) GoString() string {
	return s.String()
}

func (s *DescribeRenderingSessionResponseBodyAdditionalIngressesPortMappings) GetExternalPort() *string {
	return s.ExternalPort
}

func (s *DescribeRenderingSessionResponseBodyAdditionalIngressesPortMappings) GetInternalPort() *string {
	return s.InternalPort
}

func (s *DescribeRenderingSessionResponseBodyAdditionalIngressesPortMappings) SetExternalPort(v string) *DescribeRenderingSessionResponseBodyAdditionalIngressesPortMappings {
	s.ExternalPort = &v
	return s
}

func (s *DescribeRenderingSessionResponseBodyAdditionalIngressesPortMappings) SetInternalPort(v string) *DescribeRenderingSessionResponseBodyAdditionalIngressesPortMappings {
	s.InternalPort = &v
	return s
}

func (s *DescribeRenderingSessionResponseBodyAdditionalIngressesPortMappings) Validate() error {
	return dara.Validate(s)
}

type DescribeRenderingSessionResponseBodyLocation struct {
	// Province code of the cloud application service instance
	//
	// example:
	//
	// 310000
	ProvinceCode *string `json:"ProvinceCode,omitempty" xml:"ProvinceCode,omitempty"`
}

func (s DescribeRenderingSessionResponseBodyLocation) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenderingSessionResponseBodyLocation) GoString() string {
	return s.String()
}

func (s *DescribeRenderingSessionResponseBodyLocation) GetProvinceCode() *string {
	return s.ProvinceCode
}

func (s *DescribeRenderingSessionResponseBodyLocation) SetProvinceCode(v string) *DescribeRenderingSessionResponseBodyLocation {
	s.ProvinceCode = &v
	return s
}

func (s *DescribeRenderingSessionResponseBodyLocation) Validate() error {
	return dara.Validate(s)
}

type DescribeRenderingSessionResponseBodyPortMappings struct {
	// Public port or port range, such as 22. For a port range, use a forward slash (/) to separate the start and end ports. Example: 10/20.
	//
	// example:
	//
	// 10013/10020
	ExternalPort *string `json:"ExternalPort,omitempty" xml:"ExternalPort,omitempty"`
	// Private port or port range. Each private port maps one-to-one with a public port. For a port range, use a forward slash (/) to separate the start and end ports. Example: 10/20.
	//
	// example:
	//
	// 49008/49015
	InternalPort *string `json:"InternalPort,omitempty" xml:"InternalPort,omitempty"`
}

func (s DescribeRenderingSessionResponseBodyPortMappings) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenderingSessionResponseBodyPortMappings) GoString() string {
	return s.String()
}

func (s *DescribeRenderingSessionResponseBodyPortMappings) GetExternalPort() *string {
	return s.ExternalPort
}

func (s *DescribeRenderingSessionResponseBodyPortMappings) GetInternalPort() *string {
	return s.InternalPort
}

func (s *DescribeRenderingSessionResponseBodyPortMappings) SetExternalPort(v string) *DescribeRenderingSessionResponseBodyPortMappings {
	s.ExternalPort = &v
	return s
}

func (s *DescribeRenderingSessionResponseBodyPortMappings) SetInternalPort(v string) *DescribeRenderingSessionResponseBodyPortMappings {
	s.InternalPort = &v
	return s
}

func (s *DescribeRenderingSessionResponseBodyPortMappings) Validate() error {
	return dara.Validate(s)
}

type DescribeRenderingSessionResponseBodyStateInfo struct {
	// State description
	//
	// example:
	//
	// 已启动
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// Session state. Valid values:
	//
	// 1. SessionStarting: Starting the session
	//
	// 2. SessionStartSuspended: Session start is suspended. Retry by calling Start again.
	//
	// 3. SessionStarted: Session started or in use
	//
	// 4. SessionStartFailed: Session failed to start
	//
	// 5. SessionAbnormal: Session became abnormal after starting successfully
	//
	// 6. SessionStopping: Stopping the session
	//
	// 7. SessionStopFailed: Session failed to stop
	//
	// example:
	//
	// SessionStarted
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// Time when the state was last updated
	//
	// example:
	//
	// 2024-10-15T10:05:20+08:00
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeRenderingSessionResponseBodyStateInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenderingSessionResponseBodyStateInfo) GoString() string {
	return s.String()
}

func (s *DescribeRenderingSessionResponseBodyStateInfo) GetComment() *string {
	return s.Comment
}

func (s *DescribeRenderingSessionResponseBodyStateInfo) GetState() *string {
	return s.State
}

func (s *DescribeRenderingSessionResponseBodyStateInfo) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeRenderingSessionResponseBodyStateInfo) SetComment(v string) *DescribeRenderingSessionResponseBodyStateInfo {
	s.Comment = &v
	return s
}

func (s *DescribeRenderingSessionResponseBodyStateInfo) SetState(v string) *DescribeRenderingSessionResponseBodyStateInfo {
	s.State = &v
	return s
}

func (s *DescribeRenderingSessionResponseBodyStateInfo) SetUpdateTime(v string) *DescribeRenderingSessionResponseBodyStateInfo {
	s.UpdateTime = &v
	return s
}

func (s *DescribeRenderingSessionResponseBodyStateInfo) Validate() error {
	return dara.Validate(s)
}
