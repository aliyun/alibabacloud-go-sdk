// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMetaCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v *MetaCategory) *CreateMetaCategoryResponseBody
	GetCategory() *MetaCategory
	SetErrorCode(v string) *CreateMetaCategoryResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateMetaCategoryResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateMetaCategoryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateMetaCategoryResponseBody
	GetSuccess() *bool
}

type CreateMetaCategoryResponseBody struct {
	Category *MetaCategory `json:"Category,omitempty" xml:"Category,omitempty"`
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
	// 207176D7-A9B3-55CE-A9DA-14E223A31913
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateMetaCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMetaCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMetaCategoryResponseBody) GetCategory() *MetaCategory {
	return s.Category
}

func (s *CreateMetaCategoryResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateMetaCategoryResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateMetaCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMetaCategoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateMetaCategoryResponseBody) SetCategory(v *MetaCategory) *CreateMetaCategoryResponseBody {
	s.Category = v
	return s
}

func (s *CreateMetaCategoryResponseBody) SetErrorCode(v string) *CreateMetaCategoryResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateMetaCategoryResponseBody) SetErrorMessage(v string) *CreateMetaCategoryResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateMetaCategoryResponseBody) SetRequestId(v string) *CreateMetaCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMetaCategoryResponseBody) SetSuccess(v bool) *CreateMetaCategoryResponseBody {
	s.Success = &v
	return s
}

func (s *CreateMetaCategoryResponseBody) Validate() error {
	return dara.Validate(s)
}
