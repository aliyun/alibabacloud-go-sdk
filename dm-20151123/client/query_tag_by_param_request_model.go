// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTagByParamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyWord(v string) *QueryTagByParamRequest
	GetKeyWord() *string
	SetOwnerId(v int64) *QueryTagByParamRequest
	GetOwnerId() *int64
	SetPageNo(v int32) *QueryTagByParamRequest
	GetPageNo() *int32
	SetPageSize(v int32) *QueryTagByParamRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *QueryTagByParamRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryTagByParamRequest
	GetResourceOwnerId() *int64
}

type QueryTagByParamRequest struct {
	// Tag name, length 1-50, defaults to all tags if not specified.
	//
	// example:
	//
	// 1aTag
	KeyWord *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Page number
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// Page size
	//
	// example:
	//
	// 10
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryTagByParamRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryTagByParamRequest) GoString() string {
	return s.String()
}

func (s *QueryTagByParamRequest) GetKeyWord() *string {
	return s.KeyWord
}

func (s *QueryTagByParamRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryTagByParamRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *QueryTagByParamRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryTagByParamRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryTagByParamRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryTagByParamRequest) SetKeyWord(v string) *QueryTagByParamRequest {
	s.KeyWord = &v
	return s
}

func (s *QueryTagByParamRequest) SetOwnerId(v int64) *QueryTagByParamRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryTagByParamRequest) SetPageNo(v int32) *QueryTagByParamRequest {
	s.PageNo = &v
	return s
}

func (s *QueryTagByParamRequest) SetPageSize(v int32) *QueryTagByParamRequest {
	s.PageSize = &v
	return s
}

func (s *QueryTagByParamRequest) SetResourceOwnerAccount(v string) *QueryTagByParamRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryTagByParamRequest) SetResourceOwnerId(v int64) *QueryTagByParamRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryTagByParamRequest) Validate() error {
	return dara.Validate(s)
}
