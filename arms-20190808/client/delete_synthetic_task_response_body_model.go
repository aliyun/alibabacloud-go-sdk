// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSyntheticTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSyntheticTaskResponseBody
	GetRequestId() *string
	SetResult(v string) *DeleteSyntheticTaskResponseBody
	GetResult() *string
}

type DeleteSyntheticTaskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1A9C645C-C83F-4C9D-8CCB-29BEC9E1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the specified tasks are deleted.
	//
	// 	- `true`: The tasks are deleted.
	//
	// 	- `false`: The tasks fail to be deleted.
	//
	// example:
	//
	// true
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DeleteSyntheticTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSyntheticTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSyntheticTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSyntheticTaskResponseBody) GetResult() *string {
	return s.Result
}

func (s *DeleteSyntheticTaskResponseBody) SetRequestId(v string) *DeleteSyntheticTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSyntheticTaskResponseBody) SetResult(v string) *DeleteSyntheticTaskResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteSyntheticTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
