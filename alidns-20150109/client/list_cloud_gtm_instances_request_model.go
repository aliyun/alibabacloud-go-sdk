// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudGtmInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListCloudGtmInstancesRequest
	GetAcceptLanguage() *string
	SetChargeType(v string) *ListCloudGtmInstancesRequest
	GetChargeType() *string
	SetClientToken(v string) *ListCloudGtmInstancesRequest
	GetClientToken() *string
	SetInstanceId(v string) *ListCloudGtmInstancesRequest
	GetInstanceId() *string
	SetInstanceName(v string) *ListCloudGtmInstancesRequest
	GetInstanceName() *string
	SetPageNumber(v int32) *ListCloudGtmInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCloudGtmInstancesRequest
	GetPageSize() *int32
}

type ListCloudGtmInstancesRequest struct {
	// The language of the response. Valid values:
	//
	// - zh-CN: Chinese.
	//
	// - en-US: English.
	//
	// example:
	//
	// en-US
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// example:
	//
	// postpay/prepay
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// A client token that is used to ensure the idempotence of the request. Generate a unique value on the client. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the GTM instance.
	//
	// example:
	//
	// gtm-cn-jmp3qnw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the instance. This helps you identify the business purpose of the instance.
	//
	// example:
	//
	// test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The page number. The value starts from **1**. The default value is **1**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. The maximum value is **100**. The default value is **20**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListCloudGtmInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListCloudGtmInstancesRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListCloudGtmInstancesRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *ListCloudGtmInstancesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ListCloudGtmInstancesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCloudGtmInstancesRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListCloudGtmInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCloudGtmInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCloudGtmInstancesRequest) SetAcceptLanguage(v string) *ListCloudGtmInstancesRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListCloudGtmInstancesRequest) SetChargeType(v string) *ListCloudGtmInstancesRequest {
	s.ChargeType = &v
	return s
}

func (s *ListCloudGtmInstancesRequest) SetClientToken(v string) *ListCloudGtmInstancesRequest {
	s.ClientToken = &v
	return s
}

func (s *ListCloudGtmInstancesRequest) SetInstanceId(v string) *ListCloudGtmInstancesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListCloudGtmInstancesRequest) SetInstanceName(v string) *ListCloudGtmInstancesRequest {
	s.InstanceName = &v
	return s
}

func (s *ListCloudGtmInstancesRequest) SetPageNumber(v int32) *ListCloudGtmInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCloudGtmInstancesRequest) SetPageSize(v int32) *ListCloudGtmInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListCloudGtmInstancesRequest) Validate() error {
	return dara.Validate(s)
}
