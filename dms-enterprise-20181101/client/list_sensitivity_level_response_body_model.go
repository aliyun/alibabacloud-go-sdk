// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSensitivityLevelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListSensitivityLevelResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListSensitivityLevelResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListSensitivityLevelResponseBody
	GetRequestId() *string
	SetSensitivityLevelList(v []*ListSensitivityLevelResponseBodySensitivityLevelList) *ListSensitivityLevelResponseBody
	GetSensitivityLevelList() []*ListSensitivityLevelResponseBodySensitivityLevelList
	SetSuccess(v bool) *ListSensitivityLevelResponseBody
	GetSuccess() *bool
}

type ListSensitivityLevelResponseBody struct {
	// The status code.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID. You can use the ID to query logs and troubleshoot issues.
	//
	// example:
	//
	// 4E1D2B4D-3E53-4ABC-999D-1D2520B3471A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The sensitivity levels.
	SensitivityLevelList []*ListSensitivityLevelResponseBodySensitivityLevelList `json:"SensitivityLevelList,omitempty" xml:"SensitivityLevelList,omitempty" type:"Repeated"`
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

func (s ListSensitivityLevelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSensitivityLevelResponseBody) GoString() string {
	return s.String()
}

func (s *ListSensitivityLevelResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListSensitivityLevelResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListSensitivityLevelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSensitivityLevelResponseBody) GetSensitivityLevelList() []*ListSensitivityLevelResponseBodySensitivityLevelList {
	return s.SensitivityLevelList
}

func (s *ListSensitivityLevelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSensitivityLevelResponseBody) SetErrorCode(v string) *ListSensitivityLevelResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListSensitivityLevelResponseBody) SetErrorMessage(v string) *ListSensitivityLevelResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListSensitivityLevelResponseBody) SetRequestId(v string) *ListSensitivityLevelResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSensitivityLevelResponseBody) SetSensitivityLevelList(v []*ListSensitivityLevelResponseBodySensitivityLevelList) *ListSensitivityLevelResponseBody {
	s.SensitivityLevelList = v
	return s
}

func (s *ListSensitivityLevelResponseBody) SetSuccess(v bool) *ListSensitivityLevelResponseBody {
	s.Success = &v
	return s
}

func (s *ListSensitivityLevelResponseBody) Validate() error {
	if s.SensitivityLevelList != nil {
		for _, item := range s.SensitivityLevelList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSensitivityLevelResponseBodySensitivityLevelList struct {
	// Indicates whether the fields of the sensitive level are displayed in plaintext.
	//
	// example:
	//
	// true
	IsPlain *bool `json:"IsPlain,omitempty" xml:"IsPlain,omitempty"`
	// The name of the sensitive level.
	//
	// example:
	//
	// S2
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the classification template.
	//
	// example:
	//
	// 1070
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The type of the classification template. Valid values:
	//
	// 	- **INNER**: a built-in template.
	//
	// 	- **USER_DEFINE**: a custom template.
	//
	// example:
	//
	// INNER
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s ListSensitivityLevelResponseBodySensitivityLevelList) String() string {
	return dara.Prettify(s)
}

func (s ListSensitivityLevelResponseBodySensitivityLevelList) GoString() string {
	return s.String()
}

func (s *ListSensitivityLevelResponseBodySensitivityLevelList) GetIsPlain() *bool {
	return s.IsPlain
}

func (s *ListSensitivityLevelResponseBodySensitivityLevelList) GetName() *string {
	return s.Name
}

func (s *ListSensitivityLevelResponseBodySensitivityLevelList) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListSensitivityLevelResponseBodySensitivityLevelList) GetTemplateType() *string {
	return s.TemplateType
}

func (s *ListSensitivityLevelResponseBodySensitivityLevelList) SetIsPlain(v bool) *ListSensitivityLevelResponseBodySensitivityLevelList {
	s.IsPlain = &v
	return s
}

func (s *ListSensitivityLevelResponseBodySensitivityLevelList) SetName(v string) *ListSensitivityLevelResponseBodySensitivityLevelList {
	s.Name = &v
	return s
}

func (s *ListSensitivityLevelResponseBodySensitivityLevelList) SetTemplateId(v string) *ListSensitivityLevelResponseBodySensitivityLevelList {
	s.TemplateId = &v
	return s
}

func (s *ListSensitivityLevelResponseBodySensitivityLevelList) SetTemplateType(v string) *ListSensitivityLevelResponseBodySensitivityLevelList {
	s.TemplateType = &v
	return s
}

func (s *ListSensitivityLevelResponseBodySensitivityLevelList) Validate() error {
	return dara.Validate(s)
}
