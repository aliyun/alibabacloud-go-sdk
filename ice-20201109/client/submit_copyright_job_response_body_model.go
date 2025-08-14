// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitCopyrightJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SubmitCopyrightJobResponseBodyData) *SubmitCopyrightJobResponseBody
	GetData() *SubmitCopyrightJobResponseBodyData
	SetMessage(v string) *SubmitCopyrightJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitCopyrightJobResponseBody
	GetRequestId() *string
	SetStatusCode(v int64) *SubmitCopyrightJobResponseBody
	GetStatusCode() *int64
}

type SubmitCopyrightJobResponseBody struct {
	Data *SubmitCopyrightJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// ok
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// FA258E67-09B8-4EAA-8F33-BA567834A2C3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	StatusCode *int64 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s SubmitCopyrightJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitCopyrightJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitCopyrightJobResponseBody) GetData() *SubmitCopyrightJobResponseBodyData {
	return s.Data
}

func (s *SubmitCopyrightJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitCopyrightJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitCopyrightJobResponseBody) GetStatusCode() *int64 {
	return s.StatusCode
}

func (s *SubmitCopyrightJobResponseBody) SetData(v *SubmitCopyrightJobResponseBodyData) *SubmitCopyrightJobResponseBody {
	s.Data = v
	return s
}

func (s *SubmitCopyrightJobResponseBody) SetMessage(v string) *SubmitCopyrightJobResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitCopyrightJobResponseBody) SetRequestId(v string) *SubmitCopyrightJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitCopyrightJobResponseBody) SetStatusCode(v int64) *SubmitCopyrightJobResponseBody {
	s.StatusCode = &v
	return s
}

func (s *SubmitCopyrightJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type SubmitCopyrightJobResponseBodyData struct {
	// example:
	//
	// bfb786c63****f4d80648792021eff90
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s SubmitCopyrightJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitCopyrightJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitCopyrightJobResponseBodyData) GetJobId() *string {
	return s.JobId
}

func (s *SubmitCopyrightJobResponseBodyData) SetJobId(v string) *SubmitCopyrightJobResponseBodyData {
	s.JobId = &v
	return s
}

func (s *SubmitCopyrightJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}
