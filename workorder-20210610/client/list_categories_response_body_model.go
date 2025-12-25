// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCategoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListCategoriesResponseBody
	GetCode() *int32
	SetData(v []*ListCategoriesResponseBodyData) *ListCategoriesResponseBody
	GetData() []*ListCategoriesResponseBodyData
	SetMessage(v string) *ListCategoriesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListCategoriesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListCategoriesResponseBody
	GetSuccess() *bool
}

type ListCategoriesResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The return value, which is a list of question categories
	Data []*ListCategoriesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error message. If success is set to false, the message is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// CA6204AC-6AA9-4CFA-9310-7DFD20C19EBC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values: The value true indicates a success. The value false indicates a failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCategoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCategoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCategoriesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListCategoriesResponseBody) GetData() []*ListCategoriesResponseBodyData {
	return s.Data
}

func (s *ListCategoriesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListCategoriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCategoriesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListCategoriesResponseBody) SetCode(v int32) *ListCategoriesResponseBody {
	s.Code = &v
	return s
}

func (s *ListCategoriesResponseBody) SetData(v []*ListCategoriesResponseBodyData) *ListCategoriesResponseBody {
	s.Data = v
	return s
}

func (s *ListCategoriesResponseBody) SetMessage(v string) *ListCategoriesResponseBody {
	s.Message = &v
	return s
}

func (s *ListCategoriesResponseBody) SetRequestId(v string) *ListCategoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCategoriesResponseBody) SetSuccess(v bool) *ListCategoriesResponseBody {
	s.Success = &v
	return s
}

func (s *ListCategoriesResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCategoriesResponseBodyData struct {
	// The ID of the ticket issue category.
	//
	// example:
	//
	// 7161
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The name of the ticket issue category.
	//
	// example:
	//
	// ECS
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
}

func (s ListCategoriesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListCategoriesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCategoriesResponseBodyData) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *ListCategoriesResponseBodyData) GetCategoryName() *string {
	return s.CategoryName
}

func (s *ListCategoriesResponseBodyData) SetCategoryId(v int64) *ListCategoriesResponseBodyData {
	s.CategoryId = &v
	return s
}

func (s *ListCategoriesResponseBodyData) SetCategoryName(v string) *ListCategoriesResponseBodyData {
	s.CategoryName = &v
	return s
}

func (s *ListCategoriesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
