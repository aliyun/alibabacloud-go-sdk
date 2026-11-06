// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUserAppUsageAggregateByHourResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryUserAppUsageAggregateByHourResponseBody
	GetCode() *string
	SetIndexData(v []*QueryUserAppUsageAggregateByHourResponseBodyIndexData) *QueryUserAppUsageAggregateByHourResponseBody
	GetIndexData() []*QueryUserAppUsageAggregateByHourResponseBodyIndexData
	SetMessage(v string) *QueryUserAppUsageAggregateByHourResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryUserAppUsageAggregateByHourResponseBody
	GetRequestId() *string
}

type QueryUserAppUsageAggregateByHourResponseBody struct {
	// The business status code. A value of 200 indicates that the request is successful. Other values indicate that the request failed.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The list of hourly aggregated user application usage data. A maximum of five records are returned for each hour.
	IndexData []*QueryUserAppUsageAggregateByHourResponseBodyIndexData `json:"IndexData,omitempty" xml:"IndexData,omitempty" type:"Repeated"`
	// The description of the request processing result. A value of success is usually returned when the request is successful.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique identifier of the request, which can be used to troubleshoot request issues.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryUserAppUsageAggregateByHourResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryUserAppUsageAggregateByHourResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserAppUsageAggregateByHourResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryUserAppUsageAggregateByHourResponseBody) GetIndexData() []*QueryUserAppUsageAggregateByHourResponseBodyIndexData {
	return s.IndexData
}

func (s *QueryUserAppUsageAggregateByHourResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryUserAppUsageAggregateByHourResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryUserAppUsageAggregateByHourResponseBody) SetCode(v string) *QueryUserAppUsageAggregateByHourResponseBody {
	s.Code = &v
	return s
}

func (s *QueryUserAppUsageAggregateByHourResponseBody) SetIndexData(v []*QueryUserAppUsageAggregateByHourResponseBodyIndexData) *QueryUserAppUsageAggregateByHourResponseBody {
	s.IndexData = v
	return s
}

func (s *QueryUserAppUsageAggregateByHourResponseBody) SetMessage(v string) *QueryUserAppUsageAggregateByHourResponseBody {
	s.Message = &v
	return s
}

func (s *QueryUserAppUsageAggregateByHourResponseBody) SetRequestId(v string) *QueryUserAppUsageAggregateByHourResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryUserAppUsageAggregateByHourResponseBody) Validate() error {
	if s.IndexData != nil {
		for _, item := range s.IndexData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryUserAppUsageAggregateByHourResponseBodyIndexData struct {
	// The application ID.
	//
	// example:
	//
	// app-9ciijz60n4xsv****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application name.
	//
	// example:
	//
	// Visual Studio Code
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The cumulative usage duration of the application within the statistical range. Unit: seconds.
	//
	// example:
	//
	// 200
	AppTimeCnt *string `json:"AppTimeCnt,omitempty" xml:"AppTimeCnt,omitempty"`
	// The statistical hour in a two-digit 24-hour format string. Valid values: 00 to 23. Hour boundaries are divided based on the Asia/Shanghai (UTC+8) time zone.
	//
	// example:
	//
	// 02
	Hh *string `json:"Hh,omitempty" xml:"Hh,omitempty"`
}

func (s QueryUserAppUsageAggregateByHourResponseBodyIndexData) String() string {
	return dara.Prettify(s)
}

func (s QueryUserAppUsageAggregateByHourResponseBodyIndexData) GoString() string {
	return s.String()
}

func (s *QueryUserAppUsageAggregateByHourResponseBodyIndexData) GetAppId() *string {
	return s.AppId
}

func (s *QueryUserAppUsageAggregateByHourResponseBodyIndexData) GetAppName() *string {
	return s.AppName
}

func (s *QueryUserAppUsageAggregateByHourResponseBodyIndexData) GetAppTimeCnt() *string {
	return s.AppTimeCnt
}

func (s *QueryUserAppUsageAggregateByHourResponseBodyIndexData) GetHh() *string {
	return s.Hh
}

func (s *QueryUserAppUsageAggregateByHourResponseBodyIndexData) SetAppId(v string) *QueryUserAppUsageAggregateByHourResponseBodyIndexData {
	s.AppId = &v
	return s
}

func (s *QueryUserAppUsageAggregateByHourResponseBodyIndexData) SetAppName(v string) *QueryUserAppUsageAggregateByHourResponseBodyIndexData {
	s.AppName = &v
	return s
}

func (s *QueryUserAppUsageAggregateByHourResponseBodyIndexData) SetAppTimeCnt(v string) *QueryUserAppUsageAggregateByHourResponseBodyIndexData {
	s.AppTimeCnt = &v
	return s
}

func (s *QueryUserAppUsageAggregateByHourResponseBodyIndexData) SetHh(v string) *QueryUserAppUsageAggregateByHourResponseBodyIndexData {
	s.Hh = &v
	return s
}

func (s *QueryUserAppUsageAggregateByHourResponseBodyIndexData) Validate() error {
	return dara.Validate(s)
}
