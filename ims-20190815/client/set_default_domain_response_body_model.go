// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDefaultDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDefaultDomainName(v string) *SetDefaultDomainResponseBody
	GetDefaultDomainName() *string
	SetRequestId(v string) *SetDefaultDomainResponseBody
	GetRequestId() *string
}

type SetDefaultDomainResponseBody struct {
	// The default domain name.
	//
	// example:
	//
	// examplecompany.onaliyun.com
	DefaultDomainName *string `json:"DefaultDomainName,omitempty" xml:"DefaultDomainName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 66815255-7CCE-4759-AC37-9755794C3626
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDefaultDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetDefaultDomainResponseBody) GoString() string {
	return s.String()
}

func (s *SetDefaultDomainResponseBody) GetDefaultDomainName() *string {
	return s.DefaultDomainName
}

func (s *SetDefaultDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetDefaultDomainResponseBody) SetDefaultDomainName(v string) *SetDefaultDomainResponseBody {
	s.DefaultDomainName = &v
	return s
}

func (s *SetDefaultDomainResponseBody) SetRequestId(v string) *SetDefaultDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDefaultDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
