// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSiteMonitorStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeSiteMonitorStatisticsResponseBody
	GetCode() *string
	SetData(v string) *DescribeSiteMonitorStatisticsResponseBody
	GetData() *string
	SetMessage(v string) *DescribeSiteMonitorStatisticsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeSiteMonitorStatisticsResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeSiteMonitorStatisticsResponseBody
	GetSuccess() *string
}

type DescribeSiteMonitorStatisticsResponseBody struct {
	// The responses code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The statistics.
	//
	// example:
	//
	// 100
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Succcessful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3AD2724D-E317-4BFB-B422-D6691D071BE1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSiteMonitorStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorStatisticsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeSiteMonitorStatisticsResponseBody) GetData() *string {
	return s.Data
}

func (s *DescribeSiteMonitorStatisticsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeSiteMonitorStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSiteMonitorStatisticsResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeSiteMonitorStatisticsResponseBody) SetCode(v string) *DescribeSiteMonitorStatisticsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSiteMonitorStatisticsResponseBody) SetData(v string) *DescribeSiteMonitorStatisticsResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeSiteMonitorStatisticsResponseBody) SetMessage(v string) *DescribeSiteMonitorStatisticsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSiteMonitorStatisticsResponseBody) SetRequestId(v string) *DescribeSiteMonitorStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSiteMonitorStatisticsResponseBody) SetSuccess(v string) *DescribeSiteMonitorStatisticsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSiteMonitorStatisticsResponseBody) Validate() error {
	return dara.Validate(s)
}
