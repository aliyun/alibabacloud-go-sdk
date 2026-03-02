// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRepoForkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateRepoForkResponseBody
	GetRequestId() *string
	SetResult(v int64) *CreateRepoForkResponseBody
	GetResult() *int64
}

type CreateRepoForkResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *int64  `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CreateRepoForkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRepoForkResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRepoForkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRepoForkResponseBody) GetResult() *int64 {
	return s.Result
}

func (s *CreateRepoForkResponseBody) SetRequestId(v string) *CreateRepoForkResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRepoForkResponseBody) SetResult(v int64) *CreateRepoForkResponseBody {
	s.Result = &v
	return s
}

func (s *CreateRepoForkResponseBody) Validate() error {
	return dara.Validate(s)
}
