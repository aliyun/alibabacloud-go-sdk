// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFinishCurrentProjectNodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *FinishCurrentProjectNodeResponseBody
	GetRequestId() *string
	SetResult(v bool) *FinishCurrentProjectNodeResponseBody
	GetResult() *bool
	SetSuccess(v bool) *FinishCurrentProjectNodeResponseBody
	GetSuccess() *bool
}

type FinishCurrentProjectNodeResponseBody struct {
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

func (s FinishCurrentProjectNodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FinishCurrentProjectNodeResponseBody) GoString() string {
	return s.String()
}

func (s *FinishCurrentProjectNodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FinishCurrentProjectNodeResponseBody) GetResult() *bool {
	return s.Result
}

func (s *FinishCurrentProjectNodeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FinishCurrentProjectNodeResponseBody) SetRequestId(v string) *FinishCurrentProjectNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *FinishCurrentProjectNodeResponseBody) SetResult(v bool) *FinishCurrentProjectNodeResponseBody {
	s.Result = &v
	return s
}

func (s *FinishCurrentProjectNodeResponseBody) SetSuccess(v bool) *FinishCurrentProjectNodeResponseBody {
	s.Success = &v
	return s
}

func (s *FinishCurrentProjectNodeResponseBody) Validate() error {
	return dara.Validate(s)
}
