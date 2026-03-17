// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProbeTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteProbeTaskResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteProbeTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteProbeTaskResponseBody
	GetRequestId() *string
}

type DeleteProbeTaskResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 324223F3-93D3-4CE4-B26F-66C0C3809922
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteProbeTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteProbeTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProbeTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteProbeTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteProbeTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteProbeTaskResponseBody) SetCode(v string) *DeleteProbeTaskResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteProbeTaskResponseBody) SetMessage(v string) *DeleteProbeTaskResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteProbeTaskResponseBody) SetRequestId(v string) *DeleteProbeTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteProbeTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
