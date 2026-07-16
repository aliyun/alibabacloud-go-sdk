// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConditionIPBInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatas(v []*DescribeConditionIPBInfoResponseBodyDatas) *DescribeConditionIPBInfoResponseBody
	GetDatas() []*DescribeConditionIPBInfoResponseBodyDatas
	SetRequestId(v string) *DescribeConditionIPBInfoResponseBody
	GetRequestId() *string
}

type DescribeConditionIPBInfoResponseBody struct {
	// The data details.
	Datas []*DescribeConditionIPBInfoResponseBodyDatas `json:"Datas,omitempty" xml:"Datas,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// EEEBE525-F576-1196-8DAF-2D70CA3F4D2F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeConditionIPBInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeConditionIPBInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeConditionIPBInfoResponseBody) GetDatas() []*DescribeConditionIPBInfoResponseBodyDatas {
	return s.Datas
}

func (s *DescribeConditionIPBInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeConditionIPBInfoResponseBody) SetDatas(v []*DescribeConditionIPBInfoResponseBodyDatas) *DescribeConditionIPBInfoResponseBody {
	s.Datas = v
	return s
}

func (s *DescribeConditionIPBInfoResponseBody) SetRequestId(v string) *DescribeConditionIPBInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeConditionIPBInfoResponseBody) Validate() error {
	if s.Datas != nil {
		for _, item := range s.Datas {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeConditionIPBInfoResponseBodyDatas struct {
	// The returned data details, including the name of the country, ISP, or region and the corresponding identifier code.
	//
	// example:
	//
	// [{\\"text\\":\\"安徽\\",\\"value\\":\\"340000\\"},{\\"text\\":\\"北京\\",\\"value\\":\\"110000\\"},{\\"text\\":\\"重庆\\",\\"value\\":\\"500000\\"}]
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeConditionIPBInfoResponseBodyDatas) String() string {
	return dara.Prettify(s)
}

func (s DescribeConditionIPBInfoResponseBodyDatas) GoString() string {
	return s.String()
}

func (s *DescribeConditionIPBInfoResponseBodyDatas) GetValue() *string {
	return s.Value
}

func (s *DescribeConditionIPBInfoResponseBodyDatas) SetValue(v string) *DescribeConditionIPBInfoResponseBodyDatas {
	s.Value = &v
	return s
}

func (s *DescribeConditionIPBInfoResponseBodyDatas) Validate() error {
	return dara.Validate(s)
}
