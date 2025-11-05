// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainCnameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCnameDatas(v *DescribeDomainCnameResponseBodyCnameDatas) *DescribeDomainCnameResponseBody
	GetCnameDatas() *DescribeDomainCnameResponseBodyCnameDatas
	SetRequestId(v string) *DescribeDomainCnameResponseBody
	GetRequestId() *string
}

type DescribeDomainCnameResponseBody struct {
	// Details about the CNAME detection results.
	CnameDatas *DescribeDomainCnameResponseBodyCnameDatas `json:"CnameDatas,omitempty" xml:"CnameDatas,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 06D29681-B7CD-4034-A8CC-28AFFA213539
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDomainCnameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainCnameResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainCnameResponseBody) GetCnameDatas() *DescribeDomainCnameResponseBodyCnameDatas {
	return s.CnameDatas
}

func (s *DescribeDomainCnameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainCnameResponseBody) SetCnameDatas(v *DescribeDomainCnameResponseBodyCnameDatas) *DescribeDomainCnameResponseBody {
	s.CnameDatas = v
	return s
}

func (s *DescribeDomainCnameResponseBody) SetRequestId(v string) *DescribeDomainCnameResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainCnameResponseBody) Validate() error {
	if s.CnameDatas != nil {
		if err := s.CnameDatas.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDomainCnameResponseBodyCnameDatas struct {
	Data []*DescribeDomainCnameResponseBodyCnameDatasData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s DescribeDomainCnameResponseBodyCnameDatas) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainCnameResponseBodyCnameDatas) GoString() string {
	return s.String()
}

func (s *DescribeDomainCnameResponseBodyCnameDatas) GetData() []*DescribeDomainCnameResponseBodyCnameDatasData {
	return s.Data
}

func (s *DescribeDomainCnameResponseBodyCnameDatas) SetData(v []*DescribeDomainCnameResponseBodyCnameDatasData) *DescribeDomainCnameResponseBodyCnameDatas {
	s.Data = v
	return s
}

func (s *DescribeDomainCnameResponseBodyCnameDatas) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDomainCnameResponseBodyCnameDatasData struct {
	// The CNAME assigned to the domain name by Alibaba Cloud CDN.
	//
	// example:
	//
	// a.com.w.alikunlun.net
	Cname *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	// The accelerated domain name.
	//
	// example:
	//
	// a.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	ErrMsg *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	Passed *string `json:"Passed,omitempty" xml:"Passed,omitempty"`
	// The CNAME detection result. Valid values:
	//
	// 	- 0: The DNS can detect the CNAME assigned to the domain name.
	//
	// 	- Value other than 0: The DNS cannot detect the CNAME assigned to the domain name.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDomainCnameResponseBodyCnameDatasData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainCnameResponseBodyCnameDatasData) GoString() string {
	return s.String()
}

func (s *DescribeDomainCnameResponseBodyCnameDatasData) GetCname() *string {
	return s.Cname
}

func (s *DescribeDomainCnameResponseBodyCnameDatasData) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDomainCnameResponseBodyCnameDatasData) GetErrMsg() *string {
	return s.ErrMsg
}

func (s *DescribeDomainCnameResponseBodyCnameDatasData) GetPassed() *string {
	return s.Passed
}

func (s *DescribeDomainCnameResponseBodyCnameDatasData) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeDomainCnameResponseBodyCnameDatasData) SetCname(v string) *DescribeDomainCnameResponseBodyCnameDatasData {
	s.Cname = &v
	return s
}

func (s *DescribeDomainCnameResponseBodyCnameDatasData) SetDomain(v string) *DescribeDomainCnameResponseBodyCnameDatasData {
	s.Domain = &v
	return s
}

func (s *DescribeDomainCnameResponseBodyCnameDatasData) SetErrMsg(v string) *DescribeDomainCnameResponseBodyCnameDatasData {
	s.ErrMsg = &v
	return s
}

func (s *DescribeDomainCnameResponseBodyCnameDatasData) SetPassed(v string) *DescribeDomainCnameResponseBodyCnameDatasData {
	s.Passed = &v
	return s
}

func (s *DescribeDomainCnameResponseBodyCnameDatasData) SetStatus(v int32) *DescribeDomainCnameResponseBodyCnameDatasData {
	s.Status = &v
	return s
}

func (s *DescribeDomainCnameResponseBodyCnameDatasData) Validate() error {
	return dara.Validate(s)
}
