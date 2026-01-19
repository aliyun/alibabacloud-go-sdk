// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClientPublicKeysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *ListClientPublicKeysRequest
	GetApplicationId() *string
	SetInstanceId(v string) *ListClientPublicKeysRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListClientPublicKeysRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListClientPublicKeysRequest
	GetNextToken() *string
}

type ListClientPublicKeysRequest struct {
	// IDaaS的应用资源ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 查询凭证（Token），取值为上一次API调用返回的NextToken参数值。
	//
	// example:
	//
	// NTxxxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListClientPublicKeysRequest) String() string {
	return dara.Prettify(s)
}

func (s ListClientPublicKeysRequest) GoString() string {
	return s.String()
}

func (s *ListClientPublicKeysRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListClientPublicKeysRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListClientPublicKeysRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListClientPublicKeysRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListClientPublicKeysRequest) SetApplicationId(v string) *ListClientPublicKeysRequest {
	s.ApplicationId = &v
	return s
}

func (s *ListClientPublicKeysRequest) SetInstanceId(v string) *ListClientPublicKeysRequest {
	s.InstanceId = &v
	return s
}

func (s *ListClientPublicKeysRequest) SetMaxResults(v int32) *ListClientPublicKeysRequest {
	s.MaxResults = &v
	return s
}

func (s *ListClientPublicKeysRequest) SetNextToken(v string) *ListClientPublicKeysRequest {
	s.NextToken = &v
	return s
}

func (s *ListClientPublicKeysRequest) Validate() error {
	return dara.Validate(s)
}
