// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPageServiceCustomerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PageServiceCustomerResponseBody
	GetCode() *string
	SetData(v []*PageServiceCustomerResponseBodyData) *PageServiceCustomerResponseBody
	GetData() []*PageServiceCustomerResponseBodyData
	SetHttpStatusCode(v int32) *PageServiceCustomerResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *PageServiceCustomerResponseBody
	GetMessage() *string
	SetPageInfo(v *PageServiceCustomerResponseBodyPageInfo) *PageServiceCustomerResponseBody
	GetPageInfo() *PageServiceCustomerResponseBodyPageInfo
	SetRequestId(v string) *PageServiceCustomerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PageServiceCustomerResponseBody
	GetSuccess() *bool
}

type PageServiceCustomerResponseBody struct {
	// Interface return code.
	//
	// example:
	//
	// System error or openapi error
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data query results.
	Data []*PageServiceCustomerResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Return message. When the request is successful, it returns a success message; when the request fails, it returns the reason for the failure.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Pagination information.
	PageInfo *PageServiceCustomerResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// Request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 808A307F-9513-5099-AAA5-98D4EF199140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Request return status.
	//
	// - true: Success.
	//
	// - false: Failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PageServiceCustomerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PageServiceCustomerResponseBody) GoString() string {
	return s.String()
}

func (s *PageServiceCustomerResponseBody) GetCode() *string {
	return s.Code
}

func (s *PageServiceCustomerResponseBody) GetData() []*PageServiceCustomerResponseBodyData {
	return s.Data
}

func (s *PageServiceCustomerResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *PageServiceCustomerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PageServiceCustomerResponseBody) GetPageInfo() *PageServiceCustomerResponseBodyPageInfo {
	return s.PageInfo
}

func (s *PageServiceCustomerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PageServiceCustomerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PageServiceCustomerResponseBody) SetCode(v string) *PageServiceCustomerResponseBody {
	s.Code = &v
	return s
}

func (s *PageServiceCustomerResponseBody) SetData(v []*PageServiceCustomerResponseBodyData) *PageServiceCustomerResponseBody {
	s.Data = v
	return s
}

func (s *PageServiceCustomerResponseBody) SetHttpStatusCode(v int32) *PageServiceCustomerResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PageServiceCustomerResponseBody) SetMessage(v string) *PageServiceCustomerResponseBody {
	s.Message = &v
	return s
}

func (s *PageServiceCustomerResponseBody) SetPageInfo(v *PageServiceCustomerResponseBodyPageInfo) *PageServiceCustomerResponseBody {
	s.PageInfo = v
	return s
}

func (s *PageServiceCustomerResponseBody) SetRequestId(v string) *PageServiceCustomerResponseBody {
	s.RequestId = &v
	return s
}

func (s *PageServiceCustomerResponseBody) SetSuccess(v bool) *PageServiceCustomerResponseBody {
	s.Success = &v
	return s
}

func (s *PageServiceCustomerResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PageServiceCustomerResponseBodyData struct {
	// Customer UID.
	//
	// example:
	//
	// 1667751131382856
	Aliuid *string `json:"Aliuid,omitempty" xml:"Aliuid,omitempty"`
	// Authorization status.
	//
	// example:
	//
	// 1
	AuthStatus *int32 `json:"AuthStatus,omitempty" xml:"AuthStatus,omitempty"`
	// Cloud Monitoring - Alert authorization status.
	//
	// example:
	//
	// 0
	CmAuthStatus *int32 `json:"CmAuthStatus,omitempty" xml:"CmAuthStatus,omitempty"`
	// End time. The format is a Unix timestamp, which is the number of milliseconds since January 1, 1970.
	//
	// example:
	//
	// 1710123149222
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Customer level.
	//
	// example:
	//
	// GC1
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// Cloud Security - Alert authorization status.
	//
	// example:
	//
	// 1
	MonitorAuthStatus *int32 `json:"MonitorAuthStatus,omitempty" xml:"MonitorAuthStatus,omitempty"`
	// Customer name.
	//
	// example:
	//
	// 中国工程院
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Owner name.
	//
	// example:
	//
	// 常温
	OwnId *string `json:"OwnId,omitempty" xml:"OwnId,omitempty"`
	// Start time. The format is a Unix timestamp, which is the number of milliseconds since January 1, 1970.
	//
	// example:
	//
	// 1710123149000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Customer ID.
	//
	// example:
	//
	// 1667751131382856
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// Version information.
	//
	// example:
	//
	// 企业版
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s PageServiceCustomerResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s PageServiceCustomerResponseBodyData) GoString() string {
	return s.String()
}

func (s *PageServiceCustomerResponseBodyData) GetAliuid() *string {
	return s.Aliuid
}

func (s *PageServiceCustomerResponseBodyData) GetAuthStatus() *int32 {
	return s.AuthStatus
}

func (s *PageServiceCustomerResponseBodyData) GetCmAuthStatus() *int32 {
	return s.CmAuthStatus
}

func (s *PageServiceCustomerResponseBodyData) GetEndTime() *int64 {
	return s.EndTime
}

func (s *PageServiceCustomerResponseBodyData) GetLevel() *string {
	return s.Level
}

func (s *PageServiceCustomerResponseBodyData) GetMonitorAuthStatus() *int32 {
	return s.MonitorAuthStatus
}

func (s *PageServiceCustomerResponseBodyData) GetName() *string {
	return s.Name
}

func (s *PageServiceCustomerResponseBodyData) GetOwnId() *string {
	return s.OwnId
}

func (s *PageServiceCustomerResponseBodyData) GetStartTime() *int64 {
	return s.StartTime
}

func (s *PageServiceCustomerResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *PageServiceCustomerResponseBodyData) GetVersion() *string {
	return s.Version
}

func (s *PageServiceCustomerResponseBodyData) SetAliuid(v string) *PageServiceCustomerResponseBodyData {
	s.Aliuid = &v
	return s
}

func (s *PageServiceCustomerResponseBodyData) SetAuthStatus(v int32) *PageServiceCustomerResponseBodyData {
	s.AuthStatus = &v
	return s
}

func (s *PageServiceCustomerResponseBodyData) SetCmAuthStatus(v int32) *PageServiceCustomerResponseBodyData {
	s.CmAuthStatus = &v
	return s
}

func (s *PageServiceCustomerResponseBodyData) SetEndTime(v int64) *PageServiceCustomerResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *PageServiceCustomerResponseBodyData) SetLevel(v string) *PageServiceCustomerResponseBodyData {
	s.Level = &v
	return s
}

func (s *PageServiceCustomerResponseBodyData) SetMonitorAuthStatus(v int32) *PageServiceCustomerResponseBodyData {
	s.MonitorAuthStatus = &v
	return s
}

func (s *PageServiceCustomerResponseBodyData) SetName(v string) *PageServiceCustomerResponseBodyData {
	s.Name = &v
	return s
}

func (s *PageServiceCustomerResponseBodyData) SetOwnId(v string) *PageServiceCustomerResponseBodyData {
	s.OwnId = &v
	return s
}

func (s *PageServiceCustomerResponseBodyData) SetStartTime(v int64) *PageServiceCustomerResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *PageServiceCustomerResponseBodyData) SetUserId(v string) *PageServiceCustomerResponseBodyData {
	s.UserId = &v
	return s
}

func (s *PageServiceCustomerResponseBodyData) SetVersion(v string) *PageServiceCustomerResponseBodyData {
	s.Version = &v
	return s
}

func (s *PageServiceCustomerResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type PageServiceCustomerResponseBodyPageInfo struct {
	// The current page number in pagination queries.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Number of items per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Total number of query results.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s PageServiceCustomerResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s PageServiceCustomerResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *PageServiceCustomerResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *PageServiceCustomerResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *PageServiceCustomerResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *PageServiceCustomerResponseBodyPageInfo) SetCurrentPage(v int32) *PageServiceCustomerResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *PageServiceCustomerResponseBodyPageInfo) SetPageSize(v int32) *PageServiceCustomerResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *PageServiceCustomerResponseBodyPageInfo) SetTotalCount(v int32) *PageServiceCustomerResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *PageServiceCustomerResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
