// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrgResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrgId(v string) *CreateOrgResponseBody
	GetOrgId() *string
	SetRequestId(v string) *CreateOrgResponseBody
	GetRequestId() *string
}

type CreateOrgResponseBody struct {
	// example:
	//
	// org-evk12ozjvmlxl****
	OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	// example:
	//
	// 7A2C3803-C975-5871-A232-80A91009****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateOrgResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOrgResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrgResponseBody) GetOrgId() *string {
	return s.OrgId
}

func (s *CreateOrgResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOrgResponseBody) SetOrgId(v string) *CreateOrgResponseBody {
	s.OrgId = &v
	return s
}

func (s *CreateOrgResponseBody) SetRequestId(v string) *CreateOrgResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOrgResponseBody) Validate() error {
	return dara.Validate(s)
}
