// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainCnameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCnameDatas(v *DescribeDcdnDomainCnameResponseBodyCnameDatas) *DescribeDcdnDomainCnameResponseBody
	GetCnameDatas() *DescribeDcdnDomainCnameResponseBodyCnameDatas
	SetRequestId(v string) *DescribeDcdnDomainCnameResponseBody
	GetRequestId() *string
}

type DescribeDcdnDomainCnameResponseBody struct {
	// The CNAME information.
	CnameDatas *DescribeDcdnDomainCnameResponseBodyCnameDatas `json:"CnameDatas,omitempty" xml:"CnameDatas,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnDomainCnameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainCnameResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainCnameResponseBody) GetCnameDatas() *DescribeDcdnDomainCnameResponseBodyCnameDatas {
	return s.CnameDatas
}

func (s *DescribeDcdnDomainCnameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnDomainCnameResponseBody) SetCnameDatas(v *DescribeDcdnDomainCnameResponseBodyCnameDatas) *DescribeDcdnDomainCnameResponseBody {
	s.CnameDatas = v
	return s
}

func (s *DescribeDcdnDomainCnameResponseBody) SetRequestId(v string) *DescribeDcdnDomainCnameResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainCnameResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainCnameResponseBodyCnameDatas struct {
	Data []*DescribeDcdnDomainCnameResponseBodyCnameDatasData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainCnameResponseBodyCnameDatas) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainCnameResponseBodyCnameDatas) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainCnameResponseBodyCnameDatas) GetData() []*DescribeDcdnDomainCnameResponseBodyCnameDatasData {
	return s.Data
}

func (s *DescribeDcdnDomainCnameResponseBodyCnameDatas) SetData(v []*DescribeDcdnDomainCnameResponseBodyCnameDatasData) *DescribeDcdnDomainCnameResponseBodyCnameDatas {
	s.Data = v
	return s
}

func (s *DescribeDcdnDomainCnameResponseBodyCnameDatas) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainCnameResponseBodyCnameDatasData struct {
	// The CNAME assigned to the domain name.
	//
	// example:
	//
	// *.com
	Cname *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	// The accelerated domain name.
	//
	// example:
	//
	// .example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	ErrMsg *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	Passed *string `json:"Passed,omitempty" xml:"Passed,omitempty"`
	// The configuration status of the CNAME record. If the operation returns 0 for the parameter, the configuration was successful. Otherwise, the configuration failed.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDcdnDomainCnameResponseBodyCnameDatasData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainCnameResponseBodyCnameDatasData) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainCnameResponseBodyCnameDatasData) GetCname() *string {
	return s.Cname
}

func (s *DescribeDcdnDomainCnameResponseBodyCnameDatasData) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDcdnDomainCnameResponseBodyCnameDatasData) GetErrMsg() *string {
	return s.ErrMsg
}

func (s *DescribeDcdnDomainCnameResponseBodyCnameDatasData) GetPassed() *string {
	return s.Passed
}

func (s *DescribeDcdnDomainCnameResponseBodyCnameDatasData) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeDcdnDomainCnameResponseBodyCnameDatasData) SetCname(v string) *DescribeDcdnDomainCnameResponseBodyCnameDatasData {
	s.Cname = &v
	return s
}

func (s *DescribeDcdnDomainCnameResponseBodyCnameDatasData) SetDomain(v string) *DescribeDcdnDomainCnameResponseBodyCnameDatasData {
	s.Domain = &v
	return s
}

func (s *DescribeDcdnDomainCnameResponseBodyCnameDatasData) SetErrMsg(v string) *DescribeDcdnDomainCnameResponseBodyCnameDatasData {
	s.ErrMsg = &v
	return s
}

func (s *DescribeDcdnDomainCnameResponseBodyCnameDatasData) SetPassed(v string) *DescribeDcdnDomainCnameResponseBodyCnameDatasData {
	s.Passed = &v
	return s
}

func (s *DescribeDcdnDomainCnameResponseBodyCnameDatasData) SetStatus(v int32) *DescribeDcdnDomainCnameResponseBodyCnameDatasData {
	s.Status = &v
	return s
}

func (s *DescribeDcdnDomainCnameResponseBodyCnameDatasData) Validate() error {
	return dara.Validate(s)
}
