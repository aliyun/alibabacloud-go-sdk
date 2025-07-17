// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUndeployHttpApiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UndeployHttpApiResponseBody
	GetCode() *string
	SetMessage(v string) *UndeployHttpApiResponseBody
	GetMessage() *string
	SetRequestId(v string) *UndeployHttpApiResponseBody
	GetRequestId() *string
}

type UndeployHttpApiResponseBody struct {
	// Response code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 3ACFC7A7-45A9-58CF-B2D5-765B60254695
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UndeployHttpApiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UndeployHttpApiResponseBody) GoString() string {
	return s.String()
}

func (s *UndeployHttpApiResponseBody) GetCode() *string {
	return s.Code
}

func (s *UndeployHttpApiResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UndeployHttpApiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UndeployHttpApiResponseBody) SetCode(v string) *UndeployHttpApiResponseBody {
	s.Code = &v
	return s
}

func (s *UndeployHttpApiResponseBody) SetMessage(v string) *UndeployHttpApiResponseBody {
	s.Message = &v
	return s
}

func (s *UndeployHttpApiResponseBody) SetRequestId(v string) *UndeployHttpApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *UndeployHttpApiResponseBody) Validate() error {
	return dara.Validate(s)
}
