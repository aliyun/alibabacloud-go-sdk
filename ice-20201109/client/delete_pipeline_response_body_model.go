// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePipelineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeletePipelineResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeletePipelineResponseBody
	GetSuccess() *bool
}

type DeletePipelineResponseBody struct {
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

func (s DeletePipelineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePipelineResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePipelineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePipelineResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeletePipelineResponseBody) SetRequestId(v string) *DeletePipelineResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePipelineResponseBody) SetSuccess(v bool) *DeletePipelineResponseBody {
	s.Success = &v
	return s
}

func (s *DeletePipelineResponseBody) Validate() error {
	return dara.Validate(s)
}
