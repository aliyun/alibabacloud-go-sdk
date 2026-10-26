// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCopilotEmbedConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyCopilotEmbedConfigResponseBody
	GetRequestId() *string
	SetResult(v bool) *ModifyCopilotEmbedConfigResponseBody
	GetResult() *bool
	SetSuccess(v bool) *ModifyCopilotEmbedConfigResponseBody
	GetSuccess() *bool
}

type ModifyCopilotEmbedConfigResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 4BAA469******A9289FEE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returns the result of the API execution. Possible values:
	//
	// - true: Execution succeeded
	//
	// - false: Execution failed
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request was successful. Possible values:
	//
	// - true: The request was successful
	//
	// - false: The request failed
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyCopilotEmbedConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCopilotEmbedConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCopilotEmbedConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCopilotEmbedConfigResponseBody) GetResult() *bool {
	return s.Result
}

func (s *ModifyCopilotEmbedConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyCopilotEmbedConfigResponseBody) SetRequestId(v string) *ModifyCopilotEmbedConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCopilotEmbedConfigResponseBody) SetResult(v bool) *ModifyCopilotEmbedConfigResponseBody {
	s.Result = &v
	return s
}

func (s *ModifyCopilotEmbedConfigResponseBody) SetSuccess(v bool) *ModifyCopilotEmbedConfigResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyCopilotEmbedConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
