// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteGatewayResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteGatewayResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteGatewayResponseBody
	GetRequestId() *string
}

type DeleteGatewayResponseBody struct {
	// Response status code.
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
	// DE97DFDB-7DF0-5AB9-941C-10D27D769E4B
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGatewayResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteGatewayResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteGatewayResponseBody) SetCode(v string) *DeleteGatewayResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteGatewayResponseBody) SetMessage(v string) *DeleteGatewayResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteGatewayResponseBody) SetRequestId(v string) *DeleteGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGatewayResponseBody) Validate() error {
	return dara.Validate(s)
}
