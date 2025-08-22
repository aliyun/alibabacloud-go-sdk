// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteAppResponseBody
	GetCode() *int32
	SetMessage(v string) *DeleteAppResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteAppResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteAppResponseBody
	GetSuccess() *bool
}

type DeleteAppResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// CF99C381-C8F6-5A8D-8C24-57F46B706D2D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteAppResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAppResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteAppResponseBody) SetCode(v int32) *DeleteAppResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAppResponseBody) SetMessage(v string) *DeleteAppResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAppResponseBody) SetRequestId(v string) *DeleteAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAppResponseBody) SetSuccess(v bool) *DeleteAppResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteAppResponseBody) Validate() error {
	return dara.Validate(s)
}
