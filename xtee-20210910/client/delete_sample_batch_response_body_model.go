// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSampleBatchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSampleBatchResponseBody
	GetRequestId() *string
	SetData(v bool) *DeleteSampleBatchResponseBody
	GetData() *bool
}

type DeleteSampleBatchResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned data object.
	//
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
}

func (s DeleteSampleBatchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSampleBatchResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSampleBatchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSampleBatchResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteSampleBatchResponseBody) SetRequestId(v string) *DeleteSampleBatchResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSampleBatchResponseBody) SetData(v bool) *DeleteSampleBatchResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteSampleBatchResponseBody) Validate() error {
	return dara.Validate(s)
}
