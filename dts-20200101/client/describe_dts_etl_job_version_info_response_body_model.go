// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDtsEtlJobVersionInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDtsEtlJobVersionInfos(v []*DescribeDtsEtlJobVersionInfoResponseBodyDtsEtlJobVersionInfos) *DescribeDtsEtlJobVersionInfoResponseBody
	GetDtsEtlJobVersionInfos() []*DescribeDtsEtlJobVersionInfoResponseBodyDtsEtlJobVersionInfos
	SetDynamicCode(v string) *DescribeDtsEtlJobVersionInfoResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DescribeDtsEtlJobVersionInfoResponseBody
	GetDynamicMessage() *string
	SetErrCode(v string) *DescribeDtsEtlJobVersionInfoResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeDtsEtlJobVersionInfoResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v string) *DescribeDtsEtlJobVersionInfoResponseBody
	GetHttpStatusCode() *string
	SetPageNumber(v int32) *DescribeDtsEtlJobVersionInfoResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeDtsEtlJobVersionInfoResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeDtsEtlJobVersionInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeDtsEtlJobVersionInfoResponseBody
	GetSuccess() *bool
	SetTotalRecordCount(v int32) *DescribeDtsEtlJobVersionInfoResponseBody
	GetTotalRecordCount() *int32
}

type DescribeDtsEtlJobVersionInfoResponseBody struct {
	// The details of ETL tasks.
	DtsEtlJobVersionInfos []*DescribeDtsEtlJobVersionInfoResponseBodyDtsEtlJobVersionInfos `json:"DtsEtlJobVersionInfos,omitempty" xml:"DtsEtlJobVersionInfos,omitempty" type:"Repeated"`
	// The dynamic error code.
	//
	// example:
	//
	// 403
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic part in the error message.
	//
	// example:
	//
	// present environment is not support,so skip
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error code. This example indicates that the specified ETL task ID is invalid.
	//
	// example:
	//
	// InvalidJobId
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message. This example indicates that the specified ETL task ID does not exist. In this case, the ETL task may have been deleted.
	//
	// example:
	//
	// The specified dts job id %s is not exists.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The page number of the returned page. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records returned on the current page.
	//
	// example:
	//
	// 20
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 224DB9F7-3100-4899-AB9C-C938BCCB43E7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. If the call failed, false is returned.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 200
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeDtsEtlJobVersionInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsEtlJobVersionInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDtsEtlJobVersionInfoResponseBody) GetDtsEtlJobVersionInfos() []*DescribeDtsEtlJobVersionInfoResponseBodyDtsEtlJobVersionInfos {
	return s.DtsEtlJobVersionInfos
}

func (s *DescribeDtsEtlJobVersionInfoResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DescribeDtsEtlJobVersionInfoResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DescribeDtsEtlJobVersionInfoResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeDtsEtlJobVersionInfoResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeDtsEtlJobVersionInfoResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DescribeDtsEtlJobVersionInfoResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDtsEtlJobVersionInfoResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeDtsEtlJobVersionInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDtsEtlJobVersionInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDtsEtlJobVersionInfoResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeDtsEtlJobVersionInfoResponseBody) SetDtsEtlJobVersionInfos(v []*DescribeDtsEtlJobVersionInfoResponseBodyDtsEtlJobVersionInfos) *DescribeDtsEtlJobVersionInfoResponseBody {
	s.DtsEtlJobVersionInfos = v
	return s
}

func (s *DescribeDtsEtlJobVersionInfoResponseBody) SetDynamicCode(v string) *DescribeDtsEtlJobVersionInfoResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeDtsEtlJobVersionInfoResponseBody) SetDynamicMessage(v string) *DescribeDtsEtlJobVersionInfoResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeDtsEtlJobVersionInfoResponseBody) SetErrCode(v string) *DescribeDtsEtlJobVersionInfoResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeDtsEtlJobVersionInfoResponseBody) SetErrMessage(v string) *DescribeDtsEtlJobVersionInfoResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeDtsEtlJobVersionInfoResponseBody) SetHttpStatusCode(v string) *DescribeDtsEtlJobVersionInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeDtsEtlJobVersionInfoResponseBody) SetPageNumber(v int32) *DescribeDtsEtlJobVersionInfoResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDtsEtlJobVersionInfoResponseBody) SetPageRecordCount(v int32) *DescribeDtsEtlJobVersionInfoResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeDtsEtlJobVersionInfoResponseBody) SetRequestId(v string) *DescribeDtsEtlJobVersionInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDtsEtlJobVersionInfoResponseBody) SetSuccess(v bool) *DescribeDtsEtlJobVersionInfoResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDtsEtlJobVersionInfoResponseBody) SetTotalRecordCount(v int32) *DescribeDtsEtlJobVersionInfoResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeDtsEtlJobVersionInfoResponseBody) Validate() error {
	if s.DtsEtlJobVersionInfos != nil {
		for _, item := range s.DtsEtlJobVersionInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDtsEtlJobVersionInfoResponseBodyDtsEtlJobVersionInfos struct {
	// The time when the ETL task was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1637229315000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the creator.
	//
	// example:
	//
	// 10000000
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The username of the creator.
	//
	// example:
	//
	// ***@****.com
	CreatorName *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	// The ID of the DTS instance.
	//
	// example:
	//
	// dtsg******gd
	DtsInstanceId *string `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
	// The ID of the ETL task.
	//
	// example:
	//
	// l5512es7w15****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The name of the ETL task.
	//
	// example:
	//
	// test_sql
	DtsJobName *string `json:"DtsJobName,omitempty" xml:"DtsJobName,omitempty"`
	// The time when the ETL task was last modified. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1637230117000
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The safe checkpoint of the ETL task.
	//
	// example:
	//
	// 1637230117000
	SafeCheckpoint *string `json:"SafeCheckpoint,omitempty" xml:"SafeCheckpoint,omitempty"`
	// The log level. Valid values: ERROR, WARN, INFO, and DEBUG.
	//
	// example:
	//
	// INFO
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The version number of the ETL task.
	//
	// example:
	//
	// 9
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeDtsEtlJobVersionInfoResponseBodyDtsEtlJobVersionInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsEtlJobVersionInfoResponseBodyDtsEtlJobVersionInfos) GoString() string {
	return s.String()
}

func (s *DescribeDtsEtlJobVersionInfoResponseBodyDtsEtlJobVersionInfos) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDtsEtlJobVersionInfoResponseBodyDtsEtlJobVersionInfos) GetCreator() *string {
	return s.Creator
}

func (s *DescribeDtsEtlJobVersionInfoResponseBodyDtsEtlJobVersionInfos) GetCreatorName() *string {
	return s.CreatorName
}

func (s *DescribeDtsEtlJobVersionInfoResponseBodyDtsEtlJobVersionInfos) GetDtsInstanceId() *string {
	return s.DtsInstanceId
}

func (s *DescribeDtsEtlJobVersionInfoResponseBodyDtsEtlJobVersionInfos) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *DescribeDtsEtlJobVersionInfoResponseBodyDtsEtlJobVersionInfos) GetDtsJobName() *string {
	return s.DtsJobName
}

func (s *DescribeDtsEtlJobVersionInfoResponseBodyDtsEtlJobVersionInfos) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *DescribeDtsEtlJobVersionInfoResponseBodyDtsEtlJobVersionInfos) GetSafeCheckpoint() *string {
	return s.SafeCheckpoint
}

func (s *DescribeDtsEtlJobVersionInfoResponseBodyDtsEtlJobVersionInfos) GetStatus() *string {
	return s.Status
}

func (s *DescribeDtsEtlJobVersionInfoResponseBodyDtsEtlJobVersionInfos) GetVersion() *int32 {
	return s.Version
}

func (s *DescribeDtsEtlJobVersionInfoResponseBodyDtsEtlJobVersionInfos) SetCreateTime(v string) *DescribeDtsEtlJobVersionInfoResponseBodyDtsEtlJobVersionInfos {
	s.CreateTime = &v
	return s
}

func (s *DescribeDtsEtlJobVersionInfoResponseBodyDtsEtlJobVersionInfos) SetCreator(v string) *DescribeDtsEtlJobVersionInfoResponseBodyDtsEtlJobVersionInfos {
	s.Creator = &v
	return s
}

func (s *DescribeDtsEtlJobVersionInfoResponseBodyDtsEtlJobVersionInfos) SetCreatorName(v string) *DescribeDtsEtlJobVersionInfoResponseBodyDtsEtlJobVersionInfos {
	s.CreatorName = &v
	return s
}

func (s *DescribeDtsEtlJobVersionInfoResponseBodyDtsEtlJobVersionInfos) SetDtsInstanceId(v string) *DescribeDtsEtlJobVersionInfoResponseBodyDtsEtlJobVersionInfos {
	s.DtsInstanceId = &v
	return s
}

func (s *DescribeDtsEtlJobVersionInfoResponseBodyDtsEtlJobVersionInfos) SetDtsJobId(v string) *DescribeDtsEtlJobVersionInfoResponseBodyDtsEtlJobVersionInfos {
	s.DtsJobId = &v
	return s
}

func (s *DescribeDtsEtlJobVersionInfoResponseBodyDtsEtlJobVersionInfos) SetDtsJobName(v string) *DescribeDtsEtlJobVersionInfoResponseBodyDtsEtlJobVersionInfos {
	s.DtsJobName = &v
	return s
}

func (s *DescribeDtsEtlJobVersionInfoResponseBodyDtsEtlJobVersionInfos) SetModifyTime(v string) *DescribeDtsEtlJobVersionInfoResponseBodyDtsEtlJobVersionInfos {
	s.ModifyTime = &v
	return s
}

func (s *DescribeDtsEtlJobVersionInfoResponseBodyDtsEtlJobVersionInfos) SetSafeCheckpoint(v string) *DescribeDtsEtlJobVersionInfoResponseBodyDtsEtlJobVersionInfos {
	s.SafeCheckpoint = &v
	return s
}

func (s *DescribeDtsEtlJobVersionInfoResponseBodyDtsEtlJobVersionInfos) SetStatus(v string) *DescribeDtsEtlJobVersionInfoResponseBodyDtsEtlJobVersionInfos {
	s.Status = &v
	return s
}

func (s *DescribeDtsEtlJobVersionInfoResponseBodyDtsEtlJobVersionInfos) SetVersion(v int32) *DescribeDtsEtlJobVersionInfoResponseBodyDtsEtlJobVersionInfos {
	s.Version = &v
	return s
}

func (s *DescribeDtsEtlJobVersionInfoResponseBodyDtsEtlJobVersionInfos) Validate() error {
	return dara.Validate(s)
}
