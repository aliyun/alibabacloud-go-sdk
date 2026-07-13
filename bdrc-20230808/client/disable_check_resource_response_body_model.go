// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableCheckResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableCheckResourceResponseBody
	GetRequestId() *string
}

type DisableCheckResourceResponseBody struct {
	// Unique request identity
	//
	// example:
	//
	// 86DEBAC9-AB6A-59AB-9E5C-A540E579ECC9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableCheckResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableCheckResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DisableCheckResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableCheckResourceResponseBody) SetRequestId(v string) *DisableCheckResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableCheckResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
