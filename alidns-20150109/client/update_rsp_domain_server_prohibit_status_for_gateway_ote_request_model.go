// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRspDomainServerProhibitStatusForGatewayOteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddStatusList(v []*UpdateRspDomainServerProhibitStatusForGatewayOteRequestAddStatusList) *UpdateRspDomainServerProhibitStatusForGatewayOteRequest
	GetAddStatusList() []*UpdateRspDomainServerProhibitStatusForGatewayOteRequestAddStatusList
	SetClientToken(v string) *UpdateRspDomainServerProhibitStatusForGatewayOteRequest
	GetClientToken() *string
	SetDeleteStatusList(v []*UpdateRspDomainServerProhibitStatusForGatewayOteRequestDeleteStatusList) *UpdateRspDomainServerProhibitStatusForGatewayOteRequest
	GetDeleteStatusList() []*UpdateRspDomainServerProhibitStatusForGatewayOteRequestDeleteStatusList
	SetDomainName(v string) *UpdateRspDomainServerProhibitStatusForGatewayOteRequest
	GetDomainName() *string
}

type UpdateRspDomainServerProhibitStatusForGatewayOteRequest struct {
	AddStatusList []*UpdateRspDomainServerProhibitStatusForGatewayOteRequestAddStatusList `json:"AddStatusList,omitempty" xml:"AddStatusList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// qwoefasdf
	ClientToken      *string                                                                    `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DeleteStatusList []*UpdateRspDomainServerProhibitStatusForGatewayOteRequestDeleteStatusList `json:"DeleteStatusList,omitempty" xml:"DeleteStatusList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s UpdateRspDomainServerProhibitStatusForGatewayOteRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRspDomainServerProhibitStatusForGatewayOteRequest) GoString() string {
	return s.String()
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteRequest) GetAddStatusList() []*UpdateRspDomainServerProhibitStatusForGatewayOteRequestAddStatusList {
	return s.AddStatusList
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteRequest) GetDeleteStatusList() []*UpdateRspDomainServerProhibitStatusForGatewayOteRequestDeleteStatusList {
	return s.DeleteStatusList
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteRequest) SetAddStatusList(v []*UpdateRspDomainServerProhibitStatusForGatewayOteRequestAddStatusList) *UpdateRspDomainServerProhibitStatusForGatewayOteRequest {
	s.AddStatusList = v
	return s
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteRequest) SetClientToken(v string) *UpdateRspDomainServerProhibitStatusForGatewayOteRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteRequest) SetDeleteStatusList(v []*UpdateRspDomainServerProhibitStatusForGatewayOteRequestDeleteStatusList) *UpdateRspDomainServerProhibitStatusForGatewayOteRequest {
	s.DeleteStatusList = v
	return s
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteRequest) SetDomainName(v string) *UpdateRspDomainServerProhibitStatusForGatewayOteRequest {
	s.DomainName = &v
	return s
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteRequest) Validate() error {
	if s.AddStatusList != nil {
		for _, item := range s.AddStatusList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DeleteStatusList != nil {
		for _, item := range s.DeleteStatusList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateRspDomainServerProhibitStatusForGatewayOteRequestAddStatusList struct {
	// example:
	//
	// serverDeleteProhibited
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusMsg *string `json:"StatusMsg,omitempty" xml:"StatusMsg,omitempty"`
}

func (s UpdateRspDomainServerProhibitStatusForGatewayOteRequestAddStatusList) String() string {
	return dara.Prettify(s)
}

func (s UpdateRspDomainServerProhibitStatusForGatewayOteRequestAddStatusList) GoString() string {
	return s.String()
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteRequestAddStatusList) GetStatus() *string {
	return s.Status
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteRequestAddStatusList) GetStatusMsg() *string {
	return s.StatusMsg
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteRequestAddStatusList) SetStatus(v string) *UpdateRspDomainServerProhibitStatusForGatewayOteRequestAddStatusList {
	s.Status = &v
	return s
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteRequestAddStatusList) SetStatusMsg(v string) *UpdateRspDomainServerProhibitStatusForGatewayOteRequestAddStatusList {
	s.StatusMsg = &v
	return s
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteRequestAddStatusList) Validate() error {
	return dara.Validate(s)
}

type UpdateRspDomainServerProhibitStatusForGatewayOteRequestDeleteStatusList struct {
	// example:
	//
	// serverDeleteProhibited
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusMsg *string `json:"StatusMsg,omitempty" xml:"StatusMsg,omitempty"`
}

func (s UpdateRspDomainServerProhibitStatusForGatewayOteRequestDeleteStatusList) String() string {
	return dara.Prettify(s)
}

func (s UpdateRspDomainServerProhibitStatusForGatewayOteRequestDeleteStatusList) GoString() string {
	return s.String()
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteRequestDeleteStatusList) GetStatus() *string {
	return s.Status
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteRequestDeleteStatusList) GetStatusMsg() *string {
	return s.StatusMsg
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteRequestDeleteStatusList) SetStatus(v string) *UpdateRspDomainServerProhibitStatusForGatewayOteRequestDeleteStatusList {
	s.Status = &v
	return s
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteRequestDeleteStatusList) SetStatusMsg(v string) *UpdateRspDomainServerProhibitStatusForGatewayOteRequestDeleteStatusList {
	s.StatusMsg = &v
	return s
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayOteRequestDeleteStatusList) Validate() error {
	return dara.Validate(s)
}
