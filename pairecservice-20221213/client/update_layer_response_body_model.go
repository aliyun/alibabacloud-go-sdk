// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLayerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateLayerResponseBody
	GetRequestId() *string
}

type UpdateLayerResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 0EA9215E-EC21-53AB-B8D9-D3DEA90D040A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLayerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLayerResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLayerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLayerResponseBody) SetRequestId(v string) *UpdateLayerResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLayerResponseBody) Validate() error {
	return dara.Validate(s)
}
