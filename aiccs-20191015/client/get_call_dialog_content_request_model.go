// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCallDialogContentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallDate(v string) *GetCallDialogContentRequest
	GetCallDate() *string
	SetCallId(v string) *GetCallDialogContentRequest
	GetCallId() *string
	SetOwnerId(v int64) *GetCallDialogContentRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetCallDialogContentRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetCallDialogContentRequest
	GetResourceOwnerId() *int64
}

type GetCallDialogContentRequest struct {
	// The outbound call date, in yyyy-MM-dd format. You can only query data from the last 30 days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2025-06-06
	CallDate *string `json:"CallDate,omitempty" xml:"CallDate,omitempty"`
	// The unique identifier for a call.
	//
	// > The LlmSmartCall API returns this ID in the CallId parameter of its response.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456^123478
	CallId               *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetCallDialogContentRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCallDialogContentRequest) GoString() string {
	return s.String()
}

func (s *GetCallDialogContentRequest) GetCallDate() *string {
	return s.CallDate
}

func (s *GetCallDialogContentRequest) GetCallId() *string {
	return s.CallId
}

func (s *GetCallDialogContentRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetCallDialogContentRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetCallDialogContentRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetCallDialogContentRequest) SetCallDate(v string) *GetCallDialogContentRequest {
	s.CallDate = &v
	return s
}

func (s *GetCallDialogContentRequest) SetCallId(v string) *GetCallDialogContentRequest {
	s.CallId = &v
	return s
}

func (s *GetCallDialogContentRequest) SetOwnerId(v int64) *GetCallDialogContentRequest {
	s.OwnerId = &v
	return s
}

func (s *GetCallDialogContentRequest) SetResourceOwnerAccount(v string) *GetCallDialogContentRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetCallDialogContentRequest) SetResourceOwnerId(v int64) *GetCallDialogContentRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetCallDialogContentRequest) Validate() error {
	return dara.Validate(s)
}
