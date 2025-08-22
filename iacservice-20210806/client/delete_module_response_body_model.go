// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteModuleResponseBody
	GetRequestId() *string
}

type DeleteModuleResponseBody struct {
	// example:
	//
	// 49DA6457-E545-5095-A930-EB8F0BCD4DAA
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteModuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteModuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteModuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteModuleResponseBody) SetRequestId(v string) *DeleteModuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteModuleResponseBody) Validate() error {
	return dara.Validate(s)
}
