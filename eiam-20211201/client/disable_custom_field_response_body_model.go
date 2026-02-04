// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableCustomFieldResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableCustomFieldResponseBody
	GetRequestId() *string
}

type DisableCustomFieldResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableCustomFieldResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableCustomFieldResponseBody) GoString() string {
	return s.String()
}

func (s *DisableCustomFieldResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableCustomFieldResponseBody) SetRequestId(v string) *DisableCustomFieldResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableCustomFieldResponseBody) Validate() error {
	return dara.Validate(s)
}
