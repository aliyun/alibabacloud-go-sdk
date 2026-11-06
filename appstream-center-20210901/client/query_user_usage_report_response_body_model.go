// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUserUsageReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryUserUsageReportResponseBody
	GetCode() *string
	SetCount(v int32) *QueryUserUsageReportResponseBody
	GetCount() *int32
	SetMessage(v string) *QueryUserUsageReportResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *QueryUserUsageReportResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *QueryUserUsageReportResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *QueryUserUsageReportResponseBody
	GetRequestId() *string
	SetUserUsageReportDataModel(v []*QueryUserUsageReportResponseBodyUserUsageReportDataModel) *QueryUserUsageReportResponseBody
	GetUserUsageReportDataModel() []*QueryUserUsageReportResponseBodyUserUsageReportDataModel
}

type QueryUserUsageReportResponseBody struct {
	// The business status code. A value of 200 indicates that the request is successful. Other values indicate that the request failed.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The total number of records that meet the query conditions.
	//
	// example:
	//
	// 3
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The description of the request processing result. Typically, success is returned when the request is successful.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number corresponding to the current response.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page used in the current response.
	//
	// example:
	//
	// 200
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The unique identifier of the request. You can use this ID to troubleshoot request issues.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of user usage data, sorted in descending order of cumulative usage duration.
	UserUsageReportDataModel []*QueryUserUsageReportResponseBodyUserUsageReportDataModel `json:"UserUsageReportDataModel,omitempty" xml:"UserUsageReportDataModel,omitempty" type:"Repeated"`
}

func (s QueryUserUsageReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryUserUsageReportResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserUsageReportResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryUserUsageReportResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *QueryUserUsageReportResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryUserUsageReportResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *QueryUserUsageReportResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryUserUsageReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryUserUsageReportResponseBody) GetUserUsageReportDataModel() []*QueryUserUsageReportResponseBodyUserUsageReportDataModel {
	return s.UserUsageReportDataModel
}

func (s *QueryUserUsageReportResponseBody) SetCode(v string) *QueryUserUsageReportResponseBody {
	s.Code = &v
	return s
}

func (s *QueryUserUsageReportResponseBody) SetCount(v int32) *QueryUserUsageReportResponseBody {
	s.Count = &v
	return s
}

func (s *QueryUserUsageReportResponseBody) SetMessage(v string) *QueryUserUsageReportResponseBody {
	s.Message = &v
	return s
}

func (s *QueryUserUsageReportResponseBody) SetPageNumber(v int32) *QueryUserUsageReportResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryUserUsageReportResponseBody) SetPageSize(v int32) *QueryUserUsageReportResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryUserUsageReportResponseBody) SetRequestId(v string) *QueryUserUsageReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryUserUsageReportResponseBody) SetUserUsageReportDataModel(v []*QueryUserUsageReportResponseBodyUserUsageReportDataModel) *QueryUserUsageReportResponseBody {
	s.UserUsageReportDataModel = v
	return s
}

func (s *QueryUserUsageReportResponseBody) Validate() error {
	if s.UserUsageReportDataModel != nil {
		for _, item := range s.UserUsageReportDataModel {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryUserUsageReportResponseBodyUserUsageReportDataModel struct {
	// The ID of the app instance group to which the user usage record belongs.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The name of the app instance group to which the user usage record belongs.
	//
	// example:
	//
	// R&D workstation group
	AppInstanceGroupName *string `json:"AppInstanceGroupName,omitempty" xml:"AppInstanceGroupName,omitempty"`
	// The current connection status of the user.
	//
	// Valid values:
	//
	// - `Connected`: At least one currently connected session exists.
	//
	// - `Disconnected`: The session binding is retained, but no currently connected session exists.
	//
	// - `Unbound`: No session binding exists.
	//
	// example:
	//
	// Connected
	ConnectionStatus *string `json:"ConnectionStatus,omitempty" xml:"ConnectionStatus,omitempty"`
	// The last connection time. Format: `yyyy-MM-dd HH:mm:ss`. If no current or historical connection record exists, - is returned. The string itself does not carry a time zone and is interpreted as Asia/Shanghai (UTC+8).
	//
	// example:
	//
	// 2026-09-01 10:00:00
	LastConnectedTime *string `json:"LastConnectedTime,omitempty" xml:"LastConnectedTime,omitempty"`
	// The current version does not return a valid region ID. Callers should not rely on this field. To filter by business region, use the BizRegionId request parameter.
	//
	// example:
	//
	// null
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The cumulative session persistence duration within the query range. This is the duration during which the user occupies the session from session attach to session release. Multiple sessions under the same app instance group are accumulated. Unit: hours.
	//
	// example:
	//
	// 3.5
	SessionsHoldDuration *float64 `json:"SessionsHoldDuration,omitempty" xml:"SessionsHoldDuration,omitempty"`
	// The cumulative user usage duration within the query range. The usage duration of the same user under the same app instance group is aggregated by day. Unit: hours.
	//
	// example:
	//
	// 2
	UsageDuration *float64 `json:"UsageDuration,omitempty" xml:"UsageDuration,omitempty"`
	// The ID of the end user.
	//
	// example:
	//
	// user01
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryUserUsageReportResponseBodyUserUsageReportDataModel) String() string {
	return dara.Prettify(s)
}

func (s QueryUserUsageReportResponseBodyUserUsageReportDataModel) GoString() string {
	return s.String()
}

func (s *QueryUserUsageReportResponseBodyUserUsageReportDataModel) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *QueryUserUsageReportResponseBodyUserUsageReportDataModel) GetAppInstanceGroupName() *string {
	return s.AppInstanceGroupName
}

func (s *QueryUserUsageReportResponseBodyUserUsageReportDataModel) GetConnectionStatus() *string {
	return s.ConnectionStatus
}

func (s *QueryUserUsageReportResponseBodyUserUsageReportDataModel) GetLastConnectedTime() *string {
	return s.LastConnectedTime
}

func (s *QueryUserUsageReportResponseBodyUserUsageReportDataModel) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryUserUsageReportResponseBodyUserUsageReportDataModel) GetSessionsHoldDuration() *float64 {
	return s.SessionsHoldDuration
}

func (s *QueryUserUsageReportResponseBodyUserUsageReportDataModel) GetUsageDuration() *float64 {
	return s.UsageDuration
}

func (s *QueryUserUsageReportResponseBodyUserUsageReportDataModel) GetUserId() *string {
	return s.UserId
}

func (s *QueryUserUsageReportResponseBodyUserUsageReportDataModel) SetAppInstanceGroupId(v string) *QueryUserUsageReportResponseBodyUserUsageReportDataModel {
	s.AppInstanceGroupId = &v
	return s
}

func (s *QueryUserUsageReportResponseBodyUserUsageReportDataModel) SetAppInstanceGroupName(v string) *QueryUserUsageReportResponseBodyUserUsageReportDataModel {
	s.AppInstanceGroupName = &v
	return s
}

func (s *QueryUserUsageReportResponseBodyUserUsageReportDataModel) SetConnectionStatus(v string) *QueryUserUsageReportResponseBodyUserUsageReportDataModel {
	s.ConnectionStatus = &v
	return s
}

func (s *QueryUserUsageReportResponseBodyUserUsageReportDataModel) SetLastConnectedTime(v string) *QueryUserUsageReportResponseBodyUserUsageReportDataModel {
	s.LastConnectedTime = &v
	return s
}

func (s *QueryUserUsageReportResponseBodyUserUsageReportDataModel) SetRegionId(v string) *QueryUserUsageReportResponseBodyUserUsageReportDataModel {
	s.RegionId = &v
	return s
}

func (s *QueryUserUsageReportResponseBodyUserUsageReportDataModel) SetSessionsHoldDuration(v float64) *QueryUserUsageReportResponseBodyUserUsageReportDataModel {
	s.SessionsHoldDuration = &v
	return s
}

func (s *QueryUserUsageReportResponseBodyUserUsageReportDataModel) SetUsageDuration(v float64) *QueryUserUsageReportResponseBodyUserUsageReportDataModel {
	s.UsageDuration = &v
	return s
}

func (s *QueryUserUsageReportResponseBodyUserUsageReportDataModel) SetUserId(v string) *QueryUserUsageReportResponseBodyUserUsageReportDataModel {
	s.UserId = &v
	return s
}

func (s *QueryUserUsageReportResponseBodyUserUsageReportDataModel) Validate() error {
	return dara.Validate(s)
}
