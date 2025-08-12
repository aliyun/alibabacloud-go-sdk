// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutExporterOutputResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PutExporterOutputResponseBody
	GetCode() *string
	SetMessage(v string) *PutExporterOutputResponseBody
	GetMessage() *string
	SetRequestId(v string) *PutExporterOutputResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PutExporterOutputResponseBody
	GetSuccess() *bool
}

type PutExporterOutputResponseBody struct {
	// The HTTP status code.
	//
	// > The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// The Request is not authorization.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6A5F022D-AC7C-460E-94AE-B9E75083D027
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. true: The request was successful. false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PutExporterOutputResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PutExporterOutputResponseBody) GoString() string {
	return s.String()
}

func (s *PutExporterOutputResponseBody) GetCode() *string {
	return s.Code
}

func (s *PutExporterOutputResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PutExporterOutputResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PutExporterOutputResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PutExporterOutputResponseBody) SetCode(v string) *PutExporterOutputResponseBody {
	s.Code = &v
	return s
}

func (s *PutExporterOutputResponseBody) SetMessage(v string) *PutExporterOutputResponseBody {
	s.Message = &v
	return s
}

func (s *PutExporterOutputResponseBody) SetRequestId(v string) *PutExporterOutputResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutExporterOutputResponseBody) SetSuccess(v bool) *PutExporterOutputResponseBody {
	s.Success = &v
	return s
}

func (s *PutExporterOutputResponseBody) Validate() error {
	return dara.Validate(s)
}
