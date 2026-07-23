// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSiteInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateSiteInstanceResponseBody
	GetRequestId() *string
}

type UpdateSiteInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0E50DBC5-0F50-583C-AC2D-87B2764776CF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSiteInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSiteInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSiteInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSiteInstanceResponseBody) SetRequestId(v string) *UpdateSiteInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSiteInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
