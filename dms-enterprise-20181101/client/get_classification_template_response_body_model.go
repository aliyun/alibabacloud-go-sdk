// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClassificationTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClassificationResourceTemplateMap(v *GetClassificationTemplateResponseBodyClassificationResourceTemplateMap) *GetClassificationTemplateResponseBody
	GetClassificationResourceTemplateMap() *GetClassificationTemplateResponseBodyClassificationResourceTemplateMap
	SetErrorCode(v string) *GetClassificationTemplateResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetClassificationTemplateResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetClassificationTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetClassificationTemplateResponseBody
	GetSuccess() *bool
}

type GetClassificationTemplateResponseBody struct {
	ClassificationResourceTemplateMap *GetClassificationTemplateResponseBodyClassificationResourceTemplateMap `json:"ClassificationResourceTemplateMap,omitempty" xml:"ClassificationResourceTemplateMap,omitempty" type:"Struct"`
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
	// 427688B8-ADFB-4C4E-9D45-EF5C1FD6E23D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetClassificationTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetClassificationTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetClassificationTemplateResponseBody) GetClassificationResourceTemplateMap() *GetClassificationTemplateResponseBodyClassificationResourceTemplateMap {
	return s.ClassificationResourceTemplateMap
}

func (s *GetClassificationTemplateResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetClassificationTemplateResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetClassificationTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetClassificationTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetClassificationTemplateResponseBody) SetClassificationResourceTemplateMap(v *GetClassificationTemplateResponseBodyClassificationResourceTemplateMap) *GetClassificationTemplateResponseBody {
	s.ClassificationResourceTemplateMap = v
	return s
}

func (s *GetClassificationTemplateResponseBody) SetErrorCode(v string) *GetClassificationTemplateResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetClassificationTemplateResponseBody) SetErrorMessage(v string) *GetClassificationTemplateResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetClassificationTemplateResponseBody) SetRequestId(v string) *GetClassificationTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetClassificationTemplateResponseBody) SetSuccess(v bool) *GetClassificationTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *GetClassificationTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetClassificationTemplateResponseBodyClassificationResourceTemplateMap struct {
	// example:
	//
	// 24****
	ResourceId *int64 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// INSTANCE
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// 3***
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// example:
	//
	// INNER
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s GetClassificationTemplateResponseBodyClassificationResourceTemplateMap) String() string {
	return dara.Prettify(s)
}

func (s GetClassificationTemplateResponseBodyClassificationResourceTemplateMap) GoString() string {
	return s.String()
}

func (s *GetClassificationTemplateResponseBodyClassificationResourceTemplateMap) GetResourceId() *int64 {
	return s.ResourceId
}

func (s *GetClassificationTemplateResponseBodyClassificationResourceTemplateMap) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetClassificationTemplateResponseBodyClassificationResourceTemplateMap) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *GetClassificationTemplateResponseBodyClassificationResourceTemplateMap) GetTemplateType() *string {
	return s.TemplateType
}

func (s *GetClassificationTemplateResponseBodyClassificationResourceTemplateMap) SetResourceId(v int64) *GetClassificationTemplateResponseBodyClassificationResourceTemplateMap {
	s.ResourceId = &v
	return s
}

func (s *GetClassificationTemplateResponseBodyClassificationResourceTemplateMap) SetResourceType(v string) *GetClassificationTemplateResponseBodyClassificationResourceTemplateMap {
	s.ResourceType = &v
	return s
}

func (s *GetClassificationTemplateResponseBodyClassificationResourceTemplateMap) SetTemplateId(v int64) *GetClassificationTemplateResponseBodyClassificationResourceTemplateMap {
	s.TemplateId = &v
	return s
}

func (s *GetClassificationTemplateResponseBodyClassificationResourceTemplateMap) SetTemplateType(v string) *GetClassificationTemplateResponseBodyClassificationResourceTemplateMap {
	s.TemplateType = &v
	return s
}

func (s *GetClassificationTemplateResponseBodyClassificationResourceTemplateMap) Validate() error {
	return dara.Validate(s)
}
