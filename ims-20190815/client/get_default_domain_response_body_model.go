// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDefaultDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDefaultDomainName(v string) *GetDefaultDomainResponseBody
	GetDefaultDomainName() *string
	SetRequestId(v string) *GetDefaultDomainResponseBody
	GetRequestId() *string
}

type GetDefaultDomainResponseBody struct {
	DefaultDomainName *string `json:"DefaultDomainName,omitempty" xml:"DefaultDomainName,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDefaultDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDefaultDomainResponseBody) GoString() string {
	return s.String()
}

func (s *GetDefaultDomainResponseBody) GetDefaultDomainName() *string {
	return s.DefaultDomainName
}

func (s *GetDefaultDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDefaultDomainResponseBody) SetDefaultDomainName(v string) *GetDefaultDomainResponseBody {
	s.DefaultDomainName = &v
	return s
}

func (s *GetDefaultDomainResponseBody) SetRequestId(v string) *GetDefaultDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDefaultDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
