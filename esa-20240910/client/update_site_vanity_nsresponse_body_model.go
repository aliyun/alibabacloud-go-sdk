// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSiteVanityNSResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateSiteVanityNSResponseBody
	GetRequestId() *string
}

type UpdateSiteVanityNSResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSiteVanityNSResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSiteVanityNSResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSiteVanityNSResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSiteVanityNSResponseBody) SetRequestId(v string) *UpdateSiteVanityNSResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSiteVanityNSResponseBody) Validate() error {
	return dara.Validate(s)
}
