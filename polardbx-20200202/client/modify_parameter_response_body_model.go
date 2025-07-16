// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyParameterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyParameterResponseBody
	GetRequestId() *string
}

type ModifyParameterResponseBody struct {
	// example:
	//
	// FEA5DC20-6D8A-5979-97AA-FC57546ADC20
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyParameterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyParameterResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyParameterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyParameterResponseBody) SetRequestId(v string) *ModifyParameterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyParameterResponseBody) Validate() error {
	return dara.Validate(s)
}
