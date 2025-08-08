// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackCurrentProjectNodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RollbackCurrentProjectNodeResponseBody
	GetRequestId() *string
	SetResult(v bool) *RollbackCurrentProjectNodeResponseBody
	GetResult() *bool
	SetSuccess(v bool) *RollbackCurrentProjectNodeResponseBody
	GetSuccess() *bool
}

type RollbackCurrentProjectNodeResponseBody struct {
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

func (s RollbackCurrentProjectNodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RollbackCurrentProjectNodeResponseBody) GoString() string {
	return s.String()
}

func (s *RollbackCurrentProjectNodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RollbackCurrentProjectNodeResponseBody) GetResult() *bool {
	return s.Result
}

func (s *RollbackCurrentProjectNodeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RollbackCurrentProjectNodeResponseBody) SetRequestId(v string) *RollbackCurrentProjectNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *RollbackCurrentProjectNodeResponseBody) SetResult(v bool) *RollbackCurrentProjectNodeResponseBody {
	s.Result = &v
	return s
}

func (s *RollbackCurrentProjectNodeResponseBody) SetSuccess(v bool) *RollbackCurrentProjectNodeResponseBody {
	s.Success = &v
	return s
}

func (s *RollbackCurrentProjectNodeResponseBody) Validate() error {
	return dara.Validate(s)
}
