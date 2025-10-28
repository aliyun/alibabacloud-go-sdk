// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLogPathResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteLogPathResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteLogPathResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteLogPathResponseBody
	GetRequestId() *string
}

type DeleteLogPathResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The additional information that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 3616cdca-4f92-4413**********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLogPathResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLogPathResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLogPathResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteLogPathResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteLogPathResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLogPathResponseBody) SetCode(v string) *DeleteLogPathResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteLogPathResponseBody) SetMessage(v string) *DeleteLogPathResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteLogPathResponseBody) SetRequestId(v string) *DeleteLogPathResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLogPathResponseBody) Validate() error {
	return dara.Validate(s)
}
