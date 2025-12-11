// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteSourceResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteSourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteSourceResponseBody
	GetRequestId() *string
}

type DeleteSourceResponseBody struct {
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// C67DED2B-F19B-5BEC-88C1-D6EB854CD0D4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteSourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSourceResponseBody) SetCode(v string) *DeleteSourceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSourceResponseBody) SetMessage(v string) *DeleteSourceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSourceResponseBody) SetRequestId(v string) *DeleteSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
