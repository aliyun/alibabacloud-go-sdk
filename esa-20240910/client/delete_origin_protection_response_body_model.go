// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOriginProtectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteOriginProtectionResponseBody
	GetRequestId() *string
}

type DeleteOriginProtectionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CB1A380B-09F0-41BB-A198-72F8FD6DA2FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteOriginProtectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteOriginProtectionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteOriginProtectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteOriginProtectionResponseBody) SetRequestId(v string) *DeleteOriginProtectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteOriginProtectionResponseBody) Validate() error {
	return dara.Validate(s)
}
