// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMessageAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteMessageAppResponseBody
	GetRequestId() *string
	SetResult(v *DeleteMessageAppResponseBodyResult) *DeleteMessageAppResponseBody
	GetResult() *DeleteMessageAppResponseBodyResult
}

type DeleteMessageAppResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 16A96B9A-****-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	Result *DeleteMessageAppResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DeleteMessageAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMessageAppResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMessageAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMessageAppResponseBody) GetResult() *DeleteMessageAppResponseBodyResult {
	return s.Result
}

func (s *DeleteMessageAppResponseBody) SetRequestId(v string) *DeleteMessageAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMessageAppResponseBody) SetResult(v *DeleteMessageAppResponseBodyResult) *DeleteMessageAppResponseBody {
	s.Result = v
	return s
}

func (s *DeleteMessageAppResponseBody) Validate() error {
	return dara.Validate(s)
}

type DeleteMessageAppResponseBodyResult struct {
	// Indicates whether the application was deleted. Valid values:
	//
	// 	- true: The application was deleted.
	//
	// 	- false: The application failed to be deleted.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteMessageAppResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DeleteMessageAppResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteMessageAppResponseBodyResult) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteMessageAppResponseBodyResult) SetSuccess(v bool) *DeleteMessageAppResponseBodyResult {
	s.Success = &v
	return s
}

func (s *DeleteMessageAppResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
