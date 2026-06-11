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
}

type DeletePipelineResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CD8BA7D6-995D-578D-9941-78B0FECD14B5
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
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

func (s *DeletePipelineResponseBody) SetRequestId(v string) *DeletePipelineResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePipelineResponseBody) Validate() error {
	return dara.Validate(s)
}
