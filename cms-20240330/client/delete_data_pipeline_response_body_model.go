// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataPipelineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDataPipelineResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDataPipelineResponseBody
	GetSuccess() *bool
}

type DeleteDataPipelineResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The deletion result.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteDataPipelineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataPipelineResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataPipelineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDataPipelineResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDataPipelineResponseBody) SetRequestId(v string) *DeleteDataPipelineResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataPipelineResponseBody) SetSuccess(v bool) *DeleteDataPipelineResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDataPipelineResponseBody) Validate() error {
	return dara.Validate(s)
}
