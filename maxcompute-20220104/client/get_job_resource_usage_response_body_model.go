// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobResourceUsageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetJobResourceUsageResponseBodyData) *GetJobResourceUsageResponseBody
	GetData() *GetJobResourceUsageResponseBodyData
	SetErrorCode(v string) *GetJobResourceUsageResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *GetJobResourceUsageResponseBody
	GetErrorMsg() *string
	SetHttpCode(v int32) *GetJobResourceUsageResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *GetJobResourceUsageResponseBody
	GetRequestId() *string
}

type GetJobResourceUsageResponseBody struct {
	// The data returned.
	Data *GetJobResourceUsageResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// success
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// 0A3B1E82006A23A918C70905BF08AEC7
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// Indicates whether the request was successful. If this parameter was not empty and the value of this parameter was not 200, the request failed.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0b57ff7616612271051086500ea3ce
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetJobResourceUsageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetJobResourceUsageResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobResourceUsageResponseBody) GetData() *GetJobResourceUsageResponseBodyData {
	return s.Data
}

func (s *GetJobResourceUsageResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetJobResourceUsageResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GetJobResourceUsageResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *GetJobResourceUsageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetJobResourceUsageResponseBody) SetData(v *GetJobResourceUsageResponseBodyData) *GetJobResourceUsageResponseBody {
	s.Data = v
	return s
}

func (s *GetJobResourceUsageResponseBody) SetErrorCode(v string) *GetJobResourceUsageResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetJobResourceUsageResponseBody) SetErrorMsg(v string) *GetJobResourceUsageResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetJobResourceUsageResponseBody) SetHttpCode(v int32) *GetJobResourceUsageResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetJobResourceUsageResponseBody) SetRequestId(v string) *GetJobResourceUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJobResourceUsageResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetJobResourceUsageResponseBodyData struct {
	// The data list returned.
	JobResourceUsageList []*GetJobResourceUsageResponseBodyDataJobResourceUsageList `json:"jobResourceUsageList,omitempty" xml:"jobResourceUsageList,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 2
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 64
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetJobResourceUsageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetJobResourceUsageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetJobResourceUsageResponseBodyData) GetJobResourceUsageList() []*GetJobResourceUsageResponseBodyDataJobResourceUsageList {
	return s.JobResourceUsageList
}

func (s *GetJobResourceUsageResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *GetJobResourceUsageResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetJobResourceUsageResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetJobResourceUsageResponseBodyData) SetJobResourceUsageList(v []*GetJobResourceUsageResponseBodyDataJobResourceUsageList) *GetJobResourceUsageResponseBodyData {
	s.JobResourceUsageList = v
	return s
}

func (s *GetJobResourceUsageResponseBodyData) SetPageNumber(v int64) *GetJobResourceUsageResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *GetJobResourceUsageResponseBodyData) SetPageSize(v int64) *GetJobResourceUsageResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetJobResourceUsageResponseBodyData) SetTotalCount(v int64) *GetJobResourceUsageResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *GetJobResourceUsageResponseBodyData) Validate() error {
	if s.JobResourceUsageList != nil {
		for _, item := range s.JobResourceUsageList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetJobResourceUsageResponseBodyDataJobResourceUsageList struct {
	// The total number of used compute units (CUs).
	//
	// example:
	//
	// 1185100
	CuUsage *int64 `json:"cuUsage,omitempty" xml:"cuUsage,omitempty"`
	// The start date of the query in the format of yyyy-MM-dd.
	//
	// example:
	//
	// 2023-05-09
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
	// The job executor.
	//
	// example:
	//
	// ALIYUN$xxx@test.aliyunid.com
	JobOwner *string `json:"jobOwner,omitempty" xml:"jobOwner,omitempty"`
	// The total memory usage.
	//
	// example:
	//
	// 15169536
	MemoryUsage *int64 `json:"memoryUsage,omitempty" xml:"memoryUsage,omitempty"`
	// The quota nickname.
	//
	// example:
	//
	// my_quota
	QuotaNickname *string `json:"quotaNickname,omitempty" xml:"quotaNickname,omitempty"`
}

func (s GetJobResourceUsageResponseBodyDataJobResourceUsageList) String() string {
	return dara.Prettify(s)
}

func (s GetJobResourceUsageResponseBodyDataJobResourceUsageList) GoString() string {
	return s.String()
}

func (s *GetJobResourceUsageResponseBodyDataJobResourceUsageList) GetCuUsage() *int64 {
	return s.CuUsage
}

func (s *GetJobResourceUsageResponseBodyDataJobResourceUsageList) GetDate() *string {
	return s.Date
}

func (s *GetJobResourceUsageResponseBodyDataJobResourceUsageList) GetJobOwner() *string {
	return s.JobOwner
}

func (s *GetJobResourceUsageResponseBodyDataJobResourceUsageList) GetMemoryUsage() *int64 {
	return s.MemoryUsage
}

func (s *GetJobResourceUsageResponseBodyDataJobResourceUsageList) GetQuotaNickname() *string {
	return s.QuotaNickname
}

func (s *GetJobResourceUsageResponseBodyDataJobResourceUsageList) SetCuUsage(v int64) *GetJobResourceUsageResponseBodyDataJobResourceUsageList {
	s.CuUsage = &v
	return s
}

func (s *GetJobResourceUsageResponseBodyDataJobResourceUsageList) SetDate(v string) *GetJobResourceUsageResponseBodyDataJobResourceUsageList {
	s.Date = &v
	return s
}

func (s *GetJobResourceUsageResponseBodyDataJobResourceUsageList) SetJobOwner(v string) *GetJobResourceUsageResponseBodyDataJobResourceUsageList {
	s.JobOwner = &v
	return s
}

func (s *GetJobResourceUsageResponseBodyDataJobResourceUsageList) SetMemoryUsage(v int64) *GetJobResourceUsageResponseBodyDataJobResourceUsageList {
	s.MemoryUsage = &v
	return s
}

func (s *GetJobResourceUsageResponseBodyDataJobResourceUsageList) SetQuotaNickname(v string) *GetJobResourceUsageResponseBodyDataJobResourceUsageList {
	s.QuotaNickname = &v
	return s
}

func (s *GetJobResourceUsageResponseBodyDataJobResourceUsageList) Validate() error {
	return dara.Validate(s)
}
