// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetSupabaseProjectPasswordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ResetSupabaseProjectPasswordResponseBody
	GetRequestId() *string
}

type ResetSupabaseProjectPasswordResponseBody struct {
	// The unique ID of the request.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetSupabaseProjectPasswordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetSupabaseProjectPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ResetSupabaseProjectPasswordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetSupabaseProjectPasswordResponseBody) SetRequestId(v string) *ResetSupabaseProjectPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetSupabaseProjectPasswordResponseBody) Validate() error {
	return dara.Validate(s)
}
