// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayLabelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *DeleteGatewayLabelResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteGatewayLabelResponseBody
	GetRequestId() *string
}

type DeleteGatewayLabelResponseBody struct {
	// example:
	//
	// Succeed to delete gateway labels.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteGatewayLabelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayLabelResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGatewayLabelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteGatewayLabelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteGatewayLabelResponseBody) SetMessage(v string) *DeleteGatewayLabelResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteGatewayLabelResponseBody) SetRequestId(v string) *DeleteGatewayLabelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGatewayLabelResponseBody) Validate() error {
	return dara.Validate(s)
}
