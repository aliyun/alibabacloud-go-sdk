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
	SetSuccess(v bool) *UpdatePipelineResponseBody
	GetSuccess() *bool
}

type UpdatePipelineResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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

func (s *UpdatePipelineResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdatePipelineResponseBody) SetRequestId(v string) *UpdatePipelineResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePipelineResponseBody) SetSuccess(v bool) *UpdatePipelineResponseBody {
	s.Success = &v
	return s
}

func (s *UpdatePipelineResponseBody) Validate() error {
	return dara.Validate(s)
}
