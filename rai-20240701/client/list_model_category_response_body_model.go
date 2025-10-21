// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListModelCategoryResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListModelCategoryResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListModelCategoryResponseBody
	GetMessage() *string
	SetModelCategoryInfoList(v []*ListModelCategoryResponseBodyModelCategoryInfoList) *ListModelCategoryResponseBody
	GetModelCategoryInfoList() []*ListModelCategoryResponseBodyModelCategoryInfoList
	SetRequestId(v string) *ListModelCategoryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListModelCategoryResponseBody
	GetSuccess() *bool
}

type ListModelCategoryResponseBody struct {
	// example:
	//
	// 00000
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// ""
	Message               *string                                               `json:"Message,omitempty" xml:"Message,omitempty"`
	ModelCategoryInfoList []*ListModelCategoryResponseBodyModelCategoryInfoList `json:"ModelCategoryInfoList,omitempty" xml:"ModelCategoryInfoList,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListModelCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListModelCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListModelCategoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListModelCategoryResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListModelCategoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListModelCategoryResponseBody) GetModelCategoryInfoList() []*ListModelCategoryResponseBodyModelCategoryInfoList {
	return s.ModelCategoryInfoList
}

func (s *ListModelCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListModelCategoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListModelCategoryResponseBody) SetCode(v string) *ListModelCategoryResponseBody {
	s.Code = &v
	return s
}

func (s *ListModelCategoryResponseBody) SetHttpStatusCode(v int32) *ListModelCategoryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListModelCategoryResponseBody) SetMessage(v string) *ListModelCategoryResponseBody {
	s.Message = &v
	return s
}

func (s *ListModelCategoryResponseBody) SetModelCategoryInfoList(v []*ListModelCategoryResponseBodyModelCategoryInfoList) *ListModelCategoryResponseBody {
	s.ModelCategoryInfoList = v
	return s
}

func (s *ListModelCategoryResponseBody) SetRequestId(v string) *ListModelCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListModelCategoryResponseBody) SetSuccess(v bool) *ListModelCategoryResponseBody {
	s.Success = &v
	return s
}

func (s *ListModelCategoryResponseBody) Validate() error {
	if s.ModelCategoryInfoList != nil {
		for _, item := range s.ModelCategoryInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListModelCategoryResponseBodyModelCategoryInfoList struct {
	// example:
	//
	// False
	ContentSafeImageSupported *int32 `json:"ContentSafeImageSupported,omitempty" xml:"ContentSafeImageSupported,omitempty"`
	// example:
	//
	// True
	ContentSafeTextSupported *int32 `json:"ContentSafeTextSupported,omitempty" xml:"ContentSafeTextSupported,omitempty"`
	// example:
	//
	// 2
	ModelCategoryId *int64 `json:"ModelCategoryId,omitempty" xml:"ModelCategoryId,omitempty"`
	// example:
	//
	// Qwen2.5-3B-Intruct-PAI-Guard
	ModelCategoryName *string `json:"ModelCategoryName,omitempty" xml:"ModelCategoryName,omitempty"`
	// example:
	//
	// 1
	ModelSource *int32 `json:"ModelSource,omitempty" xml:"ModelSource,omitempty"`
	// example:
	//
	// 12
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// False
	PromptAttackTextSupported *int32 `json:"PromptAttackTextSupported,omitempty" xml:"PromptAttackTextSupported,omitempty"`
	// example:
	//
	// True
	SensitiveTopicTextSupported *int32 `json:"SensitiveTopicTextSupported,omitempty" xml:"SensitiveTopicTextSupported,omitempty"`
}

func (s ListModelCategoryResponseBodyModelCategoryInfoList) String() string {
	return dara.Prettify(s)
}

func (s ListModelCategoryResponseBodyModelCategoryInfoList) GoString() string {
	return s.String()
}

func (s *ListModelCategoryResponseBodyModelCategoryInfoList) GetContentSafeImageSupported() *int32 {
	return s.ContentSafeImageSupported
}

func (s *ListModelCategoryResponseBodyModelCategoryInfoList) GetContentSafeTextSupported() *int32 {
	return s.ContentSafeTextSupported
}

func (s *ListModelCategoryResponseBodyModelCategoryInfoList) GetModelCategoryId() *int64 {
	return s.ModelCategoryId
}

func (s *ListModelCategoryResponseBodyModelCategoryInfoList) GetModelCategoryName() *string {
	return s.ModelCategoryName
}

func (s *ListModelCategoryResponseBodyModelCategoryInfoList) GetModelSource() *int32 {
	return s.ModelSource
}

func (s *ListModelCategoryResponseBodyModelCategoryInfoList) GetPriority() *int32 {
	return s.Priority
}

func (s *ListModelCategoryResponseBodyModelCategoryInfoList) GetPromptAttackTextSupported() *int32 {
	return s.PromptAttackTextSupported
}

func (s *ListModelCategoryResponseBodyModelCategoryInfoList) GetSensitiveTopicTextSupported() *int32 {
	return s.SensitiveTopicTextSupported
}

func (s *ListModelCategoryResponseBodyModelCategoryInfoList) SetContentSafeImageSupported(v int32) *ListModelCategoryResponseBodyModelCategoryInfoList {
	s.ContentSafeImageSupported = &v
	return s
}

func (s *ListModelCategoryResponseBodyModelCategoryInfoList) SetContentSafeTextSupported(v int32) *ListModelCategoryResponseBodyModelCategoryInfoList {
	s.ContentSafeTextSupported = &v
	return s
}

func (s *ListModelCategoryResponseBodyModelCategoryInfoList) SetModelCategoryId(v int64) *ListModelCategoryResponseBodyModelCategoryInfoList {
	s.ModelCategoryId = &v
	return s
}

func (s *ListModelCategoryResponseBodyModelCategoryInfoList) SetModelCategoryName(v string) *ListModelCategoryResponseBodyModelCategoryInfoList {
	s.ModelCategoryName = &v
	return s
}

func (s *ListModelCategoryResponseBodyModelCategoryInfoList) SetModelSource(v int32) *ListModelCategoryResponseBodyModelCategoryInfoList {
	s.ModelSource = &v
	return s
}

func (s *ListModelCategoryResponseBodyModelCategoryInfoList) SetPriority(v int32) *ListModelCategoryResponseBodyModelCategoryInfoList {
	s.Priority = &v
	return s
}

func (s *ListModelCategoryResponseBodyModelCategoryInfoList) SetPromptAttackTextSupported(v int32) *ListModelCategoryResponseBodyModelCategoryInfoList {
	s.PromptAttackTextSupported = &v
	return s
}

func (s *ListModelCategoryResponseBodyModelCategoryInfoList) SetSensitiveTopicTextSupported(v int32) *ListModelCategoryResponseBodyModelCategoryInfoList {
	s.SensitiveTopicTextSupported = &v
	return s
}

func (s *ListModelCategoryResponseBodyModelCategoryInfoList) Validate() error {
	return dara.Validate(s)
}
