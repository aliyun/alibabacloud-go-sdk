// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePipelineRunResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *CreatePipelineRunResponseBody
	GetId() *string
	SetRequestId(v string) *CreatePipelineRunResponseBody
	GetRequestId() *string
}

type CreatePipelineRunResponseBody struct {
	// The ID of the process.
	//
	// example:
	//
	// a7ef0634-20ec-4a7c-a214-54020f91XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7C352CB7-CD88-50CF-9D0D-E81BDF02XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePipelineRunResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePipelineRunResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePipelineRunResponseBody) GetId() *string {
	return s.Id
}

func (s *CreatePipelineRunResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePipelineRunResponseBody) SetId(v string) *CreatePipelineRunResponseBody {
	s.Id = &v
	return s
}

func (s *CreatePipelineRunResponseBody) SetRequestId(v string) *CreatePipelineRunResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePipelineRunResponseBody) Validate() error {
	return dara.Validate(s)
}
