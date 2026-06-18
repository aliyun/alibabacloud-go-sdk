// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSubPartnerListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNo(v int32) *GetSubPartnerListRequest
	GetPageNo() *int32
	SetPageSize(v int32) *GetSubPartnerListRequest
	GetPageSize() *int32
	SetSubPartnerCompanyName(v string) *GetSubPartnerListRequest
	GetSubPartnerCompanyName() *string
	SetSubPartnerPid(v string) *GetSubPartnerListRequest
	GetSubPartnerPid() *string
}

type GetSubPartnerListRequest struct {
	// The page number, starting from 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Maximum value: 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The company name of the secondary partner.
	//
	// example:
	//
	// xxx有限公司
	SubPartnerCompanyName *string `json:"SubPartnerCompanyName,omitempty" xml:"SubPartnerCompanyName,omitempty"`
	// The PID of the secondary partner.
	//
	// example:
	//
	// 2323431211
	SubPartnerPid *string `json:"SubPartnerPid,omitempty" xml:"SubPartnerPid,omitempty"`
}

func (s GetSubPartnerListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSubPartnerListRequest) GoString() string {
	return s.String()
}

func (s *GetSubPartnerListRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *GetSubPartnerListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetSubPartnerListRequest) GetSubPartnerCompanyName() *string {
	return s.SubPartnerCompanyName
}

func (s *GetSubPartnerListRequest) GetSubPartnerPid() *string {
	return s.SubPartnerPid
}

func (s *GetSubPartnerListRequest) SetPageNo(v int32) *GetSubPartnerListRequest {
	s.PageNo = &v
	return s
}

func (s *GetSubPartnerListRequest) SetPageSize(v int32) *GetSubPartnerListRequest {
	s.PageSize = &v
	return s
}

func (s *GetSubPartnerListRequest) SetSubPartnerCompanyName(v string) *GetSubPartnerListRequest {
	s.SubPartnerCompanyName = &v
	return s
}

func (s *GetSubPartnerListRequest) SetSubPartnerPid(v string) *GetSubPartnerListRequest {
	s.SubPartnerPid = &v
	return s
}

func (s *GetSubPartnerListRequest) Validate() error {
	return dara.Validate(s)
}
