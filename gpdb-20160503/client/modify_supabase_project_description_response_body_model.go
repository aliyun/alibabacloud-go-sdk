// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySupabaseProjectDescriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifySupabaseProjectDescriptionResponseBody
	GetRequestId() *string
}

type ModifySupabaseProjectDescriptionResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySupabaseProjectDescriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySupabaseProjectDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySupabaseProjectDescriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySupabaseProjectDescriptionResponseBody) SetRequestId(v string) *ModifySupabaseProjectDescriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySupabaseProjectDescriptionResponseBody) Validate() error {
	return dara.Validate(s)
}
