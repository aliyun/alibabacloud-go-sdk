// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRspDomainServerProhibitStatusForGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddStatusList(v []*UpdateRspDomainServerProhibitStatusForGatewayRequestAddStatusList) *UpdateRspDomainServerProhibitStatusForGatewayRequest
	GetAddStatusList() []*UpdateRspDomainServerProhibitStatusForGatewayRequestAddStatusList
	SetClientToken(v string) *UpdateRspDomainServerProhibitStatusForGatewayRequest
	GetClientToken() *string
	SetDeleteStatusList(v []*UpdateRspDomainServerProhibitStatusForGatewayRequestDeleteStatusList) *UpdateRspDomainServerProhibitStatusForGatewayRequest
	GetDeleteStatusList() []*UpdateRspDomainServerProhibitStatusForGatewayRequestDeleteStatusList
	SetDomainName(v string) *UpdateRspDomainServerProhibitStatusForGatewayRequest
	GetDomainName() *string
}

type UpdateRspDomainServerProhibitStatusForGatewayRequest struct {
	AddStatusList []*UpdateRspDomainServerProhibitStatusForGatewayRequestAddStatusList `json:"AddStatusList,omitempty" xml:"AddStatusList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// asdf
	ClientToken      *string                                                                 `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DeleteStatusList []*UpdateRspDomainServerProhibitStatusForGatewayRequestDeleteStatusList `json:"DeleteStatusList,omitempty" xml:"DeleteStatusList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s UpdateRspDomainServerProhibitStatusForGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRspDomainServerProhibitStatusForGatewayRequest) GoString() string {
	return s.String()
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayRequest) GetAddStatusList() []*UpdateRspDomainServerProhibitStatusForGatewayRequestAddStatusList {
	return s.AddStatusList
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayRequest) GetDeleteStatusList() []*UpdateRspDomainServerProhibitStatusForGatewayRequestDeleteStatusList {
	return s.DeleteStatusList
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayRequest) SetAddStatusList(v []*UpdateRspDomainServerProhibitStatusForGatewayRequestAddStatusList) *UpdateRspDomainServerProhibitStatusForGatewayRequest {
	s.AddStatusList = v
	return s
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayRequest) SetClientToken(v string) *UpdateRspDomainServerProhibitStatusForGatewayRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayRequest) SetDeleteStatusList(v []*UpdateRspDomainServerProhibitStatusForGatewayRequestDeleteStatusList) *UpdateRspDomainServerProhibitStatusForGatewayRequest {
	s.DeleteStatusList = v
	return s
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayRequest) SetDomainName(v string) *UpdateRspDomainServerProhibitStatusForGatewayRequest {
	s.DomainName = &v
	return s
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayRequest) Validate() error {
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

type UpdateRspDomainServerProhibitStatusForGatewayRequestAddStatusList struct {
	// example:
	//
	// serverDeleteProhibited
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusMsg *string `json:"StatusMsg,omitempty" xml:"StatusMsg,omitempty"`
}

func (s UpdateRspDomainServerProhibitStatusForGatewayRequestAddStatusList) String() string {
	return dara.Prettify(s)
}

func (s UpdateRspDomainServerProhibitStatusForGatewayRequestAddStatusList) GoString() string {
	return s.String()
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayRequestAddStatusList) GetStatus() *string {
	return s.Status
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayRequestAddStatusList) GetStatusMsg() *string {
	return s.StatusMsg
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayRequestAddStatusList) SetStatus(v string) *UpdateRspDomainServerProhibitStatusForGatewayRequestAddStatusList {
	s.Status = &v
	return s
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayRequestAddStatusList) SetStatusMsg(v string) *UpdateRspDomainServerProhibitStatusForGatewayRequestAddStatusList {
	s.StatusMsg = &v
	return s
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayRequestAddStatusList) Validate() error {
	return dara.Validate(s)
}

type UpdateRspDomainServerProhibitStatusForGatewayRequestDeleteStatusList struct {
	// example:
	//
	// serverDeleteProhibited
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusMsg *string `json:"StatusMsg,omitempty" xml:"StatusMsg,omitempty"`
}

func (s UpdateRspDomainServerProhibitStatusForGatewayRequestDeleteStatusList) String() string {
	return dara.Prettify(s)
}

func (s UpdateRspDomainServerProhibitStatusForGatewayRequestDeleteStatusList) GoString() string {
	return s.String()
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayRequestDeleteStatusList) GetStatus() *string {
	return s.Status
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayRequestDeleteStatusList) GetStatusMsg() *string {
	return s.StatusMsg
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayRequestDeleteStatusList) SetStatus(v string) *UpdateRspDomainServerProhibitStatusForGatewayRequestDeleteStatusList {
	s.Status = &v
	return s
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayRequestDeleteStatusList) SetStatusMsg(v string) *UpdateRspDomainServerProhibitStatusForGatewayRequestDeleteStatusList {
	s.StatusMsg = &v
	return s
}

func (s *UpdateRspDomainServerProhibitStatusForGatewayRequestDeleteStatusList) Validate() error {
	return dara.Validate(s)
}
