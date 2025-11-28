// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSupabaseProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSupabaseProjectResponseBody
	GetRequestId() *string
}

type DeleteSupabaseProjectResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSupabaseProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSupabaseProjectResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSupabaseProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSupabaseProjectResponseBody) SetRequestId(v string) *DeleteSupabaseProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSupabaseProjectResponseBody) Validate() error {
	return dara.Validate(s)
}
