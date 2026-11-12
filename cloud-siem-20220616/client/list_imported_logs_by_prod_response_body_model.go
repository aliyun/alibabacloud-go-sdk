// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImportedLogsByProdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListImportedLogsByProdResponseBodyData) *ListImportedLogsByProdResponseBody
	GetData() []*ListImportedLogsByProdResponseBodyData
	SetRequestId(v string) *ListImportedLogsByProdResponseBody
	GetRequestId() *string
}

type ListImportedLogsByProdResponseBody struct {
	// The response data.
	Data []*ListImportedLogsByProdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListImportedLogsByProdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListImportedLogsByProdResponseBody) GoString() string {
	return s.String()
}

func (s *ListImportedLogsByProdResponseBody) GetData() []*ListImportedLogsByProdResponseBodyData {
	return s.Data
}

func (s *ListImportedLogsByProdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListImportedLogsByProdResponseBody) SetData(v []*ListImportedLogsByProdResponseBodyData) *ListImportedLogsByProdResponseBody {
	s.Data = v
	return s
}

func (s *ListImportedLogsByProdResponseBody) SetRequestId(v string) *ListImportedLogsByProdResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListImportedLogsByProdResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListImportedLogsByProdResponseBodyData struct {
	// Indicates whether new accounts are automatically onboarded. Valid values:
	//
	// - 1: Automatically onboarded.
	//
	// - 0: Not automatically onboarded.
	//
	// example:
	//
	// 1
	AutoImported *int32 `json:"AutoImported,omitempty" xml:"AutoImported,omitempty"`
	// The multi-cloud code. Valid values:
	//
	// - hcloud: Huawei Cloud.
	//
	// - qcloud: Tencent Cloud.
	//
	// - aliyun: Alibaba Cloud.
	//
	// example:
	//
	// hcloud
	CloudCode *string `json:"CloudCode,omitempty" xml:"CloudCode,omitempty"`
	// Indicates whether the log is onboarded. Valid values:
	//
	// - 1: Onboarded.
	//
	// - 0: Not onboarded.
	//
	// example:
	//
	// 1
	Imported *int32 `json:"Imported,omitempty" xml:"Imported,omitempty"`
	// The number of users that have onboarded the log.
	//
	// example:
	//
	// 2
	ImportedUserCount *int32 `json:"ImportedUserCount,omitempty" xml:"ImportedUserCount,omitempty"`
	// The log code.
	//
	// example:
	//
	// cloud_siem_waf_xxxxx
	LogCode *string `json:"LogCode,omitempty" xml:"LogCode,omitempty"`
	// The log display code.
	//
	// example:
	//
	// ${siem.prod. cloud_siem_waf_xxxxx}
	LogMdsCode *string `json:"LogMdsCode,omitempty" xml:"LogMdsCode,omitempty"`
	// The log type. Valid values:
	//
	// - 1: Centralized onboarding.
	//
	// - 2: Predefined Log Service.
	//
	// - 3: Custom Log Service.
	//
	// example:
	//
	// 1
	LogType *int32 `json:"LogType,omitempty" xml:"LogType,omitempty"`
	// The last onboarding time of the log. The time is in the YYYY-MM-DD HH:mm:ss format.
	//
	// example:
	//
	// 2023-11-23 12:30:00
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The product code corresponding to the log.
	//
	// example:
	//
	// qcloud_waf
	ProdCode *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	// The total number of users for the log.
	//
	// example:
	//
	// 5
	TotalUserCount *int32 `json:"TotalUserCount,omitempty" xml:"TotalUserCount,omitempty"`
	// The number of users that have not onboarded the log.
	//
	// example:
	//
	// 3
	UnImportedUserCount *int32 `json:"UnImportedUserCount,omitempty" xml:"UnImportedUserCount,omitempty"`
}

func (s ListImportedLogsByProdResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListImportedLogsByProdResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListImportedLogsByProdResponseBodyData) GetAutoImported() *int32 {
	return s.AutoImported
}

func (s *ListImportedLogsByProdResponseBodyData) GetCloudCode() *string {
	return s.CloudCode
}

func (s *ListImportedLogsByProdResponseBodyData) GetImported() *int32 {
	return s.Imported
}

func (s *ListImportedLogsByProdResponseBodyData) GetImportedUserCount() *int32 {
	return s.ImportedUserCount
}

func (s *ListImportedLogsByProdResponseBodyData) GetLogCode() *string {
	return s.LogCode
}

func (s *ListImportedLogsByProdResponseBodyData) GetLogMdsCode() *string {
	return s.LogMdsCode
}

func (s *ListImportedLogsByProdResponseBodyData) GetLogType() *int32 {
	return s.LogType
}

func (s *ListImportedLogsByProdResponseBodyData) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ListImportedLogsByProdResponseBodyData) GetProdCode() *string {
	return s.ProdCode
}

func (s *ListImportedLogsByProdResponseBodyData) GetTotalUserCount() *int32 {
	return s.TotalUserCount
}

func (s *ListImportedLogsByProdResponseBodyData) GetUnImportedUserCount() *int32 {
	return s.UnImportedUserCount
}

func (s *ListImportedLogsByProdResponseBodyData) SetAutoImported(v int32) *ListImportedLogsByProdResponseBodyData {
	s.AutoImported = &v
	return s
}

func (s *ListImportedLogsByProdResponseBodyData) SetCloudCode(v string) *ListImportedLogsByProdResponseBodyData {
	s.CloudCode = &v
	return s
}

func (s *ListImportedLogsByProdResponseBodyData) SetImported(v int32) *ListImportedLogsByProdResponseBodyData {
	s.Imported = &v
	return s
}

func (s *ListImportedLogsByProdResponseBodyData) SetImportedUserCount(v int32) *ListImportedLogsByProdResponseBodyData {
	s.ImportedUserCount = &v
	return s
}

func (s *ListImportedLogsByProdResponseBodyData) SetLogCode(v string) *ListImportedLogsByProdResponseBodyData {
	s.LogCode = &v
	return s
}

func (s *ListImportedLogsByProdResponseBodyData) SetLogMdsCode(v string) *ListImportedLogsByProdResponseBodyData {
	s.LogMdsCode = &v
	return s
}

func (s *ListImportedLogsByProdResponseBodyData) SetLogType(v int32) *ListImportedLogsByProdResponseBodyData {
	s.LogType = &v
	return s
}

func (s *ListImportedLogsByProdResponseBodyData) SetModifyTime(v string) *ListImportedLogsByProdResponseBodyData {
	s.ModifyTime = &v
	return s
}

func (s *ListImportedLogsByProdResponseBodyData) SetProdCode(v string) *ListImportedLogsByProdResponseBodyData {
	s.ProdCode = &v
	return s
}

func (s *ListImportedLogsByProdResponseBodyData) SetTotalUserCount(v int32) *ListImportedLogsByProdResponseBodyData {
	s.TotalUserCount = &v
	return s
}

func (s *ListImportedLogsByProdResponseBodyData) SetUnImportedUserCount(v int32) *ListImportedLogsByProdResponseBodyData {
	s.UnImportedUserCount = &v
	return s
}

func (s *ListImportedLogsByProdResponseBodyData) Validate() error {
	return dara.Validate(s)
}
