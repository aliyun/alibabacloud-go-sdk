// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainCnameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCnameDatas(v *DescribeVodDomainCnameResponseBodyCnameDatas) *DescribeVodDomainCnameResponseBody
	GetCnameDatas() *DescribeVodDomainCnameResponseBodyCnameDatas
	SetRequestId(v string) *DescribeVodDomainCnameResponseBody
	GetRequestId() *string
}

type DescribeVodDomainCnameResponseBody struct {
	CnameDatas *DescribeVodDomainCnameResponseBodyCnameDatas `json:"CnameDatas,omitempty" xml:"CnameDatas,omitempty" type:"Struct"`
	RequestId  *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVodDomainCnameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainCnameResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainCnameResponseBody) GetCnameDatas() *DescribeVodDomainCnameResponseBodyCnameDatas {
	return s.CnameDatas
}

func (s *DescribeVodDomainCnameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodDomainCnameResponseBody) SetCnameDatas(v *DescribeVodDomainCnameResponseBodyCnameDatas) *DescribeVodDomainCnameResponseBody {
	s.CnameDatas = v
	return s
}

func (s *DescribeVodDomainCnameResponseBody) SetRequestId(v string) *DescribeVodDomainCnameResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodDomainCnameResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainCnameResponseBodyCnameDatas struct {
	Data []*DescribeVodDomainCnameResponseBodyCnameDatasData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s DescribeVodDomainCnameResponseBodyCnameDatas) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainCnameResponseBodyCnameDatas) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainCnameResponseBodyCnameDatas) GetData() []*DescribeVodDomainCnameResponseBodyCnameDatasData {
	return s.Data
}

func (s *DescribeVodDomainCnameResponseBodyCnameDatas) SetData(v []*DescribeVodDomainCnameResponseBodyCnameDatasData) *DescribeVodDomainCnameResponseBodyCnameDatas {
	s.Data = v
	return s
}

func (s *DescribeVodDomainCnameResponseBodyCnameDatas) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainCnameResponseBodyCnameDatasData struct {
	Cname  *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Status *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeVodDomainCnameResponseBodyCnameDatasData) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainCnameResponseBodyCnameDatasData) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainCnameResponseBodyCnameDatasData) GetCname() *string {
	return s.Cname
}

func (s *DescribeVodDomainCnameResponseBodyCnameDatasData) GetDomain() *string {
	return s.Domain
}

func (s *DescribeVodDomainCnameResponseBodyCnameDatasData) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeVodDomainCnameResponseBodyCnameDatasData) SetCname(v string) *DescribeVodDomainCnameResponseBodyCnameDatasData {
	s.Cname = &v
	return s
}

func (s *DescribeVodDomainCnameResponseBodyCnameDatasData) SetDomain(v string) *DescribeVodDomainCnameResponseBodyCnameDatasData {
	s.Domain = &v
	return s
}

func (s *DescribeVodDomainCnameResponseBodyCnameDatasData) SetStatus(v int32) *DescribeVodDomainCnameResponseBodyCnameDatasData {
	s.Status = &v
	return s
}

func (s *DescribeVodDomainCnameResponseBodyCnameDatasData) Validate() error {
	return dara.Validate(s)
}
