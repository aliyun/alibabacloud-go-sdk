// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApproveServiceUsageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ApproveServiceUsageRequest
	GetClientToken() *string
	SetComments(v string) *ApproveServiceUsageRequest
	GetComments() *string
	SetRegionId(v string) *ApproveServiceUsageRequest
	GetRegionId() *string
	SetServiceId(v string) *ApproveServiceUsageRequest
	GetServiceId() *string
	SetType(v int32) *ApproveServiceUsageRequest
	GetType() *int32
	SetUserAliUid(v int64) *ApproveServiceUsageRequest
	GetUserAliUid() *int64
}

type ApproveServiceUsageRequest struct {
	// A client token that ensures the idempotence of the request. Generate a unique value from your client for each request. The token can contain only ASCII characters.
	//
	// example:
	//
	// 10CM943JP0EN9D51H
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The approval comments.
	//
	// example:
	//
	// Welcome to ComputeNest service
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-39f4f251e94843xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service sharing type. The default value is SharedAccount. Valid values:
	//
	// - SharedAccount: The regular sharing type.
	//
	// - Reseller: The reseller sharing type.
	//
	// example:
	//
	// SharedAccount
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// The UID of the user\\"s Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 127383705958xxxx
	UserAliUid *int64 `json:"UserAliUid,omitempty" xml:"UserAliUid,omitempty"`
}

func (s ApproveServiceUsageRequest) String() string {
	return dara.Prettify(s)
}

func (s ApproveServiceUsageRequest) GoString() string {
	return s.String()
}

func (s *ApproveServiceUsageRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ApproveServiceUsageRequest) GetComments() *string {
	return s.Comments
}

func (s *ApproveServiceUsageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ApproveServiceUsageRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *ApproveServiceUsageRequest) GetType() *int32 {
	return s.Type
}

func (s *ApproveServiceUsageRequest) GetUserAliUid() *int64 {
	return s.UserAliUid
}

func (s *ApproveServiceUsageRequest) SetClientToken(v string) *ApproveServiceUsageRequest {
	s.ClientToken = &v
	return s
}

func (s *ApproveServiceUsageRequest) SetComments(v string) *ApproveServiceUsageRequest {
	s.Comments = &v
	return s
}

func (s *ApproveServiceUsageRequest) SetRegionId(v string) *ApproveServiceUsageRequest {
	s.RegionId = &v
	return s
}

func (s *ApproveServiceUsageRequest) SetServiceId(v string) *ApproveServiceUsageRequest {
	s.ServiceId = &v
	return s
}

func (s *ApproveServiceUsageRequest) SetType(v int32) *ApproveServiceUsageRequest {
	s.Type = &v
	return s
}

func (s *ApproveServiceUsageRequest) SetUserAliUid(v int64) *ApproveServiceUsageRequest {
	s.UserAliUid = &v
	return s
}

func (s *ApproveServiceUsageRequest) Validate() error {
	return dara.Validate(s)
}
