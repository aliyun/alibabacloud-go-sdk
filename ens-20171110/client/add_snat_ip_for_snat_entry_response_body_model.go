// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSnatIpForSnatEntryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddSnatIpForSnatEntryResponseBody
	GetRequestId() *string
}

type AddSnatIpForSnatEntryResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CD1FFCC6-5E9E-5C31-A014-13D02737B0EA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddSnatIpForSnatEntryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddSnatIpForSnatEntryResponseBody) GoString() string {
	return s.String()
}

func (s *AddSnatIpForSnatEntryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddSnatIpForSnatEntryResponseBody) SetRequestId(v string) *AddSnatIpForSnatEntryResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddSnatIpForSnatEntryResponseBody) Validate() error {
	return dara.Validate(s)
}
