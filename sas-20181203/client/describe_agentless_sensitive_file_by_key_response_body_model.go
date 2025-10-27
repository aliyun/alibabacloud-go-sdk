// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAgentlessSensitiveFileByKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageInfo(v *DescribeAgentlessSensitiveFileByKeyResponseBodyPageInfo) *DescribeAgentlessSensitiveFileByKeyResponseBody
	GetPageInfo() *DescribeAgentlessSensitiveFileByKeyResponseBodyPageInfo
	SetRequestId(v string) *DescribeAgentlessSensitiveFileByKeyResponseBody
	GetRequestId() *string
	SetSensitiveFileList(v []*DescribeAgentlessSensitiveFileByKeyResponseBodySensitiveFileList) *DescribeAgentlessSensitiveFileByKeyResponseBody
	GetSensitiveFileList() []*DescribeAgentlessSensitiveFileByKeyResponseBodySensitiveFileList
	SetSuccess(v bool) *DescribeAgentlessSensitiveFileByKeyResponseBody
	GetSuccess() *bool
}

type DescribeAgentlessSensitiveFileByKeyResponseBody struct {
	// The pagination information.
	PageInfo *DescribeAgentlessSensitiveFileByKeyResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// FBBEB173-1F43-505F-A876-C03ECDF6CE4C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the sensitive files that are detected by using the agentless detection feature.
	SensitiveFileList []*DescribeAgentlessSensitiveFileByKeyResponseBodySensitiveFileList `json:"SensitiveFileList,omitempty" xml:"SensitiveFileList,omitempty" type:"Repeated"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAgentlessSensitiveFileByKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgentlessSensitiveFileByKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAgentlessSensitiveFileByKeyResponseBody) GetPageInfo() *DescribeAgentlessSensitiveFileByKeyResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeAgentlessSensitiveFileByKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAgentlessSensitiveFileByKeyResponseBody) GetSensitiveFileList() []*DescribeAgentlessSensitiveFileByKeyResponseBodySensitiveFileList {
	return s.SensitiveFileList
}

func (s *DescribeAgentlessSensitiveFileByKeyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeAgentlessSensitiveFileByKeyResponseBody) SetPageInfo(v *DescribeAgentlessSensitiveFileByKeyResponseBodyPageInfo) *DescribeAgentlessSensitiveFileByKeyResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeAgentlessSensitiveFileByKeyResponseBody) SetRequestId(v string) *DescribeAgentlessSensitiveFileByKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAgentlessSensitiveFileByKeyResponseBody) SetSensitiveFileList(v []*DescribeAgentlessSensitiveFileByKeyResponseBodySensitiveFileList) *DescribeAgentlessSensitiveFileByKeyResponseBody {
	s.SensitiveFileList = v
	return s
}

func (s *DescribeAgentlessSensitiveFileByKeyResponseBody) SetSuccess(v bool) *DescribeAgentlessSensitiveFileByKeyResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeAgentlessSensitiveFileByKeyResponseBody) Validate() error {
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	if s.SensitiveFileList != nil {
		for _, item := range s.SensitiveFileList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAgentlessSensitiveFileByKeyResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 20
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAgentlessSensitiveFileByKeyResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgentlessSensitiveFileByKeyResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeAgentlessSensitiveFileByKeyResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribeAgentlessSensitiveFileByKeyResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeAgentlessSensitiveFileByKeyResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAgentlessSensitiveFileByKeyResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeAgentlessSensitiveFileByKeyResponseBodyPageInfo) SetCount(v int32) *DescribeAgentlessSensitiveFileByKeyResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribeAgentlessSensitiveFileByKeyResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeAgentlessSensitiveFileByKeyResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAgentlessSensitiveFileByKeyResponseBodyPageInfo) SetPageSize(v int32) *DescribeAgentlessSensitiveFileByKeyResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeAgentlessSensitiveFileByKeyResponseBodyPageInfo) SetTotalCount(v int32) *DescribeAgentlessSensitiveFileByKeyResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeAgentlessSensitiveFileByKeyResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeAgentlessSensitiveFileByKeyResponseBodySensitiveFileList struct {
	// The path to the sensitive file.
	//
	// example:
	//
	// /usr/lib/abc.txt
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// The timestamp when the first scan was performed. Unit: milliseconds.
	//
	// example:
	//
	// 1663321552000
	FirstScanTime *int64 `json:"FirstScanTime,omitempty" xml:"FirstScanTime,omitempty"`
	// The ID of the alert for the sensitive file.
	//
	// example:
	//
	// 423505573
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The instance name of the asset.
	//
	// example:
	//
	// test-server
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The public IP address of the asset.
	//
	// example:
	//
	// 95.214.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address of the asset.
	//
	// example:
	//
	// 10.68.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The timestamp when the last scan was performed. Unit: milliseconds.
	//
	// example:
	//
	// 1663691592000
	LastScanTime *int64 `json:"LastScanTime,omitempty" xml:"LastScanTime,omitempty"`
	// The MD5 hash value of the sensitive file.
	//
	// example:
	//
	// b484b0dff093f358897486b58266****
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// The content of the sensitive file.
	//
	// example:
	//
	// mysql -u -p****
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// The risk level. Valid values:
	//
	// 	- **high**
	//
	// 	- **medium**
	//
	// 	- **low**
	//
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The type of the sensitive file.
	//
	// example:
	//
	// pem
	SensitiveFileKey *string `json:"SensitiveFileKey,omitempty" xml:"SensitiveFileKey,omitempty"`
	// The status of the sensitive file. Valid values:
	//
	// 	- **0**: unhandled.
	//
	// 	- **1**: ignored.
	//
	// 	- **2**: false positive.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the asset.
	//
	// example:
	//
	// SAS_Agentless_171869952****
	TargetName *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	// The UUID of the asset.
	//
	// example:
	//
	// 73e6a851-e63b-49ef-9d7a-5542d765****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeAgentlessSensitiveFileByKeyResponseBodySensitiveFileList) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgentlessSensitiveFileByKeyResponseBodySensitiveFileList) GoString() string {
	return s.String()
}

func (s *DescribeAgentlessSensitiveFileByKeyResponseBodySensitiveFileList) GetFilePath() *string {
	return s.FilePath
}

func (s *DescribeAgentlessSensitiveFileByKeyResponseBodySensitiveFileList) GetFirstScanTime() *int64 {
	return s.FirstScanTime
}

func (s *DescribeAgentlessSensitiveFileByKeyResponseBodySensitiveFileList) GetId() *int64 {
	return s.Id
}

func (s *DescribeAgentlessSensitiveFileByKeyResponseBodySensitiveFileList) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeAgentlessSensitiveFileByKeyResponseBodySensitiveFileList) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribeAgentlessSensitiveFileByKeyResponseBodySensitiveFileList) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *DescribeAgentlessSensitiveFileByKeyResponseBodySensitiveFileList) GetLastScanTime() *int64 {
	return s.LastScanTime
}

func (s *DescribeAgentlessSensitiveFileByKeyResponseBodySensitiveFileList) GetMd5() *string {
	return s.Md5
}

func (s *DescribeAgentlessSensitiveFileByKeyResponseBodySensitiveFileList) GetPrompt() *string {
	return s.Prompt
}

func (s *DescribeAgentlessSensitiveFileByKeyResponseBodySensitiveFileList) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeAgentlessSensitiveFileByKeyResponseBodySensitiveFileList) GetSensitiveFileKey() *string {
	return s.SensitiveFileKey
}

func (s *DescribeAgentlessSensitiveFileByKeyResponseBodySensitiveFileList) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeAgentlessSensitiveFileByKeyResponseBodySensitiveFileList) GetTargetName() *string {
	return s.TargetName
}

func (s *DescribeAgentlessSensitiveFileByKeyResponseBodySensitiveFileList) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeAgentlessSensitiveFileByKeyResponseBodySensitiveFileList) SetFilePath(v string) *DescribeAgentlessSensitiveFileByKeyResponseBodySensitiveFileList {
	s.FilePath = &v
	return s
}

func (s *DescribeAgentlessSensitiveFileByKeyResponseBodySensitiveFileList) SetFirstScanTime(v int64) *DescribeAgentlessSensitiveFileByKeyResponseBodySensitiveFileList {
	s.FirstScanTime = &v
	return s
}

func (s *DescribeAgentlessSensitiveFileByKeyResponseBodySensitiveFileList) SetId(v int64) *DescribeAgentlessSensitiveFileByKeyResponseBodySensitiveFileList {
	s.Id = &v
	return s
}

func (s *DescribeAgentlessSensitiveFileByKeyResponseBodySensitiveFileList) SetInstanceName(v string) *DescribeAgentlessSensitiveFileByKeyResponseBodySensitiveFileList {
	s.InstanceName = &v
	return s
}

func (s *DescribeAgentlessSensitiveFileByKeyResponseBodySensitiveFileList) SetInternetIp(v string) *DescribeAgentlessSensitiveFileByKeyResponseBodySensitiveFileList {
	s.InternetIp = &v
	return s
}

func (s *DescribeAgentlessSensitiveFileByKeyResponseBodySensitiveFileList) SetIntranetIp(v string) *DescribeAgentlessSensitiveFileByKeyResponseBodySensitiveFileList {
	s.IntranetIp = &v
	return s
}

func (s *DescribeAgentlessSensitiveFileByKeyResponseBodySensitiveFileList) SetLastScanTime(v int64) *DescribeAgentlessSensitiveFileByKeyResponseBodySensitiveFileList {
	s.LastScanTime = &v
	return s
}

func (s *DescribeAgentlessSensitiveFileByKeyResponseBodySensitiveFileList) SetMd5(v string) *DescribeAgentlessSensitiveFileByKeyResponseBodySensitiveFileList {
	s.Md5 = &v
	return s
}

func (s *DescribeAgentlessSensitiveFileByKeyResponseBodySensitiveFileList) SetPrompt(v string) *DescribeAgentlessSensitiveFileByKeyResponseBodySensitiveFileList {
	s.Prompt = &v
	return s
}

func (s *DescribeAgentlessSensitiveFileByKeyResponseBodySensitiveFileList) SetRiskLevel(v string) *DescribeAgentlessSensitiveFileByKeyResponseBodySensitiveFileList {
	s.RiskLevel = &v
	return s
}

func (s *DescribeAgentlessSensitiveFileByKeyResponseBodySensitiveFileList) SetSensitiveFileKey(v string) *DescribeAgentlessSensitiveFileByKeyResponseBodySensitiveFileList {
	s.SensitiveFileKey = &v
	return s
}

func (s *DescribeAgentlessSensitiveFileByKeyResponseBodySensitiveFileList) SetStatus(v int32) *DescribeAgentlessSensitiveFileByKeyResponseBodySensitiveFileList {
	s.Status = &v
	return s
}

func (s *DescribeAgentlessSensitiveFileByKeyResponseBodySensitiveFileList) SetTargetName(v string) *DescribeAgentlessSensitiveFileByKeyResponseBodySensitiveFileList {
	s.TargetName = &v
	return s
}

func (s *DescribeAgentlessSensitiveFileByKeyResponseBodySensitiveFileList) SetUuid(v string) *DescribeAgentlessSensitiveFileByKeyResponseBodySensitiveFileList {
	s.Uuid = &v
	return s
}

func (s *DescribeAgentlessSensitiveFileByKeyResponseBodySensitiveFileList) Validate() error {
	return dara.Validate(s)
}
