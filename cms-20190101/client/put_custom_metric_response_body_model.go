// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutCustomMetricResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PutCustomMetricResponseBody
	GetCode() *string
	SetMessage(v string) *PutCustomMetricResponseBody
	GetMessage() *string
	SetRequestId(v string) *PutCustomMetricResponseBody
	GetRequestId() *string
}

type PutCustomMetricResponseBody struct {
	// The status code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// The request has failed due to a temporary failure of the server.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 05B36C2C-5F6E-48D5-8B41-CE36DD7EE8E0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PutCustomMetricResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PutCustomMetricResponseBody) GoString() string {
	return s.String()
}

func (s *PutCustomMetricResponseBody) GetCode() *string {
	return s.Code
}

func (s *PutCustomMetricResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PutCustomMetricResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PutCustomMetricResponseBody) SetCode(v string) *PutCustomMetricResponseBody {
	s.Code = &v
	return s
}

func (s *PutCustomMetricResponseBody) SetMessage(v string) *PutCustomMetricResponseBody {
	s.Message = &v
	return s
}

func (s *PutCustomMetricResponseBody) SetRequestId(v string) *PutCustomMetricResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutCustomMetricResponseBody) Validate() error {
	return dara.Validate(s)
}
