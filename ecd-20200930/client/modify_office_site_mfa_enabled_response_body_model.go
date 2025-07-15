// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyOfficeSiteMfaEnabledResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyOfficeSiteMfaEnabledResponseBody
	GetRequestId() *string
}

type ModifyOfficeSiteMfaEnabledResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyOfficeSiteMfaEnabledResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyOfficeSiteMfaEnabledResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyOfficeSiteMfaEnabledResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyOfficeSiteMfaEnabledResponseBody) SetRequestId(v string) *ModifyOfficeSiteMfaEnabledResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyOfficeSiteMfaEnabledResponseBody) Validate() error {
	return dara.Validate(s)
}
