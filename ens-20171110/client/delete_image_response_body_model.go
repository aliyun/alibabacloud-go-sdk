// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteImageResponseBody
	GetCode() *int32
	SetRequestId(v string) *DeleteImageResponseBody
	GetRequestId() *string
}

type DeleteImageResponseBody struct {
	// The returned service code. A value of 0 indicates that the operation was successful.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0D94920B-0349-5097-A57F-31876405F2E5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteImageResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteImageResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteImageResponseBody) SetCode(v int32) *DeleteImageResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteImageResponseBody) SetRequestId(v string) *DeleteImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteImageResponseBody) Validate() error {
	return dara.Validate(s)
}
