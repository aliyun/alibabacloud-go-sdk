// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyDcdnDomainOwnerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *VerifyDcdnDomainOwnerResponseBody
	GetContent() *string
	SetRequestId(v string) *VerifyDcdnDomainOwnerResponseBody
	GetRequestId() *string
}

type VerifyDcdnDomainOwnerResponseBody struct {
	// The verification result.
	//
	// example:
	//
	// verify_dffeb6610035dcb77b413a59c32cd9**
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 34AB41F1-04A5-496F-8C8D-634BDBE6A9FB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VerifyDcdnDomainOwnerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VerifyDcdnDomainOwnerResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyDcdnDomainOwnerResponseBody) GetContent() *string {
	return s.Content
}

func (s *VerifyDcdnDomainOwnerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VerifyDcdnDomainOwnerResponseBody) SetContent(v string) *VerifyDcdnDomainOwnerResponseBody {
	s.Content = &v
	return s
}

func (s *VerifyDcdnDomainOwnerResponseBody) SetRequestId(v string) *VerifyDcdnDomainOwnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyDcdnDomainOwnerResponseBody) Validate() error {
	return dara.Validate(s)
}
