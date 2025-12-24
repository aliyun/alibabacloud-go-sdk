// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayLabelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *UpdateGatewayLabelResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateGatewayLabelResponseBody
	GetRequestId() *string
}

type UpdateGatewayLabelResponseBody struct {
	// example:
	//
	// Succeed to update gateway gw-1uhcqmsc7x22*****	- labels.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateGatewayLabelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayLabelResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGatewayLabelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateGatewayLabelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGatewayLabelResponseBody) SetMessage(v string) *UpdateGatewayLabelResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateGatewayLabelResponseBody) SetRequestId(v string) *UpdateGatewayLabelResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGatewayLabelResponseBody) Validate() error {
	return dara.Validate(s)
}
