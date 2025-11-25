// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIpsPrivateAssocRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v string) *DescribeIpsPrivateAssocRequest
	GetCurrentPage() *string
	SetLang(v string) *DescribeIpsPrivateAssocRequest
	GetLang() *string
	SetMemberUid(v int64) *DescribeIpsPrivateAssocRequest
	GetMemberUid() *int64
	SetPageSize(v string) *DescribeIpsPrivateAssocRequest
	GetPageSize() *string
	SetPublicIp(v string) *DescribeIpsPrivateAssocRequest
	GetPublicIp() *string
	SetResourceId(v string) *DescribeIpsPrivateAssocRequest
	GetResourceId() *string
	SetStatus(v string) *DescribeIpsPrivateAssocRequest
	GetStatus() *string
}

type DescribeIpsPrivateAssocRequest struct {
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 258039427902****
	MemberUid *int64 `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 203.0.113.1
	PublicIp *string `json:"PublicIp,omitempty" xml:"PublicIp,omitempty"`
	// example:
	//
	// ngw-c5vhmjdfp5t****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// close
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeIpsPrivateAssocRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpsPrivateAssocRequest) GoString() string {
	return s.String()
}

func (s *DescribeIpsPrivateAssocRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeIpsPrivateAssocRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeIpsPrivateAssocRequest) GetMemberUid() *int64 {
	return s.MemberUid
}

func (s *DescribeIpsPrivateAssocRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeIpsPrivateAssocRequest) GetPublicIp() *string {
	return s.PublicIp
}

func (s *DescribeIpsPrivateAssocRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeIpsPrivateAssocRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeIpsPrivateAssocRequest) SetCurrentPage(v string) *DescribeIpsPrivateAssocRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeIpsPrivateAssocRequest) SetLang(v string) *DescribeIpsPrivateAssocRequest {
	s.Lang = &v
	return s
}

func (s *DescribeIpsPrivateAssocRequest) SetMemberUid(v int64) *DescribeIpsPrivateAssocRequest {
	s.MemberUid = &v
	return s
}

func (s *DescribeIpsPrivateAssocRequest) SetPageSize(v string) *DescribeIpsPrivateAssocRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeIpsPrivateAssocRequest) SetPublicIp(v string) *DescribeIpsPrivateAssocRequest {
	s.PublicIp = &v
	return s
}

func (s *DescribeIpsPrivateAssocRequest) SetResourceId(v string) *DescribeIpsPrivateAssocRequest {
	s.ResourceId = &v
	return s
}

func (s *DescribeIpsPrivateAssocRequest) SetStatus(v string) *DescribeIpsPrivateAssocRequest {
	s.Status = &v
	return s
}

func (s *DescribeIpsPrivateAssocRequest) Validate() error {
	return dara.Validate(s)
}
