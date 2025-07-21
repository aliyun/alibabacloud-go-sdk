// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIpProtectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIpProtection(v string) *GetIpProtectionResponseBody
	GetIpProtection() *string
	SetRequestId(v string) *GetIpProtectionResponseBody
	GetRequestId() *string
}

type GetIpProtectionResponseBody struct {
	// IP protection switch, On: 1 Off: 0
	//
	// example:
	//
	// 0
	IpProtection *string `json:"IpProtection,omitempty" xml:"IpProtection,omitempty"`
	// Request ID
	//
	// example:
	//
	// B30E5A62-2E64-577D-A70E-8C6781D6C975
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetIpProtectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetIpProtectionResponseBody) GoString() string {
	return s.String()
}

func (s *GetIpProtectionResponseBody) GetIpProtection() *string {
	return s.IpProtection
}

func (s *GetIpProtectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetIpProtectionResponseBody) SetIpProtection(v string) *GetIpProtectionResponseBody {
	s.IpProtection = &v
	return s
}

func (s *GetIpProtectionResponseBody) SetRequestId(v string) *GetIpProtectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIpProtectionResponseBody) Validate() error {
	return dara.Validate(s)
}
