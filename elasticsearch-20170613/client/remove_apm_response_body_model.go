// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveApmResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveApmResponseBody
	GetRequestId() *string
	SetResult(v bool) *RemoveApmResponseBody
	GetResult() *bool
}

type RemoveApmResponseBody struct {
	// example:
	//
	// 29A879FB-86BF-54CA-9426-B769A099E1A1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s RemoveApmResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveApmResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveApmResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveApmResponseBody) GetResult() *bool {
	return s.Result
}

func (s *RemoveApmResponseBody) SetRequestId(v string) *RemoveApmResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveApmResponseBody) SetResult(v bool) *RemoveApmResponseBody {
	s.Result = &v
	return s
}

func (s *RemoveApmResponseBody) Validate() error {
	return dara.Validate(s)
}
