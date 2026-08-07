// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableCheckProductResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableCheckProductResponseBody
	GetRequestId() *string
}

type DisableCheckProductResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 89E3CBB7-16F3-52AE-BD32-31A43A2A807F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableCheckProductResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableCheckProductResponseBody) GoString() string {
	return s.String()
}

func (s *DisableCheckProductResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableCheckProductResponseBody) SetRequestId(v string) *DisableCheckProductResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableCheckProductResponseBody) Validate() error {
	return dara.Validate(s)
}
