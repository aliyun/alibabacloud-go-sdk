// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDigitalWatermarkExtractJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *SubmitDigitalWatermarkExtractJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *SubmitDigitalWatermarkExtractJobResponseBody
	GetRequestId() *string
}

type SubmitDigitalWatermarkExtractJobResponseBody struct {
	// The ID of the watermark extraction job.
	//
	// example:
	//
	// ad90a501b1b9472374ad005046****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-****-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitDigitalWatermarkExtractJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitDigitalWatermarkExtractJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitDigitalWatermarkExtractJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitDigitalWatermarkExtractJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitDigitalWatermarkExtractJobResponseBody) SetJobId(v string) *SubmitDigitalWatermarkExtractJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitDigitalWatermarkExtractJobResponseBody) SetRequestId(v string) *SubmitDigitalWatermarkExtractJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitDigitalWatermarkExtractJobResponseBody) Validate() error {
	return dara.Validate(s)
}
