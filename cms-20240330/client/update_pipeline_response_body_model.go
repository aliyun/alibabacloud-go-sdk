// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePipelineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdatePipelineResponseBody
	GetRequestId() *string
}

type UpdatePipelineResponseBody struct {
	// The unique ID of the request.
	//
	// example:
	//
	// E99F1CCD-256A-5DF9-9B67-8F4A7ACE7132
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdatePipelineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePipelineResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePipelineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePipelineResponseBody) SetRequestId(v string) *UpdatePipelineResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePipelineResponseBody) Validate() error {
	return dara.Validate(s)
}
