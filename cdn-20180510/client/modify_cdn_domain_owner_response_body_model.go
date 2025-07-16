// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCdnDomainOwnerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v map[string]interface{}) *ModifyCdnDomainOwnerResponseBody
	GetContent() map[string]interface{}
	SetRequestId(v string) *ModifyCdnDomainOwnerResponseBody
	GetRequestId() *string
}

type ModifyCdnDomainOwnerResponseBody struct {
	// The description of the domain name transfer.
	//
	// example:
	//
	// The domain does not allow to transfer to a different account.
	Content map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C98E518B-024E-538E-8276-66310CB8667D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCdnDomainOwnerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCdnDomainOwnerResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCdnDomainOwnerResponseBody) GetContent() map[string]interface{} {
	return s.Content
}

func (s *ModifyCdnDomainOwnerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCdnDomainOwnerResponseBody) SetContent(v map[string]interface{}) *ModifyCdnDomainOwnerResponseBody {
	s.Content = v
	return s
}

func (s *ModifyCdnDomainOwnerResponseBody) SetRequestId(v string) *ModifyCdnDomainOwnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCdnDomainOwnerResponseBody) Validate() error {
	return dara.Validate(s)
}
