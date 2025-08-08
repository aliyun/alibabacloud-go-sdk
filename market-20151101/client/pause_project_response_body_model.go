// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPauseProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PauseProjectResponseBody
	GetRequestId() *string
	SetResult(v bool) *PauseProjectResponseBody
	GetResult() *bool
	SetSuccess(v bool) *PauseProjectResponseBody
	GetSuccess() *bool
}

type PauseProjectResponseBody struct {
	// example:
	//
	// ee69a00f-189b-400f-9fd2-af89749fb50f
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PauseProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PauseProjectResponseBody) GoString() string {
	return s.String()
}

func (s *PauseProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PauseProjectResponseBody) GetResult() *bool {
	return s.Result
}

func (s *PauseProjectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PauseProjectResponseBody) SetRequestId(v string) *PauseProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *PauseProjectResponseBody) SetResult(v bool) *PauseProjectResponseBody {
	s.Result = &v
	return s
}

func (s *PauseProjectResponseBody) SetSuccess(v bool) *PauseProjectResponseBody {
	s.Success = &v
	return s
}

func (s *PauseProjectResponseBody) Validate() error {
	return dara.Validate(s)
}
