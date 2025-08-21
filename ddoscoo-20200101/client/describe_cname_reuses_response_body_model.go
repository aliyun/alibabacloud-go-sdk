// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCnameReusesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCnameReuses(v []*DescribeCnameReusesResponseBodyCnameReuses) *DescribeCnameReusesResponseBody
	GetCnameReuses() []*DescribeCnameReusesResponseBodyCnameReuses
	SetRequestId(v string) *DescribeCnameReusesResponseBody
	GetRequestId() *string
}

type DescribeCnameReusesResponseBody struct {
	CnameReuses []*DescribeCnameReusesResponseBodyCnameReuses `json:"CnameReuses,omitempty" xml:"CnameReuses,omitempty" type:"Repeated"`
	// example:
	//
	// 0bcf28g5-d57c-11e7-9bs0-d89d6717dxbc
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCnameReusesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCnameReusesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCnameReusesResponseBody) GetCnameReuses() []*DescribeCnameReusesResponseBodyCnameReuses {
	return s.CnameReuses
}

func (s *DescribeCnameReusesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCnameReusesResponseBody) SetCnameReuses(v []*DescribeCnameReusesResponseBodyCnameReuses) *DescribeCnameReusesResponseBody {
	s.CnameReuses = v
	return s
}

func (s *DescribeCnameReusesResponseBody) SetRequestId(v string) *DescribeCnameReusesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCnameReusesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCnameReusesResponseBodyCnameReuses struct {
	// example:
	//
	// 4o6ep6q217k9****.aliyunddos0004.com
	Cname *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// 1
	Enable *int32 `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s DescribeCnameReusesResponseBodyCnameReuses) String() string {
	return dara.Prettify(s)
}

func (s DescribeCnameReusesResponseBodyCnameReuses) GoString() string {
	return s.String()
}

func (s *DescribeCnameReusesResponseBodyCnameReuses) GetCname() *string {
	return s.Cname
}

func (s *DescribeCnameReusesResponseBodyCnameReuses) GetDomain() *string {
	return s.Domain
}

func (s *DescribeCnameReusesResponseBodyCnameReuses) GetEnable() *int32 {
	return s.Enable
}

func (s *DescribeCnameReusesResponseBodyCnameReuses) SetCname(v string) *DescribeCnameReusesResponseBodyCnameReuses {
	s.Cname = &v
	return s
}

func (s *DescribeCnameReusesResponseBodyCnameReuses) SetDomain(v string) *DescribeCnameReusesResponseBodyCnameReuses {
	s.Domain = &v
	return s
}

func (s *DescribeCnameReusesResponseBodyCnameReuses) SetEnable(v int32) *DescribeCnameReusesResponseBodyCnameReuses {
	s.Enable = &v
	return s
}

func (s *DescribeCnameReusesResponseBodyCnameReuses) Validate() error {
	return dara.Validate(s)
}
