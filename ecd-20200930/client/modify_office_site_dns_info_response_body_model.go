// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyOfficeSiteDnsInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyOfficeSiteDnsInfoResponseBody
	GetRequestId() *string
}

type ModifyOfficeSiteDnsInfoResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F7E4322D-D679-5ACB-A909-490D2F0E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyOfficeSiteDnsInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyOfficeSiteDnsInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyOfficeSiteDnsInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyOfficeSiteDnsInfoResponseBody) SetRequestId(v string) *ModifyOfficeSiteDnsInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyOfficeSiteDnsInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
