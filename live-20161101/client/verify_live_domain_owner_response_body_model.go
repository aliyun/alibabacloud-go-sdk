// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyLiveDomainOwnerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *VerifyLiveDomainOwnerResponseBody
	GetContent() *string
	SetRequestId(v string) *VerifyLiveDomainOwnerResponseBody
	GetRequestId() *string
}

type VerifyLiveDomainOwnerResponseBody struct {
	// The verification information.
	//
	// example:
	//
	// verify_dffeb6610035dcb77b413******
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1BA6D7CE-55F1-5926-8764-F8663473AD0D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VerifyLiveDomainOwnerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VerifyLiveDomainOwnerResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyLiveDomainOwnerResponseBody) GetContent() *string {
	return s.Content
}

func (s *VerifyLiveDomainOwnerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VerifyLiveDomainOwnerResponseBody) SetContent(v string) *VerifyLiveDomainOwnerResponseBody {
	s.Content = &v
	return s
}

func (s *VerifyLiveDomainOwnerResponseBody) SetRequestId(v string) *VerifyLiveDomainOwnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyLiveDomainOwnerResponseBody) Validate() error {
	return dara.Validate(s)
}
