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
	// ****2876-6263-4B75-8F2C-CD0F7FCF****
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
	return dara.Validate(s)
}

type SubmitCopyrightExtractJobResponseBodyData struct {
	// example:
	//
	// bfb786c63****4d80648792021eff90
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
