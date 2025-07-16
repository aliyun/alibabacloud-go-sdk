// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyDomainOwnerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *VerifyDomainOwnerResponseBody
	GetContent() *string
	SetRequestId(v string) *VerifyDomainOwnerResponseBody
	GetRequestId() *string
}

type VerifyDomainOwnerResponseBody struct {
	// The verification result.
	//
	// > This parameter is returned only if the operation fails.
	//
	// example:
	//
	// verify_dffeb6610035dcb77b413a59c32c****
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 34AB41F1-04A5-496F-8C8D-634BDBE6A9FB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VerifyDomainOwnerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VerifyDomainOwnerResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyDomainOwnerResponseBody) GetContent() *string {
	return s.Content
}

func (s *VerifyDomainOwnerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VerifyDomainOwnerResponseBody) SetContent(v string) *VerifyDomainOwnerResponseBody {
	s.Content = &v
	return s
}

func (s *VerifyDomainOwnerResponseBody) SetRequestId(v string) *VerifyDomainOwnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyDomainOwnerResponseBody) Validate() error {
	return dara.Validate(s)
}
