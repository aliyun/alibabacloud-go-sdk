// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchCloudGtmInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *SearchCloudGtmInstancesRequest
	GetAcceptLanguage() *string
	SetChargeType(v string) *SearchCloudGtmInstancesRequest
	GetChargeType() *string
	SetClientToken(v string) *SearchCloudGtmInstancesRequest
	GetClientToken() *string
	SetInstanceId(v string) *SearchCloudGtmInstancesRequest
	GetInstanceId() *string
	SetInstanceName(v string) *SearchCloudGtmInstancesRequest
	GetInstanceName() *string
	SetPageNumber(v int32) *SearchCloudGtmInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *SearchCloudGtmInstancesRequest
	GetPageSize() *int32
}

type SearchCloudGtmInstancesRequest struct {
	// The language of the return value. Options are:
	//
	// - **zh-CN**: Chinese.
	//
	// - **en-US**: English.
	//
	// example:
	//
	// en-US
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// example:
	//
	// prepay / postpay
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see How to ensure idempotence.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24**22
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the Global Traffic Manager (GTM) 3.0 instance.
	//
	// example:
	//
	// gtm-cn-wwo3a3hbz**
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Schedule instance name, supports fuzzy search.
	//
	// example:
	//
	// test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Current page number, starting from 1, default is 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of rows per page when paginating queries, with a maximum value of **100**, and a default of **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s SearchCloudGtmInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchCloudGtmInstancesRequest) GoString() string {
	return s.String()
}

func (s *SearchCloudGtmInstancesRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *SearchCloudGtmInstancesRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *SearchCloudGtmInstancesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *SearchCloudGtmInstancesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SearchCloudGtmInstancesRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *SearchCloudGtmInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchCloudGtmInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchCloudGtmInstancesRequest) SetAcceptLanguage(v string) *SearchCloudGtmInstancesRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *SearchCloudGtmInstancesRequest) SetChargeType(v string) *SearchCloudGtmInstancesRequest {
	s.ChargeType = &v
	return s
}

func (s *SearchCloudGtmInstancesRequest) SetClientToken(v string) *SearchCloudGtmInstancesRequest {
	s.ClientToken = &v
	return s
}

func (s *SearchCloudGtmInstancesRequest) SetInstanceId(v string) *SearchCloudGtmInstancesRequest {
	s.InstanceId = &v
	return s
}

func (s *SearchCloudGtmInstancesRequest) SetInstanceName(v string) *SearchCloudGtmInstancesRequest {
	s.InstanceName = &v
	return s
}

func (s *SearchCloudGtmInstancesRequest) SetPageNumber(v int32) *SearchCloudGtmInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchCloudGtmInstancesRequest) SetPageSize(v int32) *SearchCloudGtmInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *SearchCloudGtmInstancesRequest) Validate() error {
	return dara.Validate(s)
}
