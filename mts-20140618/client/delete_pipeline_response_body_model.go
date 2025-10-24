// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePipelineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPipelineId(v string) *DeletePipelineResponseBody
	GetPipelineId() *string
	SetRequestId(v string) *DeletePipelineResponseBody
	GetRequestId() *string
}

type DeletePipelineResponseBody struct {
	// The ID of the MPS queue that is deleted.
	//
	// example:
	//
	// d1ce4d3efcb549419193f50f1fcd****
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 338CA33A-AE83-5DF4-B6F2-C6D3ED8143F5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePipelineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePipelineResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePipelineResponseBody) GetPipelineId() *string {
	return s.PipelineId
}

func (s *DeletePipelineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePipelineResponseBody) SetPipelineId(v string) *DeletePipelineResponseBody {
	s.PipelineId = &v
	return s
}

func (s *DeletePipelineResponseBody) SetRequestId(v string) *DeletePipelineResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePipelineResponseBody) Validate() error {
	return dara.Validate(s)
}
