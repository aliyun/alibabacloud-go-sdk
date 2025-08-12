// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomMetricResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteCustomMetricResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteCustomMetricResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteCustomMetricResponseBody
	GetRequestId() *string
}

type DeleteCustomMetricResponseBody struct {
	// The HTTP status code.
	//
	// >  The status code 200 indicates that the call is successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 05B36C2C-5F6E-48D5-8B41-CE36DD7EE8E0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCustomMetricResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomMetricResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomMetricResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteCustomMetricResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteCustomMetricResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCustomMetricResponseBody) SetCode(v string) *DeleteCustomMetricResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteCustomMetricResponseBody) SetMessage(v string) *DeleteCustomMetricResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteCustomMetricResponseBody) SetRequestId(v string) *DeleteCustomMetricResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCustomMetricResponseBody) Validate() error {
	return dara.Validate(s)
}
