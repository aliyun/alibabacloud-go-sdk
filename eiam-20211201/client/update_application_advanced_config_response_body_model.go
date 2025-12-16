// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationAdvancedConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateApplicationAdvancedConfigResponseBody
	GetRequestId() *string
}

type UpdateApplicationAdvancedConfigResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateApplicationAdvancedConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationAdvancedConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApplicationAdvancedConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateApplicationAdvancedConfigResponseBody) SetRequestId(v string) *UpdateApplicationAdvancedConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateApplicationAdvancedConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
