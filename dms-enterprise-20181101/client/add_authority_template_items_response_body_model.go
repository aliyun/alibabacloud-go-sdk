// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAuthorityTemplateItemsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *AddAuthorityTemplateItemsResponseBody
	GetData() *bool
	SetErrorCode(v string) *AddAuthorityTemplateItemsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *AddAuthorityTemplateItemsResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *AddAuthorityTemplateItemsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddAuthorityTemplateItemsResponseBody
	GetSuccess() *bool
}

type AddAuthorityTemplateItemsResponseBody struct {
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 3D3FB827-E667-50DB-AD59-C83F8237FECB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddAuthorityTemplateItemsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddAuthorityTemplateItemsResponseBody) GoString() string {
	return s.String()
}

func (s *AddAuthorityTemplateItemsResponseBody) GetData() *bool {
	return s.Data
}

func (s *AddAuthorityTemplateItemsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *AddAuthorityTemplateItemsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *AddAuthorityTemplateItemsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddAuthorityTemplateItemsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddAuthorityTemplateItemsResponseBody) SetData(v bool) *AddAuthorityTemplateItemsResponseBody {
	s.Data = &v
	return s
}

func (s *AddAuthorityTemplateItemsResponseBody) SetErrorCode(v string) *AddAuthorityTemplateItemsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AddAuthorityTemplateItemsResponseBody) SetErrorMessage(v string) *AddAuthorityTemplateItemsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AddAuthorityTemplateItemsResponseBody) SetRequestId(v string) *AddAuthorityTemplateItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddAuthorityTemplateItemsResponseBody) SetSuccess(v bool) *AddAuthorityTemplateItemsResponseBody {
	s.Success = &v
	return s
}

func (s *AddAuthorityTemplateItemsResponseBody) Validate() error {
	return dara.Validate(s)
}
