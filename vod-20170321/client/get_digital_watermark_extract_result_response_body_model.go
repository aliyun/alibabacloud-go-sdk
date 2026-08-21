// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDigitalWatermarkExtractResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAiExtractResultList(v []*GetDigitalWatermarkExtractResultResponseBodyAiExtractResultList) *GetDigitalWatermarkExtractResultResponseBody
	GetAiExtractResultList() []*GetDigitalWatermarkExtractResultResponseBodyAiExtractResultList
	SetRequestId(v string) *GetDigitalWatermarkExtractResultResponseBody
	GetRequestId() *string
}

type GetDigitalWatermarkExtractResultResponseBody struct {
	// The details of the watermark extraction jobs.
	AiExtractResultList []*GetDigitalWatermarkExtractResultResponseBodyAiExtractResultList `json:"AiExtractResultList,omitempty" xml:"AiExtractResultList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 42E0554B-80F4-4921-****-ACFB22CAAAD0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDigitalWatermarkExtractResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDigitalWatermarkExtractResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetDigitalWatermarkExtractResultResponseBody) GetAiExtractResultList() []*GetDigitalWatermarkExtractResultResponseBodyAiExtractResultList {
	return s.AiExtractResultList
}

func (s *GetDigitalWatermarkExtractResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDigitalWatermarkExtractResultResponseBody) SetAiExtractResultList(v []*GetDigitalWatermarkExtractResultResponseBodyAiExtractResultList) *GetDigitalWatermarkExtractResultResponseBody {
	s.AiExtractResultList = v
	return s
}

func (s *GetDigitalWatermarkExtractResultResponseBody) SetRequestId(v string) *GetDigitalWatermarkExtractResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDigitalWatermarkExtractResultResponseBody) Validate() error {
	if s.AiExtractResultList != nil {
		for _, item := range s.AiExtractResultList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetDigitalWatermarkExtractResultResponseBodyAiExtractResultList struct {
	// The time when the watermark extraction job was created.
	//
	// example:
	//
	// 2023-09-16T02:49:04Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The error message. This parameter is returned only when the watermark extraction job fails (the value of `Status` is `Failed`).
	//
	// example:
	//
	// query copyright info err
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the watermark extraction job.
	//
	// example:
	//
	// 3af004763bcf459698860f4ede20****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The time when the watermark extraction job was last updated.
	//
	// example:
	//
	// 2023-09-17T06:20:45Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The status of the watermark extraction job. Valid values:
	//
	// - **Success**: The job succeeded.
	//
	// - **Failed**: The job failed.
	//
	// - **Processing**: The job is being processed.
	//
	// example:
	//
	// Failed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The extracted watermark text.
	//
	// example:
	//
	// test mark
	WaterMarkText *string `json:"WaterMarkText,omitempty" xml:"WaterMarkText,omitempty"`
}

func (s GetDigitalWatermarkExtractResultResponseBodyAiExtractResultList) String() string {
	return dara.Prettify(s)
}

func (s GetDigitalWatermarkExtractResultResponseBodyAiExtractResultList) GoString() string {
	return s.String()
}

func (s *GetDigitalWatermarkExtractResultResponseBodyAiExtractResultList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetDigitalWatermarkExtractResultResponseBodyAiExtractResultList) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetDigitalWatermarkExtractResultResponseBodyAiExtractResultList) GetJobId() *string {
	return s.JobId
}

func (s *GetDigitalWatermarkExtractResultResponseBodyAiExtractResultList) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *GetDigitalWatermarkExtractResultResponseBodyAiExtractResultList) GetStatus() *string {
	return s.Status
}

func (s *GetDigitalWatermarkExtractResultResponseBodyAiExtractResultList) GetWaterMarkText() *string {
	return s.WaterMarkText
}

func (s *GetDigitalWatermarkExtractResultResponseBodyAiExtractResultList) SetCreateTime(v string) *GetDigitalWatermarkExtractResultResponseBodyAiExtractResultList {
	s.CreateTime = &v
	return s
}

func (s *GetDigitalWatermarkExtractResultResponseBodyAiExtractResultList) SetErrorMessage(v string) *GetDigitalWatermarkExtractResultResponseBodyAiExtractResultList {
	s.ErrorMessage = &v
	return s
}

func (s *GetDigitalWatermarkExtractResultResponseBodyAiExtractResultList) SetJobId(v string) *GetDigitalWatermarkExtractResultResponseBodyAiExtractResultList {
	s.JobId = &v
	return s
}

func (s *GetDigitalWatermarkExtractResultResponseBodyAiExtractResultList) SetModifyTime(v string) *GetDigitalWatermarkExtractResultResponseBodyAiExtractResultList {
	s.ModifyTime = &v
	return s
}

func (s *GetDigitalWatermarkExtractResultResponseBodyAiExtractResultList) SetStatus(v string) *GetDigitalWatermarkExtractResultResponseBodyAiExtractResultList {
	s.Status = &v
	return s
}

func (s *GetDigitalWatermarkExtractResultResponseBodyAiExtractResultList) SetWaterMarkText(v string) *GetDigitalWatermarkExtractResultResponseBodyAiExtractResultList {
	s.WaterMarkText = &v
	return s
}

func (s *GetDigitalWatermarkExtractResultResponseBodyAiExtractResultList) Validate() error {
	return dara.Validate(s)
}
