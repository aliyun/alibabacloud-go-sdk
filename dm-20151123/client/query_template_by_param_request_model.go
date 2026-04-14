// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTemplateByParamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFromType(v int32) *QueryTemplateByParamRequest
	GetFromType() *int32
	SetKeyWord(v string) *QueryTemplateByParamRequest
	GetKeyWord() *string
	SetOwnerId(v int64) *QueryTemplateByParamRequest
	GetOwnerId() *int64
	SetPageNo(v int32) *QueryTemplateByParamRequest
	GetPageNo() *int32
	SetPageSize(v int32) *QueryTemplateByParamRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *QueryTemplateByParamRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryTemplateByParamRequest
	GetResourceOwnerId() *int64
	SetStatus(v int32) *QueryTemplateByParamRequest
	GetStatus() *int32
}

type QueryTemplateByParamRequest struct {
	FromType *int32 `json:"FromType,omitempty" xml:"FromType,omitempty"`
	// example:
	//
	// test
	KeyWord *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 20
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Status               *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryTemplateByParamRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryTemplateByParamRequest) GoString() string {
	return s.String()
}

func (s *QueryTemplateByParamRequest) GetFromType() *int32 {
	return s.FromType
}

func (s *QueryTemplateByParamRequest) GetKeyWord() *string {
	return s.KeyWord
}

func (s *QueryTemplateByParamRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryTemplateByParamRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *QueryTemplateByParamRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryTemplateByParamRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryTemplateByParamRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryTemplateByParamRequest) GetStatus() *int32 {
	return s.Status
}

func (s *QueryTemplateByParamRequest) SetFromType(v int32) *QueryTemplateByParamRequest {
	s.FromType = &v
	return s
}

func (s *QueryTemplateByParamRequest) SetKeyWord(v string) *QueryTemplateByParamRequest {
	s.KeyWord = &v
	return s
}

func (s *QueryTemplateByParamRequest) SetOwnerId(v int64) *QueryTemplateByParamRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryTemplateByParamRequest) SetPageNo(v int32) *QueryTemplateByParamRequest {
	s.PageNo = &v
	return s
}

func (s *QueryTemplateByParamRequest) SetPageSize(v int32) *QueryTemplateByParamRequest {
	s.PageSize = &v
	return s
}

func (s *QueryTemplateByParamRequest) SetResourceOwnerAccount(v string) *QueryTemplateByParamRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryTemplateByParamRequest) SetResourceOwnerId(v int64) *QueryTemplateByParamRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryTemplateByParamRequest) SetStatus(v int32) *QueryTemplateByParamRequest {
	s.Status = &v
	return s
}

func (s *QueryTemplateByParamRequest) Validate() error {
	return dara.Validate(s)
}
