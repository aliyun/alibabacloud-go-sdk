// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAmountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunUidList(v []*string) *QueryAmountRequest
	GetAliyunUidList() []*string
	SetEndDate(v string) *QueryAmountRequest
	GetEndDate() *string
	SetStartDate(v string) *QueryAmountRequest
	GetStartDate() *string
}

type QueryAmountRequest struct {
	// This parameter is required.
	AliyunUidList []*string `json:"aliyunUidList,omitempty" xml:"aliyunUidList,omitempty" type:"Repeated"`
	// This parameter is required.
	EndDate *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// This parameter is required.
	StartDate *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
}

func (s QueryAmountRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAmountRequest) GoString() string {
	return s.String()
}

func (s *QueryAmountRequest) GetAliyunUidList() []*string {
	return s.AliyunUidList
}

func (s *QueryAmountRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *QueryAmountRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *QueryAmountRequest) SetAliyunUidList(v []*string) *QueryAmountRequest {
	s.AliyunUidList = v
	return s
}

func (s *QueryAmountRequest) SetEndDate(v string) *QueryAmountRequest {
	s.EndDate = &v
	return s
}

func (s *QueryAmountRequest) SetStartDate(v string) *QueryAmountRequest {
	s.StartDate = &v
	return s
}

func (s *QueryAmountRequest) Validate() error {
	return dara.Validate(s)
}
