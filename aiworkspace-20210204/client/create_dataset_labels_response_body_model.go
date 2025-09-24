// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatasetLabelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateDatasetLabelsResponseBody
	GetRequestId() *string
}

type CreateDatasetLabelsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A083731B-4973-54D1-B324-E53****4DD44
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDatasetLabelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetLabelsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDatasetLabelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDatasetLabelsResponseBody) SetRequestId(v string) *CreateDatasetLabelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDatasetLabelsResponseBody) Validate() error {
	return dara.Validate(s)
}
