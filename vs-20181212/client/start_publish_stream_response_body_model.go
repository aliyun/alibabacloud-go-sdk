// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartPublishStreamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *StartPublishStreamResponseBody
	GetCode() *int64
	SetMessage(v int64) *StartPublishStreamResponseBody
	GetMessage() *int64
	SetRequestId(v string) *StartPublishStreamResponseBody
	GetRequestId() *string
}

type StartPublishStreamResponseBody struct {
	// The return code.
	//
	// example:
	//
	// 0
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The return message.
	//
	// example:
	//
	// 0
	Message *int64 `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartPublishStreamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartPublishStreamResponseBody) GoString() string {
	return s.String()
}

func (s *StartPublishStreamResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *StartPublishStreamResponseBody) GetMessage() *int64 {
	return s.Message
}

func (s *StartPublishStreamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartPublishStreamResponseBody) SetCode(v int64) *StartPublishStreamResponseBody {
	s.Code = &v
	return s
}

func (s *StartPublishStreamResponseBody) SetMessage(v int64) *StartPublishStreamResponseBody {
	s.Message = &v
	return s
}

func (s *StartPublishStreamResponseBody) SetRequestId(v string) *StartPublishStreamResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartPublishStreamResponseBody) Validate() error {
	return dara.Validate(s)
}
