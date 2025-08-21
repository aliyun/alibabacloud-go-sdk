// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartSnatIpForSnatEntryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartSnatIpForSnatEntryResponseBody
	GetRequestId() *string
}

type StartSnatIpForSnatEntryResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 85BBD491-CE05-5BDA-979E-843FE52B74CF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartSnatIpForSnatEntryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartSnatIpForSnatEntryResponseBody) GoString() string {
	return s.String()
}

func (s *StartSnatIpForSnatEntryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartSnatIpForSnatEntryResponseBody) SetRequestId(v string) *StartSnatIpForSnatEntryResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartSnatIpForSnatEntryResponseBody) Validate() error {
	return dara.Validate(s)
}
