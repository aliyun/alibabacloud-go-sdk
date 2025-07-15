// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetIdpMetadataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIdpEntityId(v string) *SetIdpMetadataResponseBody
	GetIdpEntityId() *string
	SetRequestId(v string) *SetIdpMetadataResponseBody
	GetRequestId() *string
}

type SetIdpMetadataResponseBody struct {
	// The entity ID obtained after the IdP metadata file is parsed.
	//
	// example:
	//
	// http://test****.cn/adfs/services/trust
	IdpEntityId *string `json:"IdpEntityId,omitempty" xml:"IdpEntityId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetIdpMetadataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetIdpMetadataResponseBody) GoString() string {
	return s.String()
}

func (s *SetIdpMetadataResponseBody) GetIdpEntityId() *string {
	return s.IdpEntityId
}

func (s *SetIdpMetadataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetIdpMetadataResponseBody) SetIdpEntityId(v string) *SetIdpMetadataResponseBody {
	s.IdpEntityId = &v
	return s
}

func (s *SetIdpMetadataResponseBody) SetRequestId(v string) *SetIdpMetadataResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetIdpMetadataResponseBody) Validate() error {
	return dara.Validate(s)
}
