// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyVodDomainOwnerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *VerifyVodDomainOwnerResponseBody
	GetContent() *string
	SetRequestId(v string) *VerifyVodDomainOwnerResponseBody
	GetRequestId() *string
}

type VerifyVodDomainOwnerResponseBody struct {
	// The verification content.
	//
	// example:
	//
	// verify_dffeb661*********a59c32cd91f
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// E4EBD2BF-5EB0-4476-8829-9D94E1B1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VerifyVodDomainOwnerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VerifyVodDomainOwnerResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyVodDomainOwnerResponseBody) GetContent() *string {
	return s.Content
}

func (s *VerifyVodDomainOwnerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VerifyVodDomainOwnerResponseBody) SetContent(v string) *VerifyVodDomainOwnerResponseBody {
	s.Content = &v
	return s
}

func (s *VerifyVodDomainOwnerResponseBody) SetRequestId(v string) *VerifyVodDomainOwnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyVodDomainOwnerResponseBody) Validate() error {
	return dara.Validate(s)
}
