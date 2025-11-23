// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMetaCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v *MetaCategory) *UpdateMetaCategoryResponseBody
	GetCategory() *MetaCategory
	SetErrorCode(v string) *UpdateMetaCategoryResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateMetaCategoryResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdateMetaCategoryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateMetaCategoryResponseBody
	GetSuccess() *bool
}

type UpdateMetaCategoryResponseBody struct {
	// The information about the category.
	Category *MetaCategory `json:"Category,omitempty" xml:"Category,omitempty"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message that is returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0C1CB646-1DE4-4AD0-B4A4-7D47DD52E931
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateMetaCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMetaCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMetaCategoryResponseBody) GetCategory() *MetaCategory {
	return s.Category
}

func (s *UpdateMetaCategoryResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateMetaCategoryResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateMetaCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMetaCategoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateMetaCategoryResponseBody) SetCategory(v *MetaCategory) *UpdateMetaCategoryResponseBody {
	s.Category = v
	return s
}

func (s *UpdateMetaCategoryResponseBody) SetErrorCode(v string) *UpdateMetaCategoryResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateMetaCategoryResponseBody) SetErrorMessage(v string) *UpdateMetaCategoryResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateMetaCategoryResponseBody) SetRequestId(v string) *UpdateMetaCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMetaCategoryResponseBody) SetSuccess(v bool) *UpdateMetaCategoryResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateMetaCategoryResponseBody) Validate() error {
	if s.Category != nil {
		if err := s.Category.Validate(); err != nil {
			return err
		}
	}
	return nil
}
