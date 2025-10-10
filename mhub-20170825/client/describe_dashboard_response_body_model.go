// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDashboardResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetModel(v string) *DescribeDashboardResponseBody
	GetModel() *string
	SetRequestId(v string) *DescribeDashboardResponseBody
	GetRequestId() *string
}

type DescribeDashboardResponseBody struct {
	// example:
	//
	// {
	//
	// 	"success":true,
	//
	// 	"model":{
	//
	// 		"perfMonthCount":0,
	//
	// 		"compatibilityMonthCount":0,
	//
	// 		"perfLastMonthCount":0,
	//
	// 		"compatibilityLastMonthCount":0,
	//
	// 		"automationMonthCount":0,
	//
	// 		"automationLastMonthCount":0
	//
	// 	}
	//
	// }
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// 4CC30A8F-787C-55CA-87A6-7D1BED56067E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDashboardResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDashboardResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDashboardResponseBody) GetModel() *string {
	return s.Model
}

func (s *DescribeDashboardResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDashboardResponseBody) SetModel(v string) *DescribeDashboardResponseBody {
	s.Model = &v
	return s
}

func (s *DescribeDashboardResponseBody) SetRequestId(v string) *DescribeDashboardResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDashboardResponseBody) Validate() error {
	return dara.Validate(s)
}
