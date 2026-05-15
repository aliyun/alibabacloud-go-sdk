// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendHotlineServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *SuspendHotlineServiceRequest
	GetAccountName() *string
	SetClientToken(v string) *SuspendHotlineServiceRequest
	GetClientToken() *string
	SetInstanceId(v string) *SuspendHotlineServiceRequest
	GetInstanceId() *string
	SetType(v int32) *SuspendHotlineServiceRequest
	GetType() *int32
}

type SuspendHotlineServiceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123@123.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// 46c1341e-2648-447a-9b11-70b6a298d94d
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SuspendHotlineServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s SuspendHotlineServiceRequest) GoString() string {
	return s.String()
}

func (s *SuspendHotlineServiceRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *SuspendHotlineServiceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *SuspendHotlineServiceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SuspendHotlineServiceRequest) GetType() *int32 {
	return s.Type
}

func (s *SuspendHotlineServiceRequest) SetAccountName(v string) *SuspendHotlineServiceRequest {
	s.AccountName = &v
	return s
}

func (s *SuspendHotlineServiceRequest) SetClientToken(v string) *SuspendHotlineServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *SuspendHotlineServiceRequest) SetInstanceId(v string) *SuspendHotlineServiceRequest {
	s.InstanceId = &v
	return s
}

func (s *SuspendHotlineServiceRequest) SetType(v int32) *SuspendHotlineServiceRequest {
	s.Type = &v
	return s
}

func (s *SuspendHotlineServiceRequest) Validate() error {
	return dara.Validate(s)
}
