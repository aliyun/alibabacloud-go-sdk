// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePipelineManagementConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdatePipelineManagementConfigResponseBody
	GetRequestId() *string
	SetResult(v bool) *UpdatePipelineManagementConfigResponseBody
	GetResult() *bool
}

type UpdatePipelineManagementConfigResponseBody struct {
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UpdatePipelineManagementConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePipelineManagementConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePipelineManagementConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePipelineManagementConfigResponseBody) GetResult() *bool {
	return s.Result
}

func (s *UpdatePipelineManagementConfigResponseBody) SetRequestId(v string) *UpdatePipelineManagementConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePipelineManagementConfigResponseBody) SetResult(v bool) *UpdatePipelineManagementConfigResponseBody {
	s.Result = &v
	return s
}

func (s *UpdatePipelineManagementConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
