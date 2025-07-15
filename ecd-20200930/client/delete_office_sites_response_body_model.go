// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOfficeSitesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteOfficeSitesResponseBody
	GetRequestId() *string
}

type DeleteOfficeSitesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteOfficeSitesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteOfficeSitesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteOfficeSitesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteOfficeSitesResponseBody) SetRequestId(v string) *DeleteOfficeSitesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteOfficeSitesResponseBody) Validate() error {
	return dara.Validate(s)
}
