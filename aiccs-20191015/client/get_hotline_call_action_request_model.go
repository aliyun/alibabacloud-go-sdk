// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotlineCallActionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcc(v string) *GetHotlineCallActionRequest
	GetAcc() *string
	SetAccountName(v string) *GetHotlineCallActionRequest
	GetAccountName() *string
	SetAct(v int32) *GetHotlineCallActionRequest
	GetAct() *int32
	SetBiz(v string) *GetHotlineCallActionRequest
	GetBiz() *string
	SetClientToken(v string) *GetHotlineCallActionRequest
	GetClientToken() *string
	SetFromSource(v string) *GetHotlineCallActionRequest
	GetFromSource() *string
	SetInstanceId(v string) *GetHotlineCallActionRequest
	GetInstanceId() *string
}

type GetHotlineCallActionRequest struct {
	// example:
	//
	// {"time":1}
	Acc *string `json:"Acc,omitempty" xml:"Acc,omitempty"`
	// example:
	//
	// username@example.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// 1
	Act *int32 `json:"Act,omitempty" xml:"Act,omitempty"`
	// example:
	//
	// {"name":123}
	Biz *string `json:"Biz,omitempty" xml:"Biz,omitempty"`
	// example:
	//
	// 46c1341e-2648-447a-9b11-70b6a298d9****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// hotlinebs_out
	FromSource *string `json:"FromSource,omitempty" xml:"FromSource,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetHotlineCallActionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHotlineCallActionRequest) GoString() string {
	return s.String()
}

func (s *GetHotlineCallActionRequest) GetAcc() *string {
	return s.Acc
}

func (s *GetHotlineCallActionRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *GetHotlineCallActionRequest) GetAct() *int32 {
	return s.Act
}

func (s *GetHotlineCallActionRequest) GetBiz() *string {
	return s.Biz
}

func (s *GetHotlineCallActionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetHotlineCallActionRequest) GetFromSource() *string {
	return s.FromSource
}

func (s *GetHotlineCallActionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetHotlineCallActionRequest) SetAcc(v string) *GetHotlineCallActionRequest {
	s.Acc = &v
	return s
}

func (s *GetHotlineCallActionRequest) SetAccountName(v string) *GetHotlineCallActionRequest {
	s.AccountName = &v
	return s
}

func (s *GetHotlineCallActionRequest) SetAct(v int32) *GetHotlineCallActionRequest {
	s.Act = &v
	return s
}

func (s *GetHotlineCallActionRequest) SetBiz(v string) *GetHotlineCallActionRequest {
	s.Biz = &v
	return s
}

func (s *GetHotlineCallActionRequest) SetClientToken(v string) *GetHotlineCallActionRequest {
	s.ClientToken = &v
	return s
}

func (s *GetHotlineCallActionRequest) SetFromSource(v string) *GetHotlineCallActionRequest {
	s.FromSource = &v
	return s
}

func (s *GetHotlineCallActionRequest) SetInstanceId(v string) *GetHotlineCallActionRequest {
	s.InstanceId = &v
	return s
}

func (s *GetHotlineCallActionRequest) Validate() error {
	return dara.Validate(s)
}
