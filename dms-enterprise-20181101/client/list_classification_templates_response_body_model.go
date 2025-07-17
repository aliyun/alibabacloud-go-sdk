// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClassificationTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListClassificationTemplatesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListClassificationTemplatesResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListClassificationTemplatesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListClassificationTemplatesResponseBody
	GetSuccess() *bool
	SetTemplateList(v []*ListClassificationTemplatesResponseBodyTemplateList) *ListClassificationTemplatesResponseBody
	GetTemplateList() []*ListClassificationTemplatesResponseBodyTemplateList
}

type ListClassificationTemplatesResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 90260530-565C-42B9-A6E8-893481FE6AB6
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
	// The list of templates.
	TemplateList []*ListClassificationTemplatesResponseBodyTemplateList `json:"TemplateList,omitempty" xml:"TemplateList,omitempty" type:"Repeated"`
}

func (s ListClassificationTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListClassificationTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListClassificationTemplatesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListClassificationTemplatesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListClassificationTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListClassificationTemplatesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListClassificationTemplatesResponseBody) GetTemplateList() []*ListClassificationTemplatesResponseBodyTemplateList {
	return s.TemplateList
}

func (s *ListClassificationTemplatesResponseBody) SetErrorCode(v string) *ListClassificationTemplatesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListClassificationTemplatesResponseBody) SetErrorMessage(v string) *ListClassificationTemplatesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListClassificationTemplatesResponseBody) SetRequestId(v string) *ListClassificationTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClassificationTemplatesResponseBody) SetSuccess(v bool) *ListClassificationTemplatesResponseBody {
	s.Success = &v
	return s
}

func (s *ListClassificationTemplatesResponseBody) SetTemplateList(v []*ListClassificationTemplatesResponseBodyTemplateList) *ListClassificationTemplatesResponseBody {
	s.TemplateList = v
	return s
}

func (s *ListClassificationTemplatesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListClassificationTemplatesResponseBodyTemplateList struct {
	// The name of the classification template.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The remarks.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The ID of the classification template.
	//
	// example:
	//
	// 3**
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The type of the classification template. Valid values:
	//
	// 	- **INNER**: built-in template
	//
	// 	- **USER_DEFINE**: custom template
	//
	// example:
	//
	// USER_DEFINE
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s ListClassificationTemplatesResponseBodyTemplateList) String() string {
	return dara.Prettify(s)
}

func (s ListClassificationTemplatesResponseBodyTemplateList) GoString() string {
	return s.String()
}

func (s *ListClassificationTemplatesResponseBodyTemplateList) GetName() *string {
	return s.Name
}

func (s *ListClassificationTemplatesResponseBodyTemplateList) GetRemark() *string {
	return s.Remark
}

func (s *ListClassificationTemplatesResponseBodyTemplateList) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *ListClassificationTemplatesResponseBodyTemplateList) GetTemplateType() *string {
	return s.TemplateType
}

func (s *ListClassificationTemplatesResponseBodyTemplateList) SetName(v string) *ListClassificationTemplatesResponseBodyTemplateList {
	s.Name = &v
	return s
}

func (s *ListClassificationTemplatesResponseBodyTemplateList) SetRemark(v string) *ListClassificationTemplatesResponseBodyTemplateList {
	s.Remark = &v
	return s
}

func (s *ListClassificationTemplatesResponseBodyTemplateList) SetTemplateId(v int64) *ListClassificationTemplatesResponseBodyTemplateList {
	s.TemplateId = &v
	return s
}

func (s *ListClassificationTemplatesResponseBodyTemplateList) SetTemplateType(v string) *ListClassificationTemplatesResponseBodyTemplateList {
	s.TemplateType = &v
	return s
}

func (s *ListClassificationTemplatesResponseBodyTemplateList) Validate() error {
	return dara.Validate(s)
}
