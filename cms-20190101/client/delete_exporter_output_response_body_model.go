// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExporterOutputResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteExporterOutputResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteExporterOutputResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteExporterOutputResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteExporterOutputResponseBody
	GetSuccess() *bool
}

type DeleteExporterOutputResponseBody struct {
	// The HTTP status code.
	//
	// > The status code 200 indicates that the request was successful. Other status codes indicate that the request failed.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2DECF751-7AFA-43BB-8C63-2B6B07E51AE1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteExporterOutputResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteExporterOutputResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteExporterOutputResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteExporterOutputResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteExporterOutputResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteExporterOutputResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteExporterOutputResponseBody) SetCode(v string) *DeleteExporterOutputResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteExporterOutputResponseBody) SetMessage(v string) *DeleteExporterOutputResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteExporterOutputResponseBody) SetRequestId(v string) *DeleteExporterOutputResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteExporterOutputResponseBody) SetSuccess(v bool) *DeleteExporterOutputResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteExporterOutputResponseBody) Validate() error {
	return dara.Validate(s)
}
