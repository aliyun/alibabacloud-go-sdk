// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationAccountsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *ListApplicationAccountsRequest
	GetApplicationId() *string
	SetInstanceId(v string) *ListApplicationAccountsRequest
	GetInstanceId() *string
	SetPageNumber(v int64) *ListApplicationAccountsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListApplicationAccountsRequest
	GetPageSize() *int64
}

type ListApplicationAccountsRequest struct {
	// IDaaS的应用主键id
	//
	// This parameter is required.
	//
	// example:
	//
	// app_11111
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// IDaaS EIAM的实例id
	//
	// This parameter is required.
	//
	// example:
	//
	// eiam-111ccc1111
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 当前查询的列表页码，默认为1
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 当前查询的列表页码，默认为20
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListApplicationAccountsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationAccountsRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationAccountsRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListApplicationAccountsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListApplicationAccountsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListApplicationAccountsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListApplicationAccountsRequest) SetApplicationId(v string) *ListApplicationAccountsRequest {
	s.ApplicationId = &v
	return s
}

func (s *ListApplicationAccountsRequest) SetInstanceId(v string) *ListApplicationAccountsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListApplicationAccountsRequest) SetPageNumber(v int64) *ListApplicationAccountsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListApplicationAccountsRequest) SetPageSize(v int64) *ListApplicationAccountsRequest {
	s.PageSize = &v
	return s
}

func (s *ListApplicationAccountsRequest) Validate() error {
	return dara.Validate(s)
}
