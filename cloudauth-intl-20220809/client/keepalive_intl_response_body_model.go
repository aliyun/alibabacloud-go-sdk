// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKeepaliveIntlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *KeepaliveIntlResponseBody
	GetCode() *string
	SetMessage(v string) *KeepaliveIntlResponseBody
	GetMessage() *string
	SetRequestId(v string) *KeepaliveIntlResponseBody
	GetRequestId() *string
	SetResult(v *KeepaliveIntlResponseBodyResult) *KeepaliveIntlResponseBody
	GetResult() *KeepaliveIntlResponseBodyResult
}

type KeepaliveIntlResponseBody struct {
	// The return code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The return message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4EB35****87EBA1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	Result *KeepaliveIntlResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s KeepaliveIntlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s KeepaliveIntlResponseBody) GoString() string {
	return s.String()
}

func (s *KeepaliveIntlResponseBody) GetCode() *string {
	return s.Code
}

func (s *KeepaliveIntlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *KeepaliveIntlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *KeepaliveIntlResponseBody) GetResult() *KeepaliveIntlResponseBodyResult {
	return s.Result
}

func (s *KeepaliveIntlResponseBody) SetCode(v string) *KeepaliveIntlResponseBody {
	s.Code = &v
	return s
}

func (s *KeepaliveIntlResponseBody) SetMessage(v string) *KeepaliveIntlResponseBody {
	s.Message = &v
	return s
}

func (s *KeepaliveIntlResponseBody) SetRequestId(v string) *KeepaliveIntlResponseBody {
	s.RequestId = &v
	return s
}

func (s *KeepaliveIntlResponseBody) SetResult(v *KeepaliveIntlResponseBodyResult) *KeepaliveIntlResponseBody {
	s.Result = v
	return s
}

func (s *KeepaliveIntlResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type KeepaliveIntlResponseBodyResult struct {
	// The returned result.
	//
	// example:
	//
	// SUCCESS
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s KeepaliveIntlResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s KeepaliveIntlResponseBodyResult) GoString() string {
	return s.String()
}

func (s *KeepaliveIntlResponseBodyResult) GetResult() *string {
	return s.Result
}

func (s *KeepaliveIntlResponseBodyResult) SetResult(v string) *KeepaliveIntlResponseBodyResult {
	s.Result = &v
	return s
}

func (s *KeepaliveIntlResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
