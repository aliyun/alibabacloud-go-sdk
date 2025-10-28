// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLogPathResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *AddLogPathResponseBody
	GetCode() *int32
	SetMessage(v string) *AddLogPathResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddLogPathResponseBody
	GetRequestId() *string
}

type AddLogPathResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
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

func (s AddLogPathResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddLogPathResponseBody) GoString() string {
	return s.String()
}

func (s *AddLogPathResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *AddLogPathResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddLogPathResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddLogPathResponseBody) SetCode(v int32) *AddLogPathResponseBody {
	s.Code = &v
	return s
}

func (s *AddLogPathResponseBody) SetMessage(v string) *AddLogPathResponseBody {
	s.Message = &v
	return s
}

func (s *AddLogPathResponseBody) SetRequestId(v string) *AddLogPathResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddLogPathResponseBody) Validate() error {
	return dara.Validate(s)
}
