// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotlineTransferRegisterFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotlineNumber(v string) *ListHotlineTransferRegisterFileRequest
	GetHotlineNumber() *string
	SetOwnerId(v int64) *ListHotlineTransferRegisterFileRequest
	GetOwnerId() *int64
	SetPageNo(v int32) *ListHotlineTransferRegisterFileRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListHotlineTransferRegisterFileRequest
	GetPageSize() *int32
	SetQualificationId(v string) *ListHotlineTransferRegisterFileRequest
	GetQualificationId() *string
	SetResourceOwnerAccount(v string) *ListHotlineTransferRegisterFileRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListHotlineTransferRegisterFileRequest
	GetResourceOwnerId() *int64
}

type ListHotlineTransferRegisterFileRequest struct {
	// The China 400 number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 400****
	HotlineNumber *string `json:"HotlineNumber,omitempty" xml:"HotlineNumber,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Valid values: 1 to 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The qualification ID. You can call the [GetHotlineQualificationByOrder](https://help.aliyun.com/document_detail/393548.html) operation to obtain the qualification ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000****
	QualificationId      *string `json:"QualificationId,omitempty" xml:"QualificationId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListHotlineTransferRegisterFileRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHotlineTransferRegisterFileRequest) GoString() string {
	return s.String()
}

func (s *ListHotlineTransferRegisterFileRequest) GetHotlineNumber() *string {
	return s.HotlineNumber
}

func (s *ListHotlineTransferRegisterFileRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListHotlineTransferRegisterFileRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListHotlineTransferRegisterFileRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHotlineTransferRegisterFileRequest) GetQualificationId() *string {
	return s.QualificationId
}

func (s *ListHotlineTransferRegisterFileRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListHotlineTransferRegisterFileRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListHotlineTransferRegisterFileRequest) SetHotlineNumber(v string) *ListHotlineTransferRegisterFileRequest {
	s.HotlineNumber = &v
	return s
}

func (s *ListHotlineTransferRegisterFileRequest) SetOwnerId(v int64) *ListHotlineTransferRegisterFileRequest {
	s.OwnerId = &v
	return s
}

func (s *ListHotlineTransferRegisterFileRequest) SetPageNo(v int32) *ListHotlineTransferRegisterFileRequest {
	s.PageNo = &v
	return s
}

func (s *ListHotlineTransferRegisterFileRequest) SetPageSize(v int32) *ListHotlineTransferRegisterFileRequest {
	s.PageSize = &v
	return s
}

func (s *ListHotlineTransferRegisterFileRequest) SetQualificationId(v string) *ListHotlineTransferRegisterFileRequest {
	s.QualificationId = &v
	return s
}

func (s *ListHotlineTransferRegisterFileRequest) SetResourceOwnerAccount(v string) *ListHotlineTransferRegisterFileRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListHotlineTransferRegisterFileRequest) SetResourceOwnerId(v int64) *ListHotlineTransferRegisterFileRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListHotlineTransferRegisterFileRequest) Validate() error {
	return dara.Validate(s)
}
