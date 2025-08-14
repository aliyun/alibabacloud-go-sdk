// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSampleBatchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSampleBatchResponseBody
	GetRequestId() *string
}

type CreateSampleBatchResponseBody struct {
	// ID of the request
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSampleBatchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSampleBatchResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSampleBatchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSampleBatchResponseBody) SetRequestId(v string) *CreateSampleBatchResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSampleBatchResponseBody) Validate() error {
	return dara.Validate(s)
}
