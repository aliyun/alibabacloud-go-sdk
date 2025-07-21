// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableKeyResponseBody
	GetRequestId() *string
}

type DisableKeyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 2fe70ce2-3303-4fd6-b3ac-472fb2705c62
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DisableKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableKeyResponseBody) SetRequestId(v string) *DisableKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
