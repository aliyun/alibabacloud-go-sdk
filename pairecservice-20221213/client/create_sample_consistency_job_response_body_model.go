// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSampleConsistencyJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSampleConsistencyJobResponseBody
	GetRequestId() *string
	SetSampleConsistencyJobId(v string) *CreateSampleConsistencyJobResponseBody
	GetSampleConsistencyJobId() *string
}

type CreateSampleConsistencyJobResponseBody struct {
	RequestId              *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SampleConsistencyJobId *string `json:"SampleConsistencyJobId,omitempty" xml:"SampleConsistencyJobId,omitempty"`
}

func (s CreateSampleConsistencyJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSampleConsistencyJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSampleConsistencyJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSampleConsistencyJobResponseBody) GetSampleConsistencyJobId() *string {
	return s.SampleConsistencyJobId
}

func (s *CreateSampleConsistencyJobResponseBody) SetRequestId(v string) *CreateSampleConsistencyJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSampleConsistencyJobResponseBody) SetSampleConsistencyJobId(v string) *CreateSampleConsistencyJobResponseBody {
	s.SampleConsistencyJobId = &v
	return s
}

func (s *CreateSampleConsistencyJobResponseBody) Validate() error {
	return dara.Validate(s)
}
