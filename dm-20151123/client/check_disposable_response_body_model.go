// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDisposableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsDisposable(v string) *CheckDisposableResponseBody
	GetIsDisposable() *string
	SetIsFormatValid(v string) *CheckDisposableResponseBody
	GetIsFormatValid() *string
	SetIsMxValid(v string) *CheckDisposableResponseBody
	GetIsMxValid() *string
	SetIsOkToSend(v string) *CheckDisposableResponseBody
	GetIsOkToSend() *string
	SetRequestId(v string) *CheckDisposableResponseBody
	GetRequestId() *string
}

type CheckDisposableResponseBody struct {
	// example:
	//
	// true
	IsDisposable *string `json:"IsDisposable,omitempty" xml:"IsDisposable,omitempty"`
	// example:
	//
	// true
	IsFormatValid *string `json:"IsFormatValid,omitempty" xml:"IsFormatValid,omitempty"`
	// example:
	//
	// true
	IsMxValid *string `json:"IsMxValid,omitempty" xml:"IsMxValid,omitempty"`
	// example:
	//
	// false
	IsOkToSend *string `json:"IsOkToSend,omitempty" xml:"IsOkToSend,omitempty"`
	// example:
	//
	// xxx-xxx-xxxxx-xxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckDisposableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckDisposableResponseBody) GoString() string {
	return s.String()
}

func (s *CheckDisposableResponseBody) GetIsDisposable() *string {
	return s.IsDisposable
}

func (s *CheckDisposableResponseBody) GetIsFormatValid() *string {
	return s.IsFormatValid
}

func (s *CheckDisposableResponseBody) GetIsMxValid() *string {
	return s.IsMxValid
}

func (s *CheckDisposableResponseBody) GetIsOkToSend() *string {
	return s.IsOkToSend
}

func (s *CheckDisposableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckDisposableResponseBody) SetIsDisposable(v string) *CheckDisposableResponseBody {
	s.IsDisposable = &v
	return s
}

func (s *CheckDisposableResponseBody) SetIsFormatValid(v string) *CheckDisposableResponseBody {
	s.IsFormatValid = &v
	return s
}

func (s *CheckDisposableResponseBody) SetIsMxValid(v string) *CheckDisposableResponseBody {
	s.IsMxValid = &v
	return s
}

func (s *CheckDisposableResponseBody) SetIsOkToSend(v string) *CheckDisposableResponseBody {
	s.IsOkToSend = &v
	return s
}

func (s *CheckDisposableResponseBody) SetRequestId(v string) *CheckDisposableResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckDisposableResponseBody) Validate() error {
	return dara.Validate(s)
}
