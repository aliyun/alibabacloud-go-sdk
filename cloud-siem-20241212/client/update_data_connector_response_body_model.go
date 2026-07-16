// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataConnectorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDataConnectorResponseBody
	GetRequestId() *string
}

type UpdateDataConnectorResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDataConnectorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataConnectorResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDataConnectorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDataConnectorResponseBody) SetRequestId(v string) *UpdateDataConnectorResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDataConnectorResponseBody) Validate() error {
	return dara.Validate(s)
}
