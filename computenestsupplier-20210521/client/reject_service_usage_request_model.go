// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRejectServiceUsageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RejectServiceUsageRequest
	GetClientToken() *string
	SetComments(v string) *RejectServiceUsageRequest
	GetComments() *string
	SetServiceId(v string) *RejectServiceUsageRequest
	GetServiceId() *string
	SetType(v int32) *RejectServiceUsageRequest
	GetType() *int32
	SetUserAliUid(v int64) *RejectServiceUsageRequest
	GetUserAliUid() *int64
}

type RejectServiceUsageRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 10CM943JP0EN9D51H
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Reject comments.
	//
	// example:
	//
	// Thanks for your application, please add your industry information.
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-2117508c874c41xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The share type of the service. Default value: SharedAccount. Valid values:
	//
	// 	- SharedAccount: The service is shared by multiple accounts.
	//
	// 	- Reseller: The service is distributed.
	//
	// example:
	//
	// SharedAccount
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// User ali uid.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1563457855xxxxxx
	UserAliUid *int64 `json:"UserAliUid,omitempty" xml:"UserAliUid,omitempty"`
}

func (s RejectServiceUsageRequest) String() string {
	return dara.Prettify(s)
}

func (s RejectServiceUsageRequest) GoString() string {
	return s.String()
}

func (s *RejectServiceUsageRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RejectServiceUsageRequest) GetComments() *string {
	return s.Comments
}

func (s *RejectServiceUsageRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *RejectServiceUsageRequest) GetType() *int32 {
	return s.Type
}

func (s *RejectServiceUsageRequest) GetUserAliUid() *int64 {
	return s.UserAliUid
}

func (s *RejectServiceUsageRequest) SetClientToken(v string) *RejectServiceUsageRequest {
	s.ClientToken = &v
	return s
}

func (s *RejectServiceUsageRequest) SetComments(v string) *RejectServiceUsageRequest {
	s.Comments = &v
	return s
}

func (s *RejectServiceUsageRequest) SetServiceId(v string) *RejectServiceUsageRequest {
	s.ServiceId = &v
	return s
}

func (s *RejectServiceUsageRequest) SetType(v int32) *RejectServiceUsageRequest {
	s.Type = &v
	return s
}

func (s *RejectServiceUsageRequest) SetUserAliUid(v int64) *RejectServiceUsageRequest {
	s.UserAliUid = &v
	return s
}

func (s *RejectServiceUsageRequest) Validate() error {
	return dara.Validate(s)
}
