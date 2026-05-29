// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAutomaticWriteOffResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AutomaticWriteOffResponseBody
	GetCode() *string
	SetData(v bool) *AutomaticWriteOffResponseBody
	GetData() *bool
	SetMessage(v string) *AutomaticWriteOffResponseBody
	GetMessage() *string
	SetRequestId(v string) *AutomaticWriteOffResponseBody
	GetRequestId() *string
}

type AutomaticWriteOffResponseBody struct {
	// code
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data    *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 6fc1309b17543600398356606d0096
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AutomaticWriteOffResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AutomaticWriteOffResponseBody) GoString() string {
	return s.String()
}

func (s *AutomaticWriteOffResponseBody) GetCode() *string {
	return s.Code
}

func (s *AutomaticWriteOffResponseBody) GetData() *bool {
	return s.Data
}

func (s *AutomaticWriteOffResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AutomaticWriteOffResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AutomaticWriteOffResponseBody) SetCode(v string) *AutomaticWriteOffResponseBody {
	s.Code = &v
	return s
}

func (s *AutomaticWriteOffResponseBody) SetData(v bool) *AutomaticWriteOffResponseBody {
	s.Data = &v
	return s
}

func (s *AutomaticWriteOffResponseBody) SetMessage(v string) *AutomaticWriteOffResponseBody {
	s.Message = &v
	return s
}

func (s *AutomaticWriteOffResponseBody) SetRequestId(v string) *AutomaticWriteOffResponseBody {
	s.RequestId = &v
	return s
}

func (s *AutomaticWriteOffResponseBody) Validate() error {
	return dara.Validate(s)
}
