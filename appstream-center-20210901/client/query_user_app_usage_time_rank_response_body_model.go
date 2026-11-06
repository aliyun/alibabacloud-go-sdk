// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUserAppUsageTimeRankResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryUserAppUsageTimeRankResponseBody
	GetCode() *string
	SetIndexData(v []*QueryUserAppUsageTimeRankResponseBodyIndexData) *QueryUserAppUsageTimeRankResponseBody
	GetIndexData() []*QueryUserAppUsageTimeRankResponseBodyIndexData
	SetMessage(v string) *QueryUserAppUsageTimeRankResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryUserAppUsageTimeRankResponseBody
	GetRequestId() *string
}

type QueryUserAppUsageTimeRankResponseBody struct {
	// The business status code. A value of `200` indicates that the request is processed successfully. Other values indicate that the processing failed.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ranking list of user application usage duration, sorted in descending order of cumulative usage duration.
	IndexData []*QueryUserAppUsageTimeRankResponseBodyIndexData `json:"IndexData,omitempty" xml:"IndexData,omitempty" type:"Repeated"`
	// The description of the request processing result. If the request is successful, `success` is usually returned.
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

func (s QueryUserAppUsageTimeRankResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryUserAppUsageTimeRankResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserAppUsageTimeRankResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryUserAppUsageTimeRankResponseBody) GetIndexData() []*QueryUserAppUsageTimeRankResponseBodyIndexData {
	return s.IndexData
}

func (s *QueryUserAppUsageTimeRankResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryUserAppUsageTimeRankResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryUserAppUsageTimeRankResponseBody) SetCode(v string) *QueryUserAppUsageTimeRankResponseBody {
	s.Code = &v
	return s
}

func (s *QueryUserAppUsageTimeRankResponseBody) SetIndexData(v []*QueryUserAppUsageTimeRankResponseBodyIndexData) *QueryUserAppUsageTimeRankResponseBody {
	s.IndexData = v
	return s
}

func (s *QueryUserAppUsageTimeRankResponseBody) SetMessage(v string) *QueryUserAppUsageTimeRankResponseBody {
	s.Message = &v
	return s
}

func (s *QueryUserAppUsageTimeRankResponseBody) SetRequestId(v string) *QueryUserAppUsageTimeRankResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryUserAppUsageTimeRankResponseBody) Validate() error {
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

type QueryUserAppUsageTimeRankResponseBodyIndexData struct {
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
	// The application version. The current version may not return this value, and callers should not rely on this field. A non-empty value indicates the application version.
	//
	// example:
	//
	// 1.96.0
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
}

func (s QueryUserAppUsageTimeRankResponseBodyIndexData) String() string {
	return dara.Prettify(s)
}

func (s QueryUserAppUsageTimeRankResponseBodyIndexData) GoString() string {
	return s.String()
}

func (s *QueryUserAppUsageTimeRankResponseBodyIndexData) GetAppId() *string {
	return s.AppId
}

func (s *QueryUserAppUsageTimeRankResponseBodyIndexData) GetAppName() *string {
	return s.AppName
}

func (s *QueryUserAppUsageTimeRankResponseBodyIndexData) GetAppTimeCnt() *string {
	return s.AppTimeCnt
}

func (s *QueryUserAppUsageTimeRankResponseBodyIndexData) GetAppVersion() *string {
	return s.AppVersion
}

func (s *QueryUserAppUsageTimeRankResponseBodyIndexData) SetAppId(v string) *QueryUserAppUsageTimeRankResponseBodyIndexData {
	s.AppId = &v
	return s
}

func (s *QueryUserAppUsageTimeRankResponseBodyIndexData) SetAppName(v string) *QueryUserAppUsageTimeRankResponseBodyIndexData {
	s.AppName = &v
	return s
}

func (s *QueryUserAppUsageTimeRankResponseBodyIndexData) SetAppTimeCnt(v string) *QueryUserAppUsageTimeRankResponseBodyIndexData {
	s.AppTimeCnt = &v
	return s
}

func (s *QueryUserAppUsageTimeRankResponseBodyIndexData) SetAppVersion(v string) *QueryUserAppUsageTimeRankResponseBodyIndexData {
	s.AppVersion = &v
	return s
}

func (s *QueryUserAppUsageTimeRankResponseBodyIndexData) Validate() error {
	return dara.Validate(s)
}
