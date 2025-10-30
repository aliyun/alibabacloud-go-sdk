// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRecordFileDownloadUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallId(v string) *QueryRecordFileDownloadUrlRequest
	GetCallId() *string
	SetCallTime(v string) *QueryRecordFileDownloadUrlRequest
	GetCallTime() *string
	SetOwnerId(v int64) *QueryRecordFileDownloadUrlRequest
	GetOwnerId() *int64
	SetPoolKey(v string) *QueryRecordFileDownloadUrlRequest
	GetPoolKey() *string
	SetProductType(v string) *QueryRecordFileDownloadUrlRequest
	GetProductType() *string
	SetResourceOwnerAccount(v string) *QueryRecordFileDownloadUrlRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryRecordFileDownloadUrlRequest
	GetResourceOwnerId() *int64
}

type QueryRecordFileDownloadUrlRequest struct {
	// The ID of the call record. Log on to the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account) and view **Call Record ID*	- on the **Call Record Query*	- page.
	//
	// This parameter is required.
	//
	// example:
	//
	// abcedf1234
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// The call initiation time in the call record. Log on to the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account). View **Call Initiated At*	- on the **Call Record Query*	- page, or view the call_time field in the Call Detail Record (CDR) receipt.
	//
	// example:
	//
	// 2019-03-05 12:00:00
	CallTime *string `json:"CallTime,omitempty" xml:"CallTime,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The key of the phone number pool. Log on to the [Phone Number Protection console](https://dypls.console.aliyun.com/dypls.htm#/account) and view the key of the phone number pool on the **Number Pool Management*	- page.
	//
	// >  This parameter is required when **ProductType*	- is left empty.
	//
	// example:
	//
	// FC123456
	PoolKey *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	// The product type. Valid values:
	//
	// 	- **AXB_170**.
	//
	// 	- **AXN_170**.
	//
	// 	- **AXN_95**.
	//
	// 	- **AXN_EXTENSION_REUSE**
	//
	// >
	//
	// 	- This parameter is applicable to the original key accounts of Alibaba Cloud. This parameter can be ignored for Alibaba Cloud users.
	//
	// 	- This parameter is required when **PoolKey*	- is left empty.
	//
	// example:
	//
	// AXB_170
	ProductType          *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryRecordFileDownloadUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryRecordFileDownloadUrlRequest) GoString() string {
	return s.String()
}

func (s *QueryRecordFileDownloadUrlRequest) GetCallId() *string {
	return s.CallId
}

func (s *QueryRecordFileDownloadUrlRequest) GetCallTime() *string {
	return s.CallTime
}

func (s *QueryRecordFileDownloadUrlRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryRecordFileDownloadUrlRequest) GetPoolKey() *string {
	return s.PoolKey
}

func (s *QueryRecordFileDownloadUrlRequest) GetProductType() *string {
	return s.ProductType
}

func (s *QueryRecordFileDownloadUrlRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryRecordFileDownloadUrlRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryRecordFileDownloadUrlRequest) SetCallId(v string) *QueryRecordFileDownloadUrlRequest {
	s.CallId = &v
	return s
}

func (s *QueryRecordFileDownloadUrlRequest) SetCallTime(v string) *QueryRecordFileDownloadUrlRequest {
	s.CallTime = &v
	return s
}

func (s *QueryRecordFileDownloadUrlRequest) SetOwnerId(v int64) *QueryRecordFileDownloadUrlRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryRecordFileDownloadUrlRequest) SetPoolKey(v string) *QueryRecordFileDownloadUrlRequest {
	s.PoolKey = &v
	return s
}

func (s *QueryRecordFileDownloadUrlRequest) SetProductType(v string) *QueryRecordFileDownloadUrlRequest {
	s.ProductType = &v
	return s
}

func (s *QueryRecordFileDownloadUrlRequest) SetResourceOwnerAccount(v string) *QueryRecordFileDownloadUrlRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryRecordFileDownloadUrlRequest) SetResourceOwnerId(v int64) *QueryRecordFileDownloadUrlRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryRecordFileDownloadUrlRequest) Validate() error {
	return dara.Validate(s)
}
