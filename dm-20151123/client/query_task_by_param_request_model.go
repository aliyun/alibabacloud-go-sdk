// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTaskByParamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyWord(v string) *QueryTaskByParamRequest
	GetKeyWord() *string
	SetOwnerId(v int64) *QueryTaskByParamRequest
	GetOwnerId() *int64
	SetPageNo(v int32) *QueryTaskByParamRequest
	GetPageNo() *int32
	SetPageSize(v int32) *QueryTaskByParamRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *QueryTaskByParamRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryTaskByParamRequest
	GetResourceOwnerId() *int64
	SetStatus(v int32) *QueryTaskByParamRequest
	GetStatus() *int32
}

type QueryTaskByParamRequest struct {
	// Keyword, defaults to all information
	//
	// example:
	//
	// mesh-notification-788717
	KeyWord *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Current page number, default is 1
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// Page size, default is 10
	//
	// example:
	//
	// 10
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Status, defaults to all statuses
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryTaskByParamRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryTaskByParamRequest) GoString() string {
	return s.String()
}

func (s *QueryTaskByParamRequest) GetKeyWord() *string {
	return s.KeyWord
}

func (s *QueryTaskByParamRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryTaskByParamRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *QueryTaskByParamRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryTaskByParamRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryTaskByParamRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryTaskByParamRequest) GetStatus() *int32 {
	return s.Status
}

func (s *QueryTaskByParamRequest) SetKeyWord(v string) *QueryTaskByParamRequest {
	s.KeyWord = &v
	return s
}

func (s *QueryTaskByParamRequest) SetOwnerId(v int64) *QueryTaskByParamRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryTaskByParamRequest) SetPageNo(v int32) *QueryTaskByParamRequest {
	s.PageNo = &v
	return s
}

func (s *QueryTaskByParamRequest) SetPageSize(v int32) *QueryTaskByParamRequest {
	s.PageSize = &v
	return s
}

func (s *QueryTaskByParamRequest) SetResourceOwnerAccount(v string) *QueryTaskByParamRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryTaskByParamRequest) SetResourceOwnerId(v int64) *QueryTaskByParamRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryTaskByParamRequest) SetStatus(v int32) *QueryTaskByParamRequest {
	s.Status = &v
	return s
}

func (s *QueryTaskByParamRequest) Validate() error {
	return dara.Validate(s)
}
