// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCrowdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCrowdResponseBody
	GetRequestId() *string
}

type UpdateCrowdResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 8C27790E-CCA5-56BB-BA17-646295DEC0A2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCrowdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCrowdResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCrowdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCrowdResponseBody) SetRequestId(v string) *UpdateCrowdResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCrowdResponseBody) Validate() error {
	return dara.Validate(s)
}
