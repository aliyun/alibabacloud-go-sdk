// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRspDomainStatusOteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddStatusList(v []*UpdateRspDomainStatusOteRequestAddStatusList) *UpdateRspDomainStatusOteRequest
	GetAddStatusList() []*UpdateRspDomainStatusOteRequestAddStatusList
	SetClientToken(v string) *UpdateRspDomainStatusOteRequest
	GetClientToken() *string
	SetDeleteStatusList(v []*UpdateRspDomainStatusOteRequestDeleteStatusList) *UpdateRspDomainStatusOteRequest
	GetDeleteStatusList() []*UpdateRspDomainStatusOteRequestDeleteStatusList
	SetDomainName(v string) *UpdateRspDomainStatusOteRequest
	GetDomainName() *string
	SetOperatorId(v string) *UpdateRspDomainStatusOteRequest
	GetOperatorId() *string
	SetOperatorType(v string) *UpdateRspDomainStatusOteRequest
	GetOperatorType() *string
}

type UpdateRspDomainStatusOteRequest struct {
	// example:
	//
	// [{"Status":"renewProhibited","StatusMsg":"test"}]
	AddStatusList []*UpdateRspDomainStatusOteRequestAddStatusList `json:"AddStatusList,omitempty" xml:"AddStatusList,omitempty" type:"Repeated"`
	// example:
	//
	// 443F1A21-XXXX-55C4-93E1-FF020DF93D7B
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// [{"Status":"renewProhibited","StatusMsg":"test"}]
	DeleteStatusList []*UpdateRspDomainStatusOteRequestDeleteStatusList `json:"DeleteStatusList,omitempty" xml:"DeleteStatusList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// gatewayId001
	OperatorId *string `json:"OperatorId,omitempty" xml:"OperatorId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// registryGateway
	OperatorType *string `json:"OperatorType,omitempty" xml:"OperatorType,omitempty"`
}

func (s UpdateRspDomainStatusOteRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRspDomainStatusOteRequest) GoString() string {
	return s.String()
}

func (s *UpdateRspDomainStatusOteRequest) GetAddStatusList() []*UpdateRspDomainStatusOteRequestAddStatusList {
	return s.AddStatusList
}

func (s *UpdateRspDomainStatusOteRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateRspDomainStatusOteRequest) GetDeleteStatusList() []*UpdateRspDomainStatusOteRequestDeleteStatusList {
	return s.DeleteStatusList
}

func (s *UpdateRspDomainStatusOteRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *UpdateRspDomainStatusOteRequest) GetOperatorId() *string {
	return s.OperatorId
}

func (s *UpdateRspDomainStatusOteRequest) GetOperatorType() *string {
	return s.OperatorType
}

func (s *UpdateRspDomainStatusOteRequest) SetAddStatusList(v []*UpdateRspDomainStatusOteRequestAddStatusList) *UpdateRspDomainStatusOteRequest {
	s.AddStatusList = v
	return s
}

func (s *UpdateRspDomainStatusOteRequest) SetClientToken(v string) *UpdateRspDomainStatusOteRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateRspDomainStatusOteRequest) SetDeleteStatusList(v []*UpdateRspDomainStatusOteRequestDeleteStatusList) *UpdateRspDomainStatusOteRequest {
	s.DeleteStatusList = v
	return s
}

func (s *UpdateRspDomainStatusOteRequest) SetDomainName(v string) *UpdateRspDomainStatusOteRequest {
	s.DomainName = &v
	return s
}

func (s *UpdateRspDomainStatusOteRequest) SetOperatorId(v string) *UpdateRspDomainStatusOteRequest {
	s.OperatorId = &v
	return s
}

func (s *UpdateRspDomainStatusOteRequest) SetOperatorType(v string) *UpdateRspDomainStatusOteRequest {
	s.OperatorType = &v
	return s
}

func (s *UpdateRspDomainStatusOteRequest) Validate() error {
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

type UpdateRspDomainStatusOteRequestAddStatusList struct {
	// example:
	//
	// serverDeleteProhibited
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusMsg *string `json:"StatusMsg,omitempty" xml:"StatusMsg,omitempty"`
}

func (s UpdateRspDomainStatusOteRequestAddStatusList) String() string {
	return dara.Prettify(s)
}

func (s UpdateRspDomainStatusOteRequestAddStatusList) GoString() string {
	return s.String()
}

func (s *UpdateRspDomainStatusOteRequestAddStatusList) GetStatus() *string {
	return s.Status
}

func (s *UpdateRspDomainStatusOteRequestAddStatusList) GetStatusMsg() *string {
	return s.StatusMsg
}

func (s *UpdateRspDomainStatusOteRequestAddStatusList) SetStatus(v string) *UpdateRspDomainStatusOteRequestAddStatusList {
	s.Status = &v
	return s
}

func (s *UpdateRspDomainStatusOteRequestAddStatusList) SetStatusMsg(v string) *UpdateRspDomainStatusOteRequestAddStatusList {
	s.StatusMsg = &v
	return s
}

func (s *UpdateRspDomainStatusOteRequestAddStatusList) Validate() error {
	return dara.Validate(s)
}

type UpdateRspDomainStatusOteRequestDeleteStatusList struct {
	// example:
	//
	// serverDeleteProhibited
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusMsg *string `json:"StatusMsg,omitempty" xml:"StatusMsg,omitempty"`
}

func (s UpdateRspDomainStatusOteRequestDeleteStatusList) String() string {
	return dara.Prettify(s)
}

func (s UpdateRspDomainStatusOteRequestDeleteStatusList) GoString() string {
	return s.String()
}

func (s *UpdateRspDomainStatusOteRequestDeleteStatusList) GetStatus() *string {
	return s.Status
}

func (s *UpdateRspDomainStatusOteRequestDeleteStatusList) GetStatusMsg() *string {
	return s.StatusMsg
}

func (s *UpdateRspDomainStatusOteRequestDeleteStatusList) SetStatus(v string) *UpdateRspDomainStatusOteRequestDeleteStatusList {
	s.Status = &v
	return s
}

func (s *UpdateRspDomainStatusOteRequestDeleteStatusList) SetStatusMsg(v string) *UpdateRspDomainStatusOteRequestDeleteStatusList {
	s.StatusMsg = &v
	return s
}

func (s *UpdateRspDomainStatusOteRequestDeleteStatusList) Validate() error {
	return dara.Validate(s)
}
