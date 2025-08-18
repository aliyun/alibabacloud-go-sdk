// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSiteCustomLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateSiteCustomLogResponseBody
	GetRequestId() *string
}

type UpdateSiteCustomLogResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ET5BF670-09D5-4D0B-BEBY-D96A2A528000
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSiteCustomLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSiteCustomLogResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSiteCustomLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSiteCustomLogResponseBody) SetRequestId(v string) *UpdateSiteCustomLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSiteCustomLogResponseBody) Validate() error {
	return dara.Validate(s)
}
