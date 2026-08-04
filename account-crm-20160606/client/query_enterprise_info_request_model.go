// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEnterpriseInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnterpriseVersion(v string) *QueryEnterpriseInfoRequest
	GetEnterpriseVersion() *string
	SetHavanaId(v string) *QueryEnterpriseInfoRequest
	GetHavanaId() *string
	SetPK(v string) *QueryEnterpriseInfoRequest
	GetPK() *string
}

type QueryEnterpriseInfoRequest struct {
	EnterpriseVersion *string `json:"EnterpriseVersion,omitempty" xml:"EnterpriseVersion,omitempty"`
	HavanaId          *string `json:"HavanaId,omitempty" xml:"HavanaId,omitempty"`
	PK                *string `json:"PK,omitempty" xml:"PK,omitempty"`
}

func (s QueryEnterpriseInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryEnterpriseInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryEnterpriseInfoRequest) GetEnterpriseVersion() *string {
	return s.EnterpriseVersion
}

func (s *QueryEnterpriseInfoRequest) GetHavanaId() *string {
	return s.HavanaId
}

func (s *QueryEnterpriseInfoRequest) GetPK() *string {
	return s.PK
}

func (s *QueryEnterpriseInfoRequest) SetEnterpriseVersion(v string) *QueryEnterpriseInfoRequest {
	s.EnterpriseVersion = &v
	return s
}

func (s *QueryEnterpriseInfoRequest) SetHavanaId(v string) *QueryEnterpriseInfoRequest {
	s.HavanaId = &v
	return s
}

func (s *QueryEnterpriseInfoRequest) SetPK(v string) *QueryEnterpriseInfoRequest {
	s.PK = &v
	return s
}

func (s *QueryEnterpriseInfoRequest) Validate() error {
	return dara.Validate(s)
}
