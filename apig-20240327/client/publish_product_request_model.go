// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishProductRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPortalId(v string) *PublishProductRequest
	GetPortalId() *string
}

type PublishProductRequest struct {
	// The ID of the HiMarket portal to which the API product is published. This parameter specifies the target portal for the publication. You can query the publication record by portal dimension by calling ListPortalPublications. This ID also appears in the publication record details and is paired with portalName.
	//
	// example:
	//
	// portal-xxxxxxxx
	PortalId *string `json:"portalId,omitempty" xml:"portalId,omitempty"`
}

func (s PublishProductRequest) String() string {
	return dara.Prettify(s)
}

func (s PublishProductRequest) GoString() string {
	return s.String()
}

func (s *PublishProductRequest) GetPortalId() *string {
	return s.PortalId
}

func (s *PublishProductRequest) SetPortalId(v string) *PublishProductRequest {
	s.PortalId = &v
	return s
}

func (s *PublishProductRequest) Validate() error {
	return dara.Validate(s)
}
