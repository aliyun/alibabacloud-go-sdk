// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSummariesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateSummariesResponseBody
	GetRequestId() *string
	SetResult(v bool) *UpdateSummariesResponseBody
	GetResult() *bool
}

type UpdateSummariesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 7A389E09-7964-5A2B-FE9D-F6CFA7162852
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateSummariesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSummariesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSummariesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSummariesResponseBody) GetResult() *bool {
	return s.Result
}

func (s *UpdateSummariesResponseBody) SetRequestId(v string) *UpdateSummariesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSummariesResponseBody) SetResult(v bool) *UpdateSummariesResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateSummariesResponseBody) Validate() error {
	return dara.Validate(s)
}
