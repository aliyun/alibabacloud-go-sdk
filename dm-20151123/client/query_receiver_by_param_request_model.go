// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryReceiverByParamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyWord(v string) *QueryReceiverByParamRequest
	GetKeyWord() *string
	SetOwnerId(v int64) *QueryReceiverByParamRequest
	GetOwnerId() *int64
	SetPageNo(v int32) *QueryReceiverByParamRequest
	GetPageNo() *int32
	SetPageSize(v int32) *QueryReceiverByParamRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *QueryReceiverByParamRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryReceiverByParamRequest
	GetResourceOwnerId() *int64
	SetStatus(v int32) *QueryReceiverByParamRequest
	GetStatus() *int32
}

type QueryReceiverByParamRequest struct {
	// Keyword, defaults to all information if not specified
	//
	// example:
	//
	// mesh-notification
	KeyWord *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Current page number
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// Number of items per page, default: 10
	//
	// example:
	//
	// 10
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Delivery result. If not filled, it represents all statuses. Values:
	//
	// - 0: Success
	//
	// - 2: Invalid address
	//
	// - 3: Spam
	//
	// - 4: Failure
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryReceiverByParamRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryReceiverByParamRequest) GoString() string {
	return s.String()
}

func (s *QueryReceiverByParamRequest) GetKeyWord() *string {
	return s.KeyWord
}

func (s *QueryReceiverByParamRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryReceiverByParamRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *QueryReceiverByParamRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryReceiverByParamRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryReceiverByParamRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryReceiverByParamRequest) GetStatus() *int32 {
	return s.Status
}

func (s *QueryReceiverByParamRequest) SetKeyWord(v string) *QueryReceiverByParamRequest {
	s.KeyWord = &v
	return s
}

func (s *QueryReceiverByParamRequest) SetOwnerId(v int64) *QueryReceiverByParamRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryReceiverByParamRequest) SetPageNo(v int32) *QueryReceiverByParamRequest {
	s.PageNo = &v
	return s
}

func (s *QueryReceiverByParamRequest) SetPageSize(v int32) *QueryReceiverByParamRequest {
	s.PageSize = &v
	return s
}

func (s *QueryReceiverByParamRequest) SetResourceOwnerAccount(v string) *QueryReceiverByParamRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryReceiverByParamRequest) SetResourceOwnerId(v int64) *QueryReceiverByParamRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryReceiverByParamRequest) SetStatus(v int32) *QueryReceiverByParamRequest {
	s.Status = &v
	return s
}

func (s *QueryReceiverByParamRequest) Validate() error {
	return dara.Validate(s)
}
