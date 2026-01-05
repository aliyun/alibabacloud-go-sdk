// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatasetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *CreateDatasetResponseBody
	GetId() *string
	SetRequestId(v string) *CreateDatasetResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDatasetResponseBody
	GetSuccess() *bool
}

type CreateDatasetResponseBody struct {
	// ID
	//
	// example:
	//
	// dataworks-dataset:3pXXXb8o0ngr07njhps1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 99EBE7CF-69C0-5089-BE3E-79563C31XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the creation was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDatasetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDatasetResponseBody) GetId() *string {
	return s.Id
}

func (s *CreateDatasetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDatasetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDatasetResponseBody) SetId(v string) *CreateDatasetResponseBody {
	s.Id = &v
	return s
}

func (s *CreateDatasetResponseBody) SetRequestId(v string) *CreateDatasetResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDatasetResponseBody) SetSuccess(v bool) *CreateDatasetResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDatasetResponseBody) Validate() error {
	return dara.Validate(s)
}
