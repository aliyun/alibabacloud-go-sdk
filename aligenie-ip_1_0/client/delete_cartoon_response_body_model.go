// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCartoonResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *DeleteCartoonResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteCartoonResponseBody
	GetRequestId() *string
	SetResult(v bool) *DeleteCartoonResponseBody
	GetResult() *bool
	SetStatusCode(v int32) *DeleteCartoonResponseBody
	GetStatusCode() *int32
}

type DeleteCartoonResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0EC7*726E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s DeleteCartoonResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCartoonResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCartoonResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteCartoonResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCartoonResponseBody) GetResult() *bool {
	return s.Result
}

func (s *DeleteCartoonResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCartoonResponseBody) SetMessage(v string) *DeleteCartoonResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteCartoonResponseBody) SetRequestId(v string) *DeleteCartoonResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCartoonResponseBody) SetResult(v bool) *DeleteCartoonResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteCartoonResponseBody) SetStatusCode(v int32) *DeleteCartoonResponseBody {
	s.StatusCode = &v
	return s
}

func (s *DeleteCartoonResponseBody) Validate() error {
	return dara.Validate(s)
}
