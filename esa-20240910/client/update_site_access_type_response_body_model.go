// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSiteAccessTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateSiteAccessTypeResponseBody
	GetRequestId() *string
}

type UpdateSiteAccessTypeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSiteAccessTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSiteAccessTypeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSiteAccessTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSiteAccessTypeResponseBody) SetRequestId(v string) *UpdateSiteAccessTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSiteAccessTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
