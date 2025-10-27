// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCompressFileDetectResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageInfo(v *ListCompressFileDetectResultResponseBodyPageInfo) *ListCompressFileDetectResultResponseBody
	GetPageInfo() *ListCompressFileDetectResultResponseBodyPageInfo
	SetRequestId(v string) *ListCompressFileDetectResultResponseBody
	GetRequestId() *string
	SetResultList(v []*ListCompressFileDetectResultResponseBodyResultList) *ListCompressFileDetectResultResponseBody
	GetResultList() []*ListCompressFileDetectResultResponseBodyResultList
}

type ListCompressFileDetectResultResponseBody struct {
	// The pagination information.
	PageInfo *ListCompressFileDetectResultResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// E10BAF1C-A6C5-51E2-866C-76D5922E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The detection results of files.
	ResultList []*ListCompressFileDetectResultResponseBodyResultList `json:"ResultList,omitempty" xml:"ResultList,omitempty" type:"Repeated"`
}

func (s ListCompressFileDetectResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCompressFileDetectResultResponseBody) GoString() string {
	return s.String()
}

func (s *ListCompressFileDetectResultResponseBody) GetPageInfo() *ListCompressFileDetectResultResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListCompressFileDetectResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCompressFileDetectResultResponseBody) GetResultList() []*ListCompressFileDetectResultResponseBodyResultList {
	return s.ResultList
}

func (s *ListCompressFileDetectResultResponseBody) SetPageInfo(v *ListCompressFileDetectResultResponseBodyPageInfo) *ListCompressFileDetectResultResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListCompressFileDetectResultResponseBody) SetRequestId(v string) *ListCompressFileDetectResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCompressFileDetectResultResponseBody) SetResultList(v []*ListCompressFileDetectResultResponseBodyResultList) *ListCompressFileDetectResultResponseBody {
	s.ResultList = v
	return s
}

func (s *ListCompressFileDetectResultResponseBody) Validate() error {
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	if s.ResultList != nil {
		for _, item := range s.ResultList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCompressFileDetectResultResponseBodyPageInfo struct {
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
	// 55
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCompressFileDetectResultResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListCompressFileDetectResultResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListCompressFileDetectResultResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListCompressFileDetectResultResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCompressFileDetectResultResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCompressFileDetectResultResponseBodyPageInfo) SetCurrentPage(v int32) *ListCompressFileDetectResultResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListCompressFileDetectResultResponseBodyPageInfo) SetPageSize(v int32) *ListCompressFileDetectResultResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListCompressFileDetectResultResponseBodyPageInfo) SetTotalCount(v int32) *ListCompressFileDetectResultResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListCompressFileDetectResultResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}

type ListCompressFileDetectResultResponseBodyResultList struct {
	// The extended information about the file detection result.
	//
	// example:
	//
	// {
	//
	//     "HighLight":
	//
	//     [
	//
	//         [
	//
	//             23245,
	//
	//             23212
	//
	//         ]
	//
	//     ],
	//
	//     "FileLabel":
	//
	//     [
	//
	//         "PE32",
	//
	//         "Zip",
	//
	//         "SFX",
	//
	//         "encrypted"
	//
	//     ]
	//
	// }
	Ext *string `json:"Ext,omitempty" xml:"Ext,omitempty"`
	// The identifier of the file.
	//
	// example:
	//
	// 0a212417e65c26ff133cfff28f6c****
	HashKey *string `json:"HashKey,omitempty" xml:"HashKey,omitempty"`
	// The path to the file within the package.
	//
	// example:
	//
	// /root/1.zip/test****
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The file detection result. Valid values:
	//
	// 	- **0**: The file is normal.
	//
	// 	- **1**: The file is suspicious.
	//
	// 	- **3**: The detection is in progress.
	//
	// example:
	//
	// 0
	Result *int32 `json:"Result,omitempty" xml:"Result,omitempty"`
	// The score of the file detection result. The following list describes mappings between the score ranges and risk levels:
	//
	// 	- 0 to 60: normal
	//
	// 	- 61 to 70: risky
	//
	// 	- 71 to 80: suspicious
	//
	// 	- 81 to 100: malicious
	//
	// >  A higher score indicates a more suspicious file.
	//
	// example:
	//
	// 100
	Score *int32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The type of the virus. Valid values:
	//
	// 	- **Trojan**: self-mutating trojan
	//
	// 	- **WebShell**: webshell
	//
	// 	- **Backdoor**: backdoor program
	//
	// 	- **RansomWare**: ransomware
	//
	// 	- **Scanner**: scanner
	//
	// 	- **Stealer**: tool that is used to steal information
	//
	// 	- **Malbaseware**: tainted basic software
	//
	// 	- **Hacktool**: attacker tool
	//
	// 	- **Engtest**: engine test program
	//
	// 	- **Downloader**: trojan downloader
	//
	// 	- **Virus**: infectious virus
	//
	// 	- **Miner**: mining program
	//
	// 	- **Worm**: worm
	//
	// 	- **DDoS**: DDoS trojan
	//
	// 	- **Malware**: malicious program
	//
	// 	- **Backdoor**: reverse shell
	//
	// 	- **RiskWare**: software that has risks
	//
	// 	- **Proxytool**: proxy
	//
	// 	- **Suspicious**: suspicious program
	//
	// 	- **MalScript**: malicious script
	//
	// 	- **Rootkit**: rootkit
	//
	// 	- **Exploit**: exploit
	//
	// example:
	//
	// WebShell
	VirusType *string `json:"VirusType,omitempty" xml:"VirusType,omitempty"`
}

func (s ListCompressFileDetectResultResponseBodyResultList) String() string {
	return dara.Prettify(s)
}

func (s ListCompressFileDetectResultResponseBodyResultList) GoString() string {
	return s.String()
}

func (s *ListCompressFileDetectResultResponseBodyResultList) GetExt() *string {
	return s.Ext
}

func (s *ListCompressFileDetectResultResponseBodyResultList) GetHashKey() *string {
	return s.HashKey
}

func (s *ListCompressFileDetectResultResponseBodyResultList) GetPath() *string {
	return s.Path
}

func (s *ListCompressFileDetectResultResponseBodyResultList) GetResult() *int32 {
	return s.Result
}

func (s *ListCompressFileDetectResultResponseBodyResultList) GetScore() *int32 {
	return s.Score
}

func (s *ListCompressFileDetectResultResponseBodyResultList) GetVirusType() *string {
	return s.VirusType
}

func (s *ListCompressFileDetectResultResponseBodyResultList) SetExt(v string) *ListCompressFileDetectResultResponseBodyResultList {
	s.Ext = &v
	return s
}

func (s *ListCompressFileDetectResultResponseBodyResultList) SetHashKey(v string) *ListCompressFileDetectResultResponseBodyResultList {
	s.HashKey = &v
	return s
}

func (s *ListCompressFileDetectResultResponseBodyResultList) SetPath(v string) *ListCompressFileDetectResultResponseBodyResultList {
	s.Path = &v
	return s
}

func (s *ListCompressFileDetectResultResponseBodyResultList) SetResult(v int32) *ListCompressFileDetectResultResponseBodyResultList {
	s.Result = &v
	return s
}

func (s *ListCompressFileDetectResultResponseBodyResultList) SetScore(v int32) *ListCompressFileDetectResultResponseBodyResultList {
	s.Score = &v
	return s
}

func (s *ListCompressFileDetectResultResponseBodyResultList) SetVirusType(v string) *ListCompressFileDetectResultResponseBodyResultList {
	s.VirusType = &v
	return s
}

func (s *ListCompressFileDetectResultResponseBodyResultList) Validate() error {
	return dara.Validate(s)
}
