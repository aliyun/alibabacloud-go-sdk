// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotlineWaitingNumberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *GetHotlineWaitingNumberRequest
	GetAccountName() *string
	SetClientToken(v string) *GetHotlineWaitingNumberRequest
	GetClientToken() *string
	SetInstanceId(v string) *GetHotlineWaitingNumberRequest
	GetInstanceId() *string
}

type GetHotlineWaitingNumberRequest struct {
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
}

func (s GetHotlineWaitingNumberRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHotlineWaitingNumberRequest) GoString() string {
	return s.String()
}

func (s *GetHotlineWaitingNumberRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *GetHotlineWaitingNumberRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetHotlineWaitingNumberRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetHotlineWaitingNumberRequest) SetAccountName(v string) *GetHotlineWaitingNumberRequest {
	s.AccountName = &v
	return s
}

func (s *GetHotlineWaitingNumberRequest) SetClientToken(v string) *GetHotlineWaitingNumberRequest {
	s.ClientToken = &v
	return s
}

func (s *GetHotlineWaitingNumberRequest) SetInstanceId(v string) *GetHotlineWaitingNumberRequest {
	s.InstanceId = &v
	return s
}

func (s *GetHotlineWaitingNumberRequest) Validate() error {
	return dara.Validate(s)
}
