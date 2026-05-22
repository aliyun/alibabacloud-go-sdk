// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOriginProtectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateOriginProtectionResponseBody
	GetRequestId() *string
}

type UpdateOriginProtectionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4C6B5E5A-42FC-5DF2-986C-4DAAE3C55086
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateOriginProtectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateOriginProtectionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateOriginProtectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateOriginProtectionResponseBody) SetRequestId(v string) *UpdateOriginProtectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateOriginProtectionResponseBody) Validate() error {
	return dara.Validate(s)
}
