// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDomainCertResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDomainCertResponseBody
	GetRequestId() *string
}

type ModifyDomainCertResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 03E3B4DD-2BE7-5D9D-80E8-***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDomainCertResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDomainCertResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDomainCertResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDomainCertResponseBody) SetRequestId(v string) *ModifyDomainCertResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDomainCertResponseBody) Validate() error {
	return dara.Validate(s)
}
