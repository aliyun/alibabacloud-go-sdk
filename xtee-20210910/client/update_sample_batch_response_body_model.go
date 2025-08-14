// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSampleBatchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateSampleBatchResponseBody
	GetRequestId() *string
	SetData(v bool) *UpdateSampleBatchResponseBody
	GetData() *bool
}

type UpdateSampleBatchResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned data.
	//
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
}

func (s UpdateSampleBatchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSampleBatchResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSampleBatchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSampleBatchResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateSampleBatchResponseBody) SetRequestId(v string) *UpdateSampleBatchResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSampleBatchResponseBody) SetData(v bool) *UpdateSampleBatchResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateSampleBatchResponseBody) Validate() error {
	return dara.Validate(s)
}
