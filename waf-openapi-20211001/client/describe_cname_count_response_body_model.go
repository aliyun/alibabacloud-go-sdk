// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCnameCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCnameCount(v *DescribeCnameCountResponseBodyCnameCount) *DescribeCnameCountResponseBody
	GetCnameCount() *DescribeCnameCountResponseBodyCnameCount
	SetRequestId(v string) *DescribeCnameCountResponseBody
	GetRequestId() *string
}

type DescribeCnameCountResponseBody struct {
	// The information about the number of domain names that are added to WAF in CNAME record mode and hybrid cloud reverse proxy mode.
	CnameCount *DescribeCnameCountResponseBodyCnameCount `json:"CnameCount,omitempty" xml:"CnameCount,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// F35F45B0-5D6B-4238-BE02-A62D****E840
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCnameCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCnameCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCnameCountResponseBody) GetCnameCount() *DescribeCnameCountResponseBodyCnameCount {
	return s.CnameCount
}

func (s *DescribeCnameCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCnameCountResponseBody) SetCnameCount(v *DescribeCnameCountResponseBodyCnameCount) *DescribeCnameCountResponseBody {
	s.CnameCount = v
	return s
}

func (s *DescribeCnameCountResponseBody) SetRequestId(v string) *DescribeCnameCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCnameCountResponseBody) Validate() error {
	if s.CnameCount != nil {
		if err := s.CnameCount.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCnameCountResponseBodyCnameCount struct {
	// The number of domain names that are added to WAF in CNAME record mode.
	//
	// example:
	//
	// 1
	Cname *int64 `json:"Cname,omitempty" xml:"Cname,omitempty"`
	// The number of domain names that are added to WAF in hybrid cloud reverse proxy mode.
	//
	// example:
	//
	// 1
	HybridCloudCname *int64 `json:"HybridCloudCname,omitempty" xml:"HybridCloudCname,omitempty"`
	// The total number of domain names that are added to WAF in CNAME record mode and hybrid cloud reverse proxy mode.
	//
	// example:
	//
	// 1
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeCnameCountResponseBodyCnameCount) String() string {
	return dara.Prettify(s)
}

func (s DescribeCnameCountResponseBodyCnameCount) GoString() string {
	return s.String()
}

func (s *DescribeCnameCountResponseBodyCnameCount) GetCname() *int64 {
	return s.Cname
}

func (s *DescribeCnameCountResponseBodyCnameCount) GetHybridCloudCname() *int64 {
	return s.HybridCloudCname
}

func (s *DescribeCnameCountResponseBodyCnameCount) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeCnameCountResponseBodyCnameCount) SetCname(v int64) *DescribeCnameCountResponseBodyCnameCount {
	s.Cname = &v
	return s
}

func (s *DescribeCnameCountResponseBodyCnameCount) SetHybridCloudCname(v int64) *DescribeCnameCountResponseBodyCnameCount {
	s.HybridCloudCname = &v
	return s
}

func (s *DescribeCnameCountResponseBodyCnameCount) SetTotal(v int64) *DescribeCnameCountResponseBodyCnameCount {
	s.Total = &v
	return s
}

func (s *DescribeCnameCountResponseBodyCnameCount) Validate() error {
	return dara.Validate(s)
}
