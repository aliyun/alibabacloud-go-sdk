// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExpressConnectTrafficQosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetQosId(v string) *CreateExpressConnectTrafficQosResponseBody
	GetQosId() *string
	SetRequestId(v string) *CreateExpressConnectTrafficQosResponseBody
	GetRequestId() *string
}

type CreateExpressConnectTrafficQosResponseBody struct {
	// The ID of the QoS policy.
	//
	// example:
	//
	// qos-2giu0a6vd5x0mv4700
	QosId *string `json:"QosId,omitempty" xml:"QosId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// DC668356-BCB4-42FD-9BC3-FA2B2E04B634
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateExpressConnectTrafficQosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateExpressConnectTrafficQosResponseBody) GoString() string {
	return s.String()
}

func (s *CreateExpressConnectTrafficQosResponseBody) GetQosId() *string {
	return s.QosId
}

func (s *CreateExpressConnectTrafficQosResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateExpressConnectTrafficQosResponseBody) SetQosId(v string) *CreateExpressConnectTrafficQosResponseBody {
	s.QosId = &v
	return s
}

func (s *CreateExpressConnectTrafficQosResponseBody) SetRequestId(v string) *CreateExpressConnectTrafficQosResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateExpressConnectTrafficQosResponseBody) Validate() error {
	return dara.Validate(s)
}
