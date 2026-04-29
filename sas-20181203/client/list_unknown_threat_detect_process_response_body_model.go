// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUnknownThreatDetectProcessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListUnknownThreatDetectProcessResponseBodyData) *ListUnknownThreatDetectProcessResponseBody
	GetData() []*ListUnknownThreatDetectProcessResponseBodyData
	SetPageInfo(v *ListUnknownThreatDetectProcessResponseBodyPageInfo) *ListUnknownThreatDetectProcessResponseBody
	GetPageInfo() *ListUnknownThreatDetectProcessResponseBodyPageInfo
	SetRequestId(v string) *ListUnknownThreatDetectProcessResponseBody
	GetRequestId() *string
}

type ListUnknownThreatDetectProcessResponseBody struct {
	Data     []*ListUnknownThreatDetectProcessResponseBodyData   `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	PageInfo *ListUnknownThreatDetectProcessResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 20456DD5-5CBF-5015-9173-12CA4246B***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListUnknownThreatDetectProcessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUnknownThreatDetectProcessResponseBody) GoString() string {
	return s.String()
}

func (s *ListUnknownThreatDetectProcessResponseBody) GetData() []*ListUnknownThreatDetectProcessResponseBodyData {
	return s.Data
}

func (s *ListUnknownThreatDetectProcessResponseBody) GetPageInfo() *ListUnknownThreatDetectProcessResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListUnknownThreatDetectProcessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUnknownThreatDetectProcessResponseBody) SetData(v []*ListUnknownThreatDetectProcessResponseBodyData) *ListUnknownThreatDetectProcessResponseBody {
	s.Data = v
	return s
}

func (s *ListUnknownThreatDetectProcessResponseBody) SetPageInfo(v *ListUnknownThreatDetectProcessResponseBodyPageInfo) *ListUnknownThreatDetectProcessResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListUnknownThreatDetectProcessResponseBody) SetRequestId(v string) *ListUnknownThreatDetectProcessResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUnknownThreatDetectProcessResponseBody) Validate() error {
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

type ListUnknownThreatDetectProcessResponseBodyData struct {
	// example:
	//
	// white
	AnalyzeResult *string `json:"AnalyzeResult,omitempty" xml:"AnalyzeResult,omitempty"`
	// example:
	//
	// 1694576692000
	FirstTime *int64 `json:"FirstTime,omitempty" xml:"FirstTime,omitempty"`
	// example:
	//
	// 5b394b54ca632fe51c4ab4a6dbaf****
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// example:
	//
	// 2025031506350619216822625103151158982
	ProcessId *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	// example:
	//
	// /usr/bin/tar
	ProcessPath *string `json:"ProcessPath,omitempty" xml:"ProcessPath,omitempty"`
	// example:
	//
	// safe process
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// 3a6fed5fc11392b3ee9f81caf017b48640d7458766a8eb0382899a605b41****
	Sha256 *string `json:"Sha256,omitempty" xml:"Sha256,omitempty"`
}

func (s ListUnknownThreatDetectProcessResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListUnknownThreatDetectProcessResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListUnknownThreatDetectProcessResponseBodyData) GetAnalyzeResult() *string {
	return s.AnalyzeResult
}

func (s *ListUnknownThreatDetectProcessResponseBodyData) GetFirstTime() *int64 {
	return s.FirstTime
}

func (s *ListUnknownThreatDetectProcessResponseBodyData) GetMd5() *string {
	return s.Md5
}

func (s *ListUnknownThreatDetectProcessResponseBodyData) GetProcessId() *string {
	return s.ProcessId
}

func (s *ListUnknownThreatDetectProcessResponseBodyData) GetProcessPath() *string {
	return s.ProcessPath
}

func (s *ListUnknownThreatDetectProcessResponseBodyData) GetRemark() *string {
	return s.Remark
}

func (s *ListUnknownThreatDetectProcessResponseBodyData) GetSha256() *string {
	return s.Sha256
}

func (s *ListUnknownThreatDetectProcessResponseBodyData) SetAnalyzeResult(v string) *ListUnknownThreatDetectProcessResponseBodyData {
	s.AnalyzeResult = &v
	return s
}

func (s *ListUnknownThreatDetectProcessResponseBodyData) SetFirstTime(v int64) *ListUnknownThreatDetectProcessResponseBodyData {
	s.FirstTime = &v
	return s
}

func (s *ListUnknownThreatDetectProcessResponseBodyData) SetMd5(v string) *ListUnknownThreatDetectProcessResponseBodyData {
	s.Md5 = &v
	return s
}

func (s *ListUnknownThreatDetectProcessResponseBodyData) SetProcessId(v string) *ListUnknownThreatDetectProcessResponseBodyData {
	s.ProcessId = &v
	return s
}

func (s *ListUnknownThreatDetectProcessResponseBodyData) SetProcessPath(v string) *ListUnknownThreatDetectProcessResponseBodyData {
	s.ProcessPath = &v
	return s
}

func (s *ListUnknownThreatDetectProcessResponseBodyData) SetRemark(v string) *ListUnknownThreatDetectProcessResponseBodyData {
	s.Remark = &v
	return s
}

func (s *ListUnknownThreatDetectProcessResponseBodyData) SetSha256(v string) *ListUnknownThreatDetectProcessResponseBodyData {
	s.Sha256 = &v
	return s
}

func (s *ListUnknownThreatDetectProcessResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListUnknownThreatDetectProcessResponseBodyPageInfo struct {
	// example:
	//
	// 2
	Count *string `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 83
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListUnknownThreatDetectProcessResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListUnknownThreatDetectProcessResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListUnknownThreatDetectProcessResponseBodyPageInfo) GetCount() *string {
	return s.Count
}

func (s *ListUnknownThreatDetectProcessResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListUnknownThreatDetectProcessResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUnknownThreatDetectProcessResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListUnknownThreatDetectProcessResponseBodyPageInfo) SetCount(v string) *ListUnknownThreatDetectProcessResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *ListUnknownThreatDetectProcessResponseBodyPageInfo) SetCurrentPage(v int32) *ListUnknownThreatDetectProcessResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListUnknownThreatDetectProcessResponseBodyPageInfo) SetPageSize(v int32) *ListUnknownThreatDetectProcessResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListUnknownThreatDetectProcessResponseBodyPageInfo) SetTotalCount(v int32) *ListUnknownThreatDetectProcessResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListUnknownThreatDetectProcessResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
