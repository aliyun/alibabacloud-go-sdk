// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachApplication2ConnectorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DetachApplication2ConnectorResponseBody
	GetRequestId() *string
}

type DetachApplication2ConnectorResponseBody struct {
	// example:
	//
	// 4D169859-A4F2-5EC8-853B-8447787C0D8A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachApplication2ConnectorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachApplication2ConnectorResponseBody) GoString() string {
	return s.String()
}

func (s *DetachApplication2ConnectorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachApplication2ConnectorResponseBody) SetRequestId(v string) *DetachApplication2ConnectorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachApplication2ConnectorResponseBody) Validate() error {
	return dara.Validate(s)
}
