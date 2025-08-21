// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopPublishStreamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *StopPublishStreamResponseBody
	GetCode() *int64
	SetMessage(v int64) *StopPublishStreamResponseBody
	GetMessage() *int64
	SetRequestId(v string) *StopPublishStreamResponseBody
	GetRequestId() *string
}

type StopPublishStreamResponseBody struct {
	Code      *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *int64  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopPublishStreamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopPublishStreamResponseBody) GoString() string {
	return s.String()
}

func (s *StopPublishStreamResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *StopPublishStreamResponseBody) GetMessage() *int64 {
	return s.Message
}

func (s *StopPublishStreamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopPublishStreamResponseBody) SetCode(v int64) *StopPublishStreamResponseBody {
	s.Code = &v
	return s
}

func (s *StopPublishStreamResponseBody) SetMessage(v int64) *StopPublishStreamResponseBody {
	s.Message = &v
	return s
}

func (s *StopPublishStreamResponseBody) SetRequestId(v string) *StopPublishStreamResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopPublishStreamResponseBody) Validate() error {
	return dara.Validate(s)
}
