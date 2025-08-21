// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyVsDomainOwnerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *VerifyVsDomainOwnerResponseBody
	GetContent() *string
	SetRequestId(v string) *VerifyVsDomainOwnerResponseBody
	GetRequestId() *string
}

type VerifyVsDomainOwnerResponseBody struct {
	// example:
	//
	// verify_dffeb6610035dcb77b413a59c32c****
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VerifyVsDomainOwnerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VerifyVsDomainOwnerResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyVsDomainOwnerResponseBody) GetContent() *string {
	return s.Content
}

func (s *VerifyVsDomainOwnerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VerifyVsDomainOwnerResponseBody) SetContent(v string) *VerifyVsDomainOwnerResponseBody {
	s.Content = &v
	return s
}

func (s *VerifyVsDomainOwnerResponseBody) SetRequestId(v string) *VerifyVsDomainOwnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyVsDomainOwnerResponseBody) Validate() error {
	return dara.Validate(s)
}
