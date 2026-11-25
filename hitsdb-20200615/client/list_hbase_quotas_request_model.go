// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHBaseQuotasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListHBaseQuotasRequest
	GetInstanceId() *string
	SetRegionId(v string) *ListHBaseQuotasRequest
	GetRegionId() *string
	SetSecurityToken(v string) *ListHBaseQuotasRequest
	GetSecurityToken() *string
}

type ListHBaseQuotasRequest struct {
	// This parameter is required.
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ListHBaseQuotasRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHBaseQuotasRequest) GoString() string {
	return s.String()
}

func (s *ListHBaseQuotasRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListHBaseQuotasRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListHBaseQuotasRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ListHBaseQuotasRequest) SetInstanceId(v string) *ListHBaseQuotasRequest {
	s.InstanceId = &v
	return s
}

func (s *ListHBaseQuotasRequest) SetRegionId(v string) *ListHBaseQuotasRequest {
	s.RegionId = &v
	return s
}

func (s *ListHBaseQuotasRequest) SetSecurityToken(v string) *ListHBaseQuotasRequest {
	s.SecurityToken = &v
	return s
}

func (s *ListHBaseQuotasRequest) Validate() error {
	return dara.Validate(s)
}
