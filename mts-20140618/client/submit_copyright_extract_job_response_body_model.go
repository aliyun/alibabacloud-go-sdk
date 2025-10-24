// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitCopyrightExtractJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SubmitCopyrightExtractJobResponseBodyData) *SubmitCopyrightExtractJobResponseBody
	GetData() *SubmitCopyrightExtractJobResponseBodyData
	SetMessage(v string) *SubmitCopyrightExtractJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitCopyrightExtractJobResponseBody
	GetRequestId() *string
	SetStatusCode(v int64) *SubmitCopyrightExtractJobResponseBody
	GetStatusCode() *int64
}

type SubmitCopyrightExtractJobResponseBody struct {
	Data *SubmitCopyrightExtractJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// ok
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 05F8B913-E9F3-4A6F-9922-48CADA0FFAAD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	StatusCode *int64 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s SubmitCopyrightExtractJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitCopyrightExtractJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitCopyrightExtractJobResponseBody) GetData() *SubmitCopyrightExtractJobResponseBodyData {
	return s.Data
}

func (s *SubmitCopyrightExtractJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitCopyrightExtractJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitCopyrightExtractJobResponseBody) GetStatusCode() *int64 {
	return s.StatusCode
}

func (s *SubmitCopyrightExtractJobResponseBody) SetData(v *SubmitCopyrightExtractJobResponseBodyData) *SubmitCopyrightExtractJobResponseBody {
	s.Data = v
	return s
}

func (s *SubmitCopyrightExtractJobResponseBody) SetMessage(v string) *SubmitCopyrightExtractJobResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitCopyrightExtractJobResponseBody) SetRequestId(v string) *SubmitCopyrightExtractJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitCopyrightExtractJobResponseBody) SetStatusCode(v int64) *SubmitCopyrightExtractJobResponseBody {
	s.StatusCode = &v
	return s
}

func (s *SubmitCopyrightExtractJobResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitCopyrightExtractJobResponseBodyData struct {
	// example:
	//
	// ebbfe90c63b54ed4b61acb2f6c44****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s SubmitCopyrightExtractJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitCopyrightExtractJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitCopyrightExtractJobResponseBodyData) GetJobId() *string {
	return s.JobId
}

func (s *SubmitCopyrightExtractJobResponseBodyData) SetJobId(v string) *SubmitCopyrightExtractJobResponseBodyData {
	s.JobId = &v
	return s
}

func (s *SubmitCopyrightExtractJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}
