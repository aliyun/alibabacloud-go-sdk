// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVerifyPersonasSexStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeVerifyPersonasSexStatisticsResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeVerifyPersonasSexStatisticsResponseBodyResultObject) *DescribeVerifyPersonasSexStatisticsResponseBody
	GetResultObject() *DescribeVerifyPersonasSexStatisticsResponseBodyResultObject
}

type DescribeVerifyPersonasSexStatisticsResponseBody struct {
	// ID of this request.
	//
	// example:
	//
	// 013DA6E1-3F37-5579-B979-2F12B7E92450
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned data.
	ResultObject *DescribeVerifyPersonasSexStatisticsResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s DescribeVerifyPersonasSexStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifyPersonasSexStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVerifyPersonasSexStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVerifyPersonasSexStatisticsResponseBody) GetResultObject() *DescribeVerifyPersonasSexStatisticsResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeVerifyPersonasSexStatisticsResponseBody) SetRequestId(v string) *DescribeVerifyPersonasSexStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVerifyPersonasSexStatisticsResponseBody) SetResultObject(v *DescribeVerifyPersonasSexStatisticsResponseBodyResultObject) *DescribeVerifyPersonasSexStatisticsResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeVerifyPersonasSexStatisticsResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVerifyPersonasSexStatisticsResponseBodyResultObject struct {
	// Number of users under 14 years old.
	//
	// example:
	//
	// 0
	Age0To14Cnt *int64 `json:"Age0To14Cnt,omitempty" xml:"Age0To14Cnt,omitempty"`
	// Proportion of users under 14 years old.
	//
	// example:
	//
	// 0
	Age0To14Rate *string `json:"Age0To14Rate,omitempty" xml:"Age0To14Rate,omitempty"`
	// Number of users between 14 and 18 years old.
	//
	// example:
	//
	// 0
	Age14To18Cnt *int64 `json:"Age14To18Cnt,omitempty" xml:"Age14To18Cnt,omitempty"`
	// Proportion of users between 14 and 18 years old.
	//
	// example:
	//
	// 0
	Age14To18Rate *string `json:"Age14To18Rate,omitempty" xml:"Age14To18Rate,omitempty"`
	// Number of authenticated users between 18 and 35 years old.
	//
	// example:
	//
	// 9
	Age18To35Cnt *int64 `json:"Age18To35Cnt,omitempty" xml:"Age18To35Cnt,omitempty"`
	// Proportion of authenticated users between 18 and 35 years old.
	//
	// example:
	//
	// 64.29
	Age18To35Rate *string `json:"Age18To35Rate,omitempty" xml:"Age18To35Rate,omitempty"`
	// Number of authenticated users between 35 and 50 years old.
	//
	// example:
	//
	// 5
	Age35To50Cnt *int64 `json:"Age35To50Cnt,omitempty" xml:"Age35To50Cnt,omitempty"`
	// Proportion of users between 35 and 50 years old.
	//
	// example:
	//
	// 35.71
	Age35To50Rate *string `json:"Age35To50Rate,omitempty" xml:"Age35To50Rate,omitempty"`
	// Number of authenticated users over 50 years old.
	//
	// example:
	//
	// 0
	Age50To999Cnt *int64 `json:"Age50To999Cnt,omitempty" xml:"Age50To999Cnt,omitempty"`
	// Proportion of authenticated users over 50 years old.
	//
	// example:
	//
	// 0
	Age50To999Rate *string `json:"Age50To999Rate,omitempty" xml:"Age50To999Rate,omitempty"`
	// Total number of authenticated users.
	//
	// example:
	//
	// 14
	AllUserCnt *int64 `json:"AllUserCnt,omitempty" xml:"AllUserCnt,omitempty"`
	// Number of female users.
	//
	// example:
	//
	// 4
	FemaleCnt *int64 `json:"FemaleCnt,omitempty" xml:"FemaleCnt,omitempty"`
	// Proportion of female authenticated users.
	//
	// example:
	//
	// 28.57
	FemaleRate *string `json:"FemaleRate,omitempty" xml:"FemaleRate,omitempty"`
	// Number of male users.
	//
	// example:
	//
	// 10
	MaleCnt *int64 `json:"MaleCnt,omitempty" xml:"MaleCnt,omitempty"`
	// Proportion of male users.
	//
	// example:
	//
	// 71.43
	MaleRate *string `json:"MaleRate,omitempty" xml:"MaleRate,omitempty"`
}

func (s DescribeVerifyPersonasSexStatisticsResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifyPersonasSexStatisticsResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeVerifyPersonasSexStatisticsResponseBodyResultObject) GetAge0To14Cnt() *int64 {
	return s.Age0To14Cnt
}

func (s *DescribeVerifyPersonasSexStatisticsResponseBodyResultObject) GetAge0To14Rate() *string {
	return s.Age0To14Rate
}

func (s *DescribeVerifyPersonasSexStatisticsResponseBodyResultObject) GetAge14To18Cnt() *int64 {
	return s.Age14To18Cnt
}

func (s *DescribeVerifyPersonasSexStatisticsResponseBodyResultObject) GetAge14To18Rate() *string {
	return s.Age14To18Rate
}

func (s *DescribeVerifyPersonasSexStatisticsResponseBodyResultObject) GetAge18To35Cnt() *int64 {
	return s.Age18To35Cnt
}

func (s *DescribeVerifyPersonasSexStatisticsResponseBodyResultObject) GetAge18To35Rate() *string {
	return s.Age18To35Rate
}

func (s *DescribeVerifyPersonasSexStatisticsResponseBodyResultObject) GetAge35To50Cnt() *int64 {
	return s.Age35To50Cnt
}

func (s *DescribeVerifyPersonasSexStatisticsResponseBodyResultObject) GetAge35To50Rate() *string {
	return s.Age35To50Rate
}

func (s *DescribeVerifyPersonasSexStatisticsResponseBodyResultObject) GetAge50To999Cnt() *int64 {
	return s.Age50To999Cnt
}

func (s *DescribeVerifyPersonasSexStatisticsResponseBodyResultObject) GetAge50To999Rate() *string {
	return s.Age50To999Rate
}

func (s *DescribeVerifyPersonasSexStatisticsResponseBodyResultObject) GetAllUserCnt() *int64 {
	return s.AllUserCnt
}

func (s *DescribeVerifyPersonasSexStatisticsResponseBodyResultObject) GetFemaleCnt() *int64 {
	return s.FemaleCnt
}

func (s *DescribeVerifyPersonasSexStatisticsResponseBodyResultObject) GetFemaleRate() *string {
	return s.FemaleRate
}

func (s *DescribeVerifyPersonasSexStatisticsResponseBodyResultObject) GetMaleCnt() *int64 {
	return s.MaleCnt
}

func (s *DescribeVerifyPersonasSexStatisticsResponseBodyResultObject) GetMaleRate() *string {
	return s.MaleRate
}

func (s *DescribeVerifyPersonasSexStatisticsResponseBodyResultObject) SetAge0To14Cnt(v int64) *DescribeVerifyPersonasSexStatisticsResponseBodyResultObject {
	s.Age0To14Cnt = &v
	return s
}

func (s *DescribeVerifyPersonasSexStatisticsResponseBodyResultObject) SetAge0To14Rate(v string) *DescribeVerifyPersonasSexStatisticsResponseBodyResultObject {
	s.Age0To14Rate = &v
	return s
}

func (s *DescribeVerifyPersonasSexStatisticsResponseBodyResultObject) SetAge14To18Cnt(v int64) *DescribeVerifyPersonasSexStatisticsResponseBodyResultObject {
	s.Age14To18Cnt = &v
	return s
}

func (s *DescribeVerifyPersonasSexStatisticsResponseBodyResultObject) SetAge14To18Rate(v string) *DescribeVerifyPersonasSexStatisticsResponseBodyResultObject {
	s.Age14To18Rate = &v
	return s
}

func (s *DescribeVerifyPersonasSexStatisticsResponseBodyResultObject) SetAge18To35Cnt(v int64) *DescribeVerifyPersonasSexStatisticsResponseBodyResultObject {
	s.Age18To35Cnt = &v
	return s
}

func (s *DescribeVerifyPersonasSexStatisticsResponseBodyResultObject) SetAge18To35Rate(v string) *DescribeVerifyPersonasSexStatisticsResponseBodyResultObject {
	s.Age18To35Rate = &v
	return s
}

func (s *DescribeVerifyPersonasSexStatisticsResponseBodyResultObject) SetAge35To50Cnt(v int64) *DescribeVerifyPersonasSexStatisticsResponseBodyResultObject {
	s.Age35To50Cnt = &v
	return s
}

func (s *DescribeVerifyPersonasSexStatisticsResponseBodyResultObject) SetAge35To50Rate(v string) *DescribeVerifyPersonasSexStatisticsResponseBodyResultObject {
	s.Age35To50Rate = &v
	return s
}

func (s *DescribeVerifyPersonasSexStatisticsResponseBodyResultObject) SetAge50To999Cnt(v int64) *DescribeVerifyPersonasSexStatisticsResponseBodyResultObject {
	s.Age50To999Cnt = &v
	return s
}

func (s *DescribeVerifyPersonasSexStatisticsResponseBodyResultObject) SetAge50To999Rate(v string) *DescribeVerifyPersonasSexStatisticsResponseBodyResultObject {
	s.Age50To999Rate = &v
	return s
}

func (s *DescribeVerifyPersonasSexStatisticsResponseBodyResultObject) SetAllUserCnt(v int64) *DescribeVerifyPersonasSexStatisticsResponseBodyResultObject {
	s.AllUserCnt = &v
	return s
}

func (s *DescribeVerifyPersonasSexStatisticsResponseBodyResultObject) SetFemaleCnt(v int64) *DescribeVerifyPersonasSexStatisticsResponseBodyResultObject {
	s.FemaleCnt = &v
	return s
}

func (s *DescribeVerifyPersonasSexStatisticsResponseBodyResultObject) SetFemaleRate(v string) *DescribeVerifyPersonasSexStatisticsResponseBodyResultObject {
	s.FemaleRate = &v
	return s
}

func (s *DescribeVerifyPersonasSexStatisticsResponseBodyResultObject) SetMaleCnt(v int64) *DescribeVerifyPersonasSexStatisticsResponseBodyResultObject {
	s.MaleCnt = &v
	return s
}

func (s *DescribeVerifyPersonasSexStatisticsResponseBodyResultObject) SetMaleRate(v string) *DescribeVerifyPersonasSexStatisticsResponseBodyResultObject {
	s.MaleRate = &v
	return s
}

func (s *DescribeVerifyPersonasSexStatisticsResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
