// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnConditionIPBInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatas(v []*DescribeCdnConditionIPBInfoResponseBodyDatas) *DescribeCdnConditionIPBInfoResponseBody
	GetDatas() []*DescribeCdnConditionIPBInfoResponseBodyDatas
	SetRequestId(v string) *DescribeCdnConditionIPBInfoResponseBody
	GetRequestId() *string
}

type DescribeCdnConditionIPBInfoResponseBody struct {
	// The data that is returned.
	Datas []*DescribeCdnConditionIPBInfoResponseBodyDatas `json:"Datas,omitempty" xml:"Datas,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 2C4AA72D-8C00-1113-BD68-8BC4E3CF4FF6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCdnConditionIPBInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnConditionIPBInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnConditionIPBInfoResponseBody) GetDatas() []*DescribeCdnConditionIPBInfoResponseBodyDatas {
	return s.Datas
}

func (s *DescribeCdnConditionIPBInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCdnConditionIPBInfoResponseBody) SetDatas(v []*DescribeCdnConditionIPBInfoResponseBodyDatas) *DescribeCdnConditionIPBInfoResponseBody {
	s.Datas = v
	return s
}

func (s *DescribeCdnConditionIPBInfoResponseBody) SetRequestId(v string) *DescribeCdnConditionIPBInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnConditionIPBInfoResponseBody) Validate() error {
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

type DescribeCdnConditionIPBInfoResponseBodyDatas struct {
	// The configuration value.
	//
	// example:
	//
	// "[{\\"text\\":\\"阿鲁巴\\",\\"value\\":\\"AW\\"}]"
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeCdnConditionIPBInfoResponseBodyDatas) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnConditionIPBInfoResponseBodyDatas) GoString() string {
	return s.String()
}

func (s *DescribeCdnConditionIPBInfoResponseBodyDatas) GetValue() *string {
	return s.Value
}

func (s *DescribeCdnConditionIPBInfoResponseBodyDatas) SetValue(v string) *DescribeCdnConditionIPBInfoResponseBodyDatas {
	s.Value = &v
	return s
}

func (s *DescribeCdnConditionIPBInfoResponseBodyDatas) Validate() error {
	return dara.Validate(s)
}
