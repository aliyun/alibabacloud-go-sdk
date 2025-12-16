// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveQueryProcessorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveQueryProcessorResponseBody
	GetRequestId() *string
	SetResult(v string) *RemoveQueryProcessorResponseBody
	GetResult() *string
}

type RemoveQueryProcessorResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0A6EB64B-B4C8-CF02-810F-E660812972FF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// \\--
	//
	// example:
	//
	// []
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s RemoveQueryProcessorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveQueryProcessorResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveQueryProcessorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveQueryProcessorResponseBody) GetResult() *string {
	return s.Result
}

func (s *RemoveQueryProcessorResponseBody) SetRequestId(v string) *RemoveQueryProcessorResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveQueryProcessorResponseBody) SetResult(v string) *RemoveQueryProcessorResponseBody {
	s.Result = &v
	return s
}

func (s *RemoveQueryProcessorResponseBody) Validate() error {
	return dara.Validate(s)
}
