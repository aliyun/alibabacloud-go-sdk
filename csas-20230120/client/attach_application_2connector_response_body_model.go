// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachApplication2ConnectorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AttachApplication2ConnectorResponseBody
	GetRequestId() *string
}

type AttachApplication2ConnectorResponseBody struct {
	// example:
	//
	// 7E9D7ACD-53D5-56EF-A913-79D148D06299
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachApplication2ConnectorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachApplication2ConnectorResponseBody) GoString() string {
	return s.String()
}

func (s *AttachApplication2ConnectorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachApplication2ConnectorResponseBody) SetRequestId(v string) *AttachApplication2ConnectorResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachApplication2ConnectorResponseBody) Validate() error {
	return dara.Validate(s)
}
