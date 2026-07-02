// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDetailByOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptor(v string) *GetDetailByOrderRequest
	GetAcceptor() *string
	SetBizNo(v string) *GetDetailByOrderRequest
	GetBizNo() *string
	SetChannel(v string) *GetDetailByOrderRequest
	GetChannel() *string
	SetCheckAuthItems(v string) *GetDetailByOrderRequest
	GetCheckAuthItems() *string
	SetEmpId(v string) *GetDetailByOrderRequest
	GetEmpId() *string
	SetLanguage(v string) *GetDetailByOrderRequest
	GetLanguage() *string
	SetOptSource(v string) *GetDetailByOrderRequest
	GetOptSource() *string
}

type GetDetailByOrderRequest struct {
	// This parameter is required.
	Acceptor *string `json:"Acceptor,omitempty" xml:"Acceptor,omitempty"`
	// This parameter is required.
	BizNo *string `json:"BizNo,omitempty" xml:"BizNo,omitempty"`
	// This parameter is required.
	Channel        *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	CheckAuthItems *string `json:"CheckAuthItems,omitempty" xml:"CheckAuthItems,omitempty"`
	// This parameter is required.
	EmpId    *string `json:"EmpId,omitempty" xml:"EmpId,omitempty"`
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// This parameter is required.
	OptSource *string `json:"OptSource,omitempty" xml:"OptSource,omitempty"`
}

func (s GetDetailByOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDetailByOrderRequest) GoString() string {
	return s.String()
}

func (s *GetDetailByOrderRequest) GetAcceptor() *string {
	return s.Acceptor
}

func (s *GetDetailByOrderRequest) GetBizNo() *string {
	return s.BizNo
}

func (s *GetDetailByOrderRequest) GetChannel() *string {
	return s.Channel
}

func (s *GetDetailByOrderRequest) GetCheckAuthItems() *string {
	return s.CheckAuthItems
}

func (s *GetDetailByOrderRequest) GetEmpId() *string {
	return s.EmpId
}

func (s *GetDetailByOrderRequest) GetLanguage() *string {
	return s.Language
}

func (s *GetDetailByOrderRequest) GetOptSource() *string {
	return s.OptSource
}

func (s *GetDetailByOrderRequest) SetAcceptor(v string) *GetDetailByOrderRequest {
	s.Acceptor = &v
	return s
}

func (s *GetDetailByOrderRequest) SetBizNo(v string) *GetDetailByOrderRequest {
	s.BizNo = &v
	return s
}

func (s *GetDetailByOrderRequest) SetChannel(v string) *GetDetailByOrderRequest {
	s.Channel = &v
	return s
}

func (s *GetDetailByOrderRequest) SetCheckAuthItems(v string) *GetDetailByOrderRequest {
	s.CheckAuthItems = &v
	return s
}

func (s *GetDetailByOrderRequest) SetEmpId(v string) *GetDetailByOrderRequest {
	s.EmpId = &v
	return s
}

func (s *GetDetailByOrderRequest) SetLanguage(v string) *GetDetailByOrderRequest {
	s.Language = &v
	return s
}

func (s *GetDetailByOrderRequest) SetOptSource(v string) *GetDetailByOrderRequest {
	s.OptSource = &v
	return s
}

func (s *GetDetailByOrderRequest) Validate() error {
	return dara.Validate(s)
}
