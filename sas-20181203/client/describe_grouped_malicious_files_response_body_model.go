// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupedMaliciousFilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroupedMaliciousFileResponse(v []*DescribeGroupedMaliciousFilesResponseBodyGroupedMaliciousFileResponse) *DescribeGroupedMaliciousFilesResponseBody
	GetGroupedMaliciousFileResponse() []*DescribeGroupedMaliciousFilesResponseBodyGroupedMaliciousFileResponse
	SetPageInfo(v *DescribeGroupedMaliciousFilesResponseBodyPageInfo) *DescribeGroupedMaliciousFilesResponseBody
	GetPageInfo() *DescribeGroupedMaliciousFilesResponseBodyPageInfo
	SetRequestId(v string) *DescribeGroupedMaliciousFilesResponseBody
	GetRequestId() *string
}

type DescribeGroupedMaliciousFilesResponseBody struct {
	// The details of the malicious image sample.
	GroupedMaliciousFileResponse []*DescribeGroupedMaliciousFilesResponseBodyGroupedMaliciousFileResponse `json:"GroupedMaliciousFileResponse,omitempty" xml:"GroupedMaliciousFileResponse,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *DescribeGroupedMaliciousFilesResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 8045E03E-6D91-4C53-9F22-5A1B84BB29D9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeGroupedMaliciousFilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupedMaliciousFilesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGroupedMaliciousFilesResponseBody) GetGroupedMaliciousFileResponse() []*DescribeGroupedMaliciousFilesResponseBodyGroupedMaliciousFileResponse {
	return s.GroupedMaliciousFileResponse
}

func (s *DescribeGroupedMaliciousFilesResponseBody) GetPageInfo() *DescribeGroupedMaliciousFilesResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeGroupedMaliciousFilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGroupedMaliciousFilesResponseBody) SetGroupedMaliciousFileResponse(v []*DescribeGroupedMaliciousFilesResponseBodyGroupedMaliciousFileResponse) *DescribeGroupedMaliciousFilesResponseBody {
	s.GroupedMaliciousFileResponse = v
	return s
}

func (s *DescribeGroupedMaliciousFilesResponseBody) SetPageInfo(v *DescribeGroupedMaliciousFilesResponseBodyPageInfo) *DescribeGroupedMaliciousFilesResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeGroupedMaliciousFilesResponseBody) SetRequestId(v string) *DescribeGroupedMaliciousFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGroupedMaliciousFilesResponseBody) Validate() error {
	if s.GroupedMaliciousFileResponse != nil {
		for _, item := range s.GroupedMaliciousFileResponse {
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

type DescribeGroupedMaliciousFilesResponseBodyGroupedMaliciousFileResponse struct {
	// The timestamp generated when the first scan was performed. Unit: milliseconds.
	//
	// example:
	//
	// 1594907349000
	FirstScanTimestamp *int64 `json:"FirstScanTimestamp,omitempty" xml:"FirstScanTimestamp,omitempty"`
	// The number of affected images.
	//
	// example:
	//
	// 3
	ImageCount *int64 `json:"ImageCount,omitempty" xml:"ImageCount,omitempty"`
	// The timestamp generated when the last scan was performed. Unit: milliseconds.
	//
	// example:
	//
	// 1596533942000
	LatestScanTimestamp *int64 `json:"LatestScanTimestamp,omitempty" xml:"LatestScanTimestamp,omitempty"`
	// The severity of the malicious image sample. Valid values:
	//
	// 	- **serious**
	//
	// 	- **suspicious**
	//
	// 	- **remind**
	//
	// example:
	//
	// serious
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The keyword of the malicious image sample.
	//
	// example:
	//
	// WEBSHELL_IMG
	MaliciousKey *string `json:"MaliciousKey,omitempty" xml:"MaliciousKey,omitempty"`
	// The MD5 hash value of the malicious image sample.
	//
	// example:
	//
	// d836968041f7683b5459****
	MaliciousMd5 *string `json:"MaliciousMd5,omitempty" xml:"MaliciousMd5,omitempty"`
	// The name of the malicious image sample.
	//
	// example:
	//
	// testFile
	MaliciousName *string `json:"MaliciousName,omitempty" xml:"MaliciousName,omitempty"`
	// The handling status of the malicious image sample. Valid values:
	//
	// 	- **0**: unhandled
	//
	// 	- **1**: handled
	//
	// 	- **2**: verifying
	//
	// 	- **3**: whitelisted
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeGroupedMaliciousFilesResponseBodyGroupedMaliciousFileResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupedMaliciousFilesResponseBodyGroupedMaliciousFileResponse) GoString() string {
	return s.String()
}

func (s *DescribeGroupedMaliciousFilesResponseBodyGroupedMaliciousFileResponse) GetFirstScanTimestamp() *int64 {
	return s.FirstScanTimestamp
}

func (s *DescribeGroupedMaliciousFilesResponseBodyGroupedMaliciousFileResponse) GetImageCount() *int64 {
	return s.ImageCount
}

func (s *DescribeGroupedMaliciousFilesResponseBodyGroupedMaliciousFileResponse) GetLatestScanTimestamp() *int64 {
	return s.LatestScanTimestamp
}

func (s *DescribeGroupedMaliciousFilesResponseBodyGroupedMaliciousFileResponse) GetLevel() *string {
	return s.Level
}

func (s *DescribeGroupedMaliciousFilesResponseBodyGroupedMaliciousFileResponse) GetMaliciousKey() *string {
	return s.MaliciousKey
}

func (s *DescribeGroupedMaliciousFilesResponseBodyGroupedMaliciousFileResponse) GetMaliciousMd5() *string {
	return s.MaliciousMd5
}

func (s *DescribeGroupedMaliciousFilesResponseBodyGroupedMaliciousFileResponse) GetMaliciousName() *string {
	return s.MaliciousName
}

func (s *DescribeGroupedMaliciousFilesResponseBodyGroupedMaliciousFileResponse) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeGroupedMaliciousFilesResponseBodyGroupedMaliciousFileResponse) SetFirstScanTimestamp(v int64) *DescribeGroupedMaliciousFilesResponseBodyGroupedMaliciousFileResponse {
	s.FirstScanTimestamp = &v
	return s
}

func (s *DescribeGroupedMaliciousFilesResponseBodyGroupedMaliciousFileResponse) SetImageCount(v int64) *DescribeGroupedMaliciousFilesResponseBodyGroupedMaliciousFileResponse {
	s.ImageCount = &v
	return s
}

func (s *DescribeGroupedMaliciousFilesResponseBodyGroupedMaliciousFileResponse) SetLatestScanTimestamp(v int64) *DescribeGroupedMaliciousFilesResponseBodyGroupedMaliciousFileResponse {
	s.LatestScanTimestamp = &v
	return s
}

func (s *DescribeGroupedMaliciousFilesResponseBodyGroupedMaliciousFileResponse) SetLevel(v string) *DescribeGroupedMaliciousFilesResponseBodyGroupedMaliciousFileResponse {
	s.Level = &v
	return s
}

func (s *DescribeGroupedMaliciousFilesResponseBodyGroupedMaliciousFileResponse) SetMaliciousKey(v string) *DescribeGroupedMaliciousFilesResponseBodyGroupedMaliciousFileResponse {
	s.MaliciousKey = &v
	return s
}

func (s *DescribeGroupedMaliciousFilesResponseBodyGroupedMaliciousFileResponse) SetMaliciousMd5(v string) *DescribeGroupedMaliciousFilesResponseBodyGroupedMaliciousFileResponse {
	s.MaliciousMd5 = &v
	return s
}

func (s *DescribeGroupedMaliciousFilesResponseBodyGroupedMaliciousFileResponse) SetMaliciousName(v string) *DescribeGroupedMaliciousFilesResponseBodyGroupedMaliciousFileResponse {
	s.MaliciousName = &v
	return s
}

func (s *DescribeGroupedMaliciousFilesResponseBodyGroupedMaliciousFileResponse) SetStatus(v int32) *DescribeGroupedMaliciousFilesResponseBodyGroupedMaliciousFileResponse {
	s.Status = &v
	return s
}

func (s *DescribeGroupedMaliciousFilesResponseBodyGroupedMaliciousFileResponse) Validate() error {
	return dara.Validate(s)
}

type DescribeGroupedMaliciousFilesResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeGroupedMaliciousFilesResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupedMaliciousFilesResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeGroupedMaliciousFilesResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribeGroupedMaliciousFilesResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeGroupedMaliciousFilesResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeGroupedMaliciousFilesResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeGroupedMaliciousFilesResponseBodyPageInfo) SetCount(v int32) *DescribeGroupedMaliciousFilesResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribeGroupedMaliciousFilesResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeGroupedMaliciousFilesResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeGroupedMaliciousFilesResponseBodyPageInfo) SetPageSize(v int32) *DescribeGroupedMaliciousFilesResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeGroupedMaliciousFilesResponseBodyPageInfo) SetTotalCount(v int32) *DescribeGroupedMaliciousFilesResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeGroupedMaliciousFilesResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
