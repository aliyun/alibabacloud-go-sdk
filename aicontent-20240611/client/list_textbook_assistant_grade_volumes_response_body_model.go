// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTextbookAssistantGradeVolumesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListTextbookAssistantGradeVolumesResponseBodyData) *ListTextbookAssistantGradeVolumesResponseBody
	GetData() []*ListTextbookAssistantGradeVolumesResponseBodyData
	SetErrCode(v string) *ListTextbookAssistantGradeVolumesResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ListTextbookAssistantGradeVolumesResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ListTextbookAssistantGradeVolumesResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ListTextbookAssistantGradeVolumesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTextbookAssistantGradeVolumesResponseBody
	GetSuccess() *bool
}

type ListTextbookAssistantGradeVolumesResponseBody struct {
	Data []*ListTextbookAssistantGradeVolumesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	ErrCode    *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// Id of the request
	//
	// example:
	//
	// F3B1AAF2-3041-5AA7-A352-BD5F998FA465
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListTextbookAssistantGradeVolumesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTextbookAssistantGradeVolumesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantGradeVolumesResponseBody) GetData() []*ListTextbookAssistantGradeVolumesResponseBodyData {
	return s.Data
}

func (s *ListTextbookAssistantGradeVolumesResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ListTextbookAssistantGradeVolumesResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ListTextbookAssistantGradeVolumesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListTextbookAssistantGradeVolumesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTextbookAssistantGradeVolumesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTextbookAssistantGradeVolumesResponseBody) SetData(v []*ListTextbookAssistantGradeVolumesResponseBodyData) *ListTextbookAssistantGradeVolumesResponseBody {
	s.Data = v
	return s
}

func (s *ListTextbookAssistantGradeVolumesResponseBody) SetErrCode(v string) *ListTextbookAssistantGradeVolumesResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListTextbookAssistantGradeVolumesResponseBody) SetErrMessage(v string) *ListTextbookAssistantGradeVolumesResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ListTextbookAssistantGradeVolumesResponseBody) SetHttpStatusCode(v int32) *ListTextbookAssistantGradeVolumesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListTextbookAssistantGradeVolumesResponseBody) SetRequestId(v string) *ListTextbookAssistantGradeVolumesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTextbookAssistantGradeVolumesResponseBody) SetSuccess(v bool) *ListTextbookAssistantGradeVolumesResponseBody {
	s.Success = &v
	return s
}

func (s *ListTextbookAssistantGradeVolumesResponseBody) Validate() error {
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

type ListTextbookAssistantGradeVolumesResponseBodyData struct {
	GradeVolumes []*ListTextbookAssistantGradeVolumesResponseBodyDataGradeVolumes `json:"gradeVolumes,omitempty" xml:"gradeVolumes,omitempty" type:"Repeated"`
	// example:
	//
	// 人教版
	TextbookVersion *string `json:"textbookVersion,omitempty" xml:"textbookVersion,omitempty"`
}

func (s ListTextbookAssistantGradeVolumesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListTextbookAssistantGradeVolumesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantGradeVolumesResponseBodyData) GetGradeVolumes() []*ListTextbookAssistantGradeVolumesResponseBodyDataGradeVolumes {
	return s.GradeVolumes
}

func (s *ListTextbookAssistantGradeVolumesResponseBodyData) GetTextbookVersion() *string {
	return s.TextbookVersion
}

func (s *ListTextbookAssistantGradeVolumesResponseBodyData) SetGradeVolumes(v []*ListTextbookAssistantGradeVolumesResponseBodyDataGradeVolumes) *ListTextbookAssistantGradeVolumesResponseBodyData {
	s.GradeVolumes = v
	return s
}

func (s *ListTextbookAssistantGradeVolumesResponseBodyData) SetTextbookVersion(v string) *ListTextbookAssistantGradeVolumesResponseBodyData {
	s.TextbookVersion = &v
	return s
}

func (s *ListTextbookAssistantGradeVolumesResponseBodyData) Validate() error {
	if s.GradeVolumes != nil {
		for _, item := range s.GradeVolumes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTextbookAssistantGradeVolumesResponseBodyDataGradeVolumes struct {
	// example:
	//
	// 3
	Grade *string `json:"grade,omitempty" xml:"grade,omitempty"`
	// example:
	//
	// 1
	Volume *string `json:"volume,omitempty" xml:"volume,omitempty"`
}

func (s ListTextbookAssistantGradeVolumesResponseBodyDataGradeVolumes) String() string {
	return dara.Prettify(s)
}

func (s ListTextbookAssistantGradeVolumesResponseBodyDataGradeVolumes) GoString() string {
	return s.String()
}

func (s *ListTextbookAssistantGradeVolumesResponseBodyDataGradeVolumes) GetGrade() *string {
	return s.Grade
}

func (s *ListTextbookAssistantGradeVolumesResponseBodyDataGradeVolumes) GetVolume() *string {
	return s.Volume
}

func (s *ListTextbookAssistantGradeVolumesResponseBodyDataGradeVolumes) SetGrade(v string) *ListTextbookAssistantGradeVolumesResponseBodyDataGradeVolumes {
	s.Grade = &v
	return s
}

func (s *ListTextbookAssistantGradeVolumesResponseBodyDataGradeVolumes) SetVolume(v string) *ListTextbookAssistantGradeVolumesResponseBodyDataGradeVolumes {
	s.Volume = &v
	return s
}

func (s *ListTextbookAssistantGradeVolumesResponseBodyDataGradeVolumes) Validate() error {
	return dara.Validate(s)
}
