// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAbolishPipelineRunResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AbolishPipelineRunResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AbolishPipelineRunResponseBody
	GetSuccess() *bool
}

type AbolishPipelineRunResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 55D786C9-DD57-524D-884C-C5239278XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AbolishPipelineRunResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AbolishPipelineRunResponseBody) GoString() string {
	return s.String()
}

func (s *AbolishPipelineRunResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AbolishPipelineRunResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AbolishPipelineRunResponseBody) SetRequestId(v string) *AbolishPipelineRunResponseBody {
	s.RequestId = &v
	return s
}

func (s *AbolishPipelineRunResponseBody) SetSuccess(v bool) *AbolishPipelineRunResponseBody {
	s.Success = &v
	return s
}

func (s *AbolishPipelineRunResponseBody) Validate() error {
	return dara.Validate(s)
}
