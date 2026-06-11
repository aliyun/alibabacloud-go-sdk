// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSkillFileCheckResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateSkillFileCheckResponseBodyData) *CreateSkillFileCheckResponseBody
	GetData() *CreateSkillFileCheckResponseBodyData
	SetRequestId(v string) *CreateSkillFileCheckResponseBody
	GetRequestId() *string
}

type CreateSkillFileCheckResponseBody struct {
	Data      *CreateSkillFileCheckResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSkillFileCheckResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSkillFileCheckResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSkillFileCheckResponseBody) GetData() *CreateSkillFileCheckResponseBodyData {
	return s.Data
}

func (s *CreateSkillFileCheckResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSkillFileCheckResponseBody) SetData(v *CreateSkillFileCheckResponseBodyData) *CreateSkillFileCheckResponseBody {
	s.Data = v
	return s
}

func (s *CreateSkillFileCheckResponseBody) SetRequestId(v string) *CreateSkillFileCheckResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSkillFileCheckResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateSkillFileCheckResponseBodyData struct {
	FailCount     *int32                                               `json:"FailCount,omitempty" xml:"FailCount,omitempty"`
	RootTaskId    *string                                              `json:"RootTaskId,omitempty" xml:"RootTaskId,omitempty"`
	SuccessCount  *int32                                               `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	UploadResults []*CreateSkillFileCheckResponseBodyDataUploadResults `json:"UploadResults,omitempty" xml:"UploadResults,omitempty" type:"Repeated"`
}

func (s CreateSkillFileCheckResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateSkillFileCheckResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateSkillFileCheckResponseBodyData) GetFailCount() *int32 {
	return s.FailCount
}

func (s *CreateSkillFileCheckResponseBodyData) GetRootTaskId() *string {
	return s.RootTaskId
}

func (s *CreateSkillFileCheckResponseBodyData) GetSuccessCount() *int32 {
	return s.SuccessCount
}

func (s *CreateSkillFileCheckResponseBodyData) GetUploadResults() []*CreateSkillFileCheckResponseBodyDataUploadResults {
	return s.UploadResults
}

func (s *CreateSkillFileCheckResponseBodyData) SetFailCount(v int32) *CreateSkillFileCheckResponseBodyData {
	s.FailCount = &v
	return s
}

func (s *CreateSkillFileCheckResponseBodyData) SetRootTaskId(v string) *CreateSkillFileCheckResponseBodyData {
	s.RootTaskId = &v
	return s
}

func (s *CreateSkillFileCheckResponseBodyData) SetSuccessCount(v int32) *CreateSkillFileCheckResponseBodyData {
	s.SuccessCount = &v
	return s
}

func (s *CreateSkillFileCheckResponseBodyData) SetUploadResults(v []*CreateSkillFileCheckResponseBodyDataUploadResults) *CreateSkillFileCheckResponseBodyData {
	s.UploadResults = v
	return s
}

func (s *CreateSkillFileCheckResponseBodyData) Validate() error {
	if s.UploadResults != nil {
		for _, item := range s.UploadResults {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateSkillFileCheckResponseBodyDataUploadResults struct {
	ErrorMsg   *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	FileHash   *string `json:"FileHash,omitempty" xml:"FileHash,omitempty"`
	FileName   *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	IdentifyId *string `json:"IdentifyId,omitempty" xml:"IdentifyId,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSkillFileCheckResponseBodyDataUploadResults) String() string {
	return dara.Prettify(s)
}

func (s CreateSkillFileCheckResponseBodyDataUploadResults) GoString() string {
	return s.String()
}

func (s *CreateSkillFileCheckResponseBodyDataUploadResults) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *CreateSkillFileCheckResponseBodyDataUploadResults) GetFileHash() *string {
	return s.FileHash
}

func (s *CreateSkillFileCheckResponseBodyDataUploadResults) GetFileName() *string {
	return s.FileName
}

func (s *CreateSkillFileCheckResponseBodyDataUploadResults) GetIdentifyId() *string {
	return s.IdentifyId
}

func (s *CreateSkillFileCheckResponseBodyDataUploadResults) GetSuccess() *bool {
	return s.Success
}

func (s *CreateSkillFileCheckResponseBodyDataUploadResults) SetErrorMsg(v string) *CreateSkillFileCheckResponseBodyDataUploadResults {
	s.ErrorMsg = &v
	return s
}

func (s *CreateSkillFileCheckResponseBodyDataUploadResults) SetFileHash(v string) *CreateSkillFileCheckResponseBodyDataUploadResults {
	s.FileHash = &v
	return s
}

func (s *CreateSkillFileCheckResponseBodyDataUploadResults) SetFileName(v string) *CreateSkillFileCheckResponseBodyDataUploadResults {
	s.FileName = &v
	return s
}

func (s *CreateSkillFileCheckResponseBodyDataUploadResults) SetIdentifyId(v string) *CreateSkillFileCheckResponseBodyDataUploadResults {
	s.IdentifyId = &v
	return s
}

func (s *CreateSkillFileCheckResponseBodyDataUploadResults) SetSuccess(v bool) *CreateSkillFileCheckResponseBodyDataUploadResults {
	s.Success = &v
	return s
}

func (s *CreateSkillFileCheckResponseBodyDataUploadResults) Validate() error {
	return dara.Validate(s)
}
