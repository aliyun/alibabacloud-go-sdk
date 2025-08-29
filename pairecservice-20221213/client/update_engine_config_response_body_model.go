// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEngineConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateEngineConfigResponseBody
	GetRequestId() *string
}

type UpdateEngineConfigResponseBody struct {
	// example:
	//
	// F8F613A9-DF1C-551A-88E1-397A3981A785
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateEngineConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateEngineConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEngineConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateEngineConfigResponseBody) SetRequestId(v string) *UpdateEngineConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEngineConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
