// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBatchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *CreateBatchResponseBody
	GetId() *string
	SetRequestId(v string) *CreateBatchResponseBody
	GetRequestId() *string
}

type CreateBatchResponseBody struct {
	// The ID of the batch processing task.
	//
	// example:
	//
	// batch-4eb9223f-3e88-42d3-a578-3f2852******
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EC564A9A-BA5C-4499-A087-D9B9E76E*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateBatchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBatchResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBatchResponseBody) GetId() *string {
	return s.Id
}

func (s *CreateBatchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBatchResponseBody) SetId(v string) *CreateBatchResponseBody {
	s.Id = &v
	return s
}

func (s *CreateBatchResponseBody) SetRequestId(v string) *CreateBatchResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBatchResponseBody) Validate() error {
	return dara.Validate(s)
}
