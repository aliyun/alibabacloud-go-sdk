// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAICaptionExtractionJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAICaptionExtractionJobList(v []*GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList) *GetAICaptionExtractionJobsResponseBody
	GetAICaptionExtractionJobList() []*GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList
	SetRequestId(v string) *GetAICaptionExtractionJobsResponseBody
	GetRequestId() *string
}

type GetAICaptionExtractionJobsResponseBody struct {
	AICaptionExtractionJobList []*GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList `json:"AICaptionExtractionJobList,omitempty" xml:"AICaptionExtractionJobList,omitempty" type:"Repeated"`
	RequestId                  *string                                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAICaptionExtractionJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAICaptionExtractionJobsResponseBody) GoString() string {
	return s.String()
}

func (s *GetAICaptionExtractionJobsResponseBody) GetAICaptionExtractionJobList() []*GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList {
	return s.AICaptionExtractionJobList
}

func (s *GetAICaptionExtractionJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAICaptionExtractionJobsResponseBody) SetAICaptionExtractionJobList(v []*GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList) *GetAICaptionExtractionJobsResponseBody {
	s.AICaptionExtractionJobList = v
	return s
}

func (s *GetAICaptionExtractionJobsResponseBody) SetRequestId(v string) *GetAICaptionExtractionJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAICaptionExtractionJobsResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList struct {
	AICaptionExtractionResult *string `json:"AICaptionExtractionResult,omitempty" xml:"AICaptionExtractionResult,omitempty"`
	Code                      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	CreationTime              *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	JobId                     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Message                   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Status                    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TemplateConfig            *string `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty"`
	UserData                  *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	VideoId                   *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList) String() string {
	return dara.Prettify(s)
}

func (s GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList) GoString() string {
	return s.String()
}

func (s *GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList) GetAICaptionExtractionResult() *string {
	return s.AICaptionExtractionResult
}

func (s *GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList) GetCode() *string {
	return s.Code
}

func (s *GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList) GetCreationTime() *string {
	return s.CreationTime
}

func (s *GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList) GetJobId() *string {
	return s.JobId
}

func (s *GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList) GetMessage() *string {
	return s.Message
}

func (s *GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList) GetStatus() *string {
	return s.Status
}

func (s *GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList) GetTemplateConfig() *string {
	return s.TemplateConfig
}

func (s *GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList) GetUserData() *string {
	return s.UserData
}

func (s *GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList) GetVideoId() *string {
	return s.VideoId
}

func (s *GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList) SetAICaptionExtractionResult(v string) *GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList {
	s.AICaptionExtractionResult = &v
	return s
}

func (s *GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList) SetCode(v string) *GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList {
	s.Code = &v
	return s
}

func (s *GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList) SetCreationTime(v string) *GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList {
	s.CreationTime = &v
	return s
}

func (s *GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList) SetJobId(v string) *GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList {
	s.JobId = &v
	return s
}

func (s *GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList) SetMessage(v string) *GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList {
	s.Message = &v
	return s
}

func (s *GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList) SetStatus(v string) *GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList {
	s.Status = &v
	return s
}

func (s *GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList) SetTemplateConfig(v string) *GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList {
	s.TemplateConfig = &v
	return s
}

func (s *GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList) SetUserData(v string) *GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList {
	s.UserData = &v
	return s
}

func (s *GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList) SetVideoId(v string) *GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList {
	s.VideoId = &v
	return s
}

func (s *GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList) Validate() error {
	return dara.Validate(s)
}
