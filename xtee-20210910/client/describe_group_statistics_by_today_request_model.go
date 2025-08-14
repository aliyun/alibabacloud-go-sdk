// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupStatisticsByTodayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeGroupStatisticsByTodayRequest
	GetLang() *string
	SetRegId(v string) *DescribeGroupStatisticsByTodayRequest
	GetRegId() *string
}

type DescribeGroupStatisticsByTodayRequest struct {
	// Sets the language type for requests and responses, with a default value of **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeGroupStatisticsByTodayRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupStatisticsByTodayRequest) GoString() string {
	return s.String()
}

func (s *DescribeGroupStatisticsByTodayRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeGroupStatisticsByTodayRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeGroupStatisticsByTodayRequest) SetLang(v string) *DescribeGroupStatisticsByTodayRequest {
	s.Lang = &v
	return s
}

func (s *DescribeGroupStatisticsByTodayRequest) SetRegId(v string) *DescribeGroupStatisticsByTodayRequest {
	s.RegId = &v
	return s
}

func (s *DescribeGroupStatisticsByTodayRequest) Validate() error {
	return dara.Validate(s)
}
