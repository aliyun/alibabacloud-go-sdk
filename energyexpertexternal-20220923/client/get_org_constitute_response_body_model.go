// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOrgConstituteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *OrgEmission) *GetOrgConstituteResponseBody
	GetData() *OrgEmission
	SetRequestId(v string) *GetOrgConstituteResponseBody
	GetRequestId() *string
}

type GetOrgConstituteResponseBody struct {
	// The data returned.
	Data *OrgEmission `json:"data,omitempty" xml:"data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetOrgConstituteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOrgConstituteResponseBody) GoString() string {
	return s.String()
}

func (s *GetOrgConstituteResponseBody) GetData() *OrgEmission {
	return s.Data
}

func (s *GetOrgConstituteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOrgConstituteResponseBody) SetData(v *OrgEmission) *GetOrgConstituteResponseBody {
	s.Data = v
	return s
}

func (s *GetOrgConstituteResponseBody) SetRequestId(v string) *GetOrgConstituteResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOrgConstituteResponseBody) Validate() error {
	return dara.Validate(s)
}
