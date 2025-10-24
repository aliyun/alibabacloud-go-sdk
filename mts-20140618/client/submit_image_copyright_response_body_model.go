// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitImageCopyrightResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SubmitImageCopyrightResponseBodyData) *SubmitImageCopyrightResponseBody
	GetData() *SubmitImageCopyrightResponseBodyData
	SetMessage(v string) *SubmitImageCopyrightResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitImageCopyrightResponseBody
	GetRequestId() *string
	SetStatusCode(v int64) *SubmitImageCopyrightResponseBody
	GetStatusCode() *int64
}

type SubmitImageCopyrightResponseBody struct {
	Data *SubmitImageCopyrightResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// ok
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// D1D5C080-8E2F-5030-8AB4-13092F17631B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	StatusCode *int64 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s SubmitImageCopyrightResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitImageCopyrightResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitImageCopyrightResponseBody) GetData() *SubmitImageCopyrightResponseBodyData {
	return s.Data
}

func (s *SubmitImageCopyrightResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitImageCopyrightResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitImageCopyrightResponseBody) GetStatusCode() *int64 {
	return s.StatusCode
}

func (s *SubmitImageCopyrightResponseBody) SetData(v *SubmitImageCopyrightResponseBodyData) *SubmitImageCopyrightResponseBody {
	s.Data = v
	return s
}

func (s *SubmitImageCopyrightResponseBody) SetMessage(v string) *SubmitImageCopyrightResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitImageCopyrightResponseBody) SetRequestId(v string) *SubmitImageCopyrightResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitImageCopyrightResponseBody) SetStatusCode(v int64) *SubmitImageCopyrightResponseBody {
	s.StatusCode = &v
	return s
}

func (s *SubmitImageCopyrightResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitImageCopyrightResponseBodyData struct {
	// example:
	//
	// bfb786c639894f4d80648792021e****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s SubmitImageCopyrightResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitImageCopyrightResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitImageCopyrightResponseBodyData) GetJobId() *string {
	return s.JobId
}

func (s *SubmitImageCopyrightResponseBodyData) SetJobId(v string) *SubmitImageCopyrightResponseBodyData {
	s.JobId = &v
	return s
}

func (s *SubmitImageCopyrightResponseBodyData) Validate() error {
	return dara.Validate(s)
}
