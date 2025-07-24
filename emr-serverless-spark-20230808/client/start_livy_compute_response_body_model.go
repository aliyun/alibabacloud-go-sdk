// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartLivyComputeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StartLivyComputeResponseBody
	GetCode() *string
	SetMessage(v string) *StartLivyComputeResponseBody
	GetMessage() *string
	SetRequestId(v string) *StartLivyComputeResponseBody
	GetRequestId() *string
}

type StartLivyComputeResponseBody struct {
	// example:
	//
	// 1000000
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s StartLivyComputeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartLivyComputeResponseBody) GoString() string {
	return s.String()
}

func (s *StartLivyComputeResponseBody) GetCode() *string {
	return s.Code
}

func (s *StartLivyComputeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StartLivyComputeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartLivyComputeResponseBody) SetCode(v string) *StartLivyComputeResponseBody {
	s.Code = &v
	return s
}

func (s *StartLivyComputeResponseBody) SetMessage(v string) *StartLivyComputeResponseBody {
	s.Message = &v
	return s
}

func (s *StartLivyComputeResponseBody) SetRequestId(v string) *StartLivyComputeResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartLivyComputeResponseBody) Validate() error {
	return dara.Validate(s)
}
