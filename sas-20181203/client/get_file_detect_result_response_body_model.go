// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileDetectResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetFileDetectResultResponseBody
	GetRequestId() *string
	SetResultList(v []*GetFileDetectResultResponseBodyResultList) *GetFileDetectResultResponseBody
	GetResultList() []*GetFileDetectResultResponseBodyResultList
}

type GetFileDetectResultResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 24A20733-10A0-4AF6-BE6B-E3322413BB68
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array that consists of file detection results.
	ResultList []*GetFileDetectResultResponseBodyResultList `json:"ResultList,omitempty" xml:"ResultList,omitempty" type:"Repeated"`
}

func (s GetFileDetectResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFileDetectResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetFileDetectResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFileDetectResultResponseBody) GetResultList() []*GetFileDetectResultResponseBodyResultList {
	return s.ResultList
}

func (s *GetFileDetectResultResponseBody) SetRequestId(v string) *GetFileDetectResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFileDetectResultResponseBody) SetResultList(v []*GetFileDetectResultResponseBodyResultList) *GetFileDetectResultResponseBody {
	s.ResultList = v
	return s
}

func (s *GetFileDetectResultResponseBody) Validate() error {
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

type GetFileDetectResultResponseBodyResultList struct {
	// The status code returned. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Whether to identify as a compressed package. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// false
	Compress *bool `json:"Compress,omitempty" xml:"Compress,omitempty"`
	// The extended information about the file detection result.
	//
	// example:
	//
	// {"HighLight":[[23245,23212]]}
	Ext *string `json:"Ext,omitempty" xml:"Ext,omitempty"`
	// The identifier of the file. Only MD5 hash values are supported.
	//
	// example:
	//
	// 0a212417e65c26ff133cfff28f6c****
	HashKey *string `json:"HashKey,omitempty" xml:"HashKey,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
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
	// The score of file detection result.
	//
	// > A higher score indicates a more suspicious file.
	//
	// example:
	//
	// 100
	Score *int32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The type of the virus. Valid values:
	//
	// 	- **Trojan**: trojan horse
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
	// WEBSHELL
	VirusType *string `json:"VirusType,omitempty" xml:"VirusType,omitempty"`
}

func (s GetFileDetectResultResponseBodyResultList) String() string {
	return dara.Prettify(s)
}

func (s GetFileDetectResultResponseBodyResultList) GoString() string {
	return s.String()
}

func (s *GetFileDetectResultResponseBodyResultList) GetCode() *string {
	return s.Code
}

func (s *GetFileDetectResultResponseBodyResultList) GetCompress() *bool {
	return s.Compress
}

func (s *GetFileDetectResultResponseBodyResultList) GetExt() *string {
	return s.Ext
}

func (s *GetFileDetectResultResponseBodyResultList) GetHashKey() *string {
	return s.HashKey
}

func (s *GetFileDetectResultResponseBodyResultList) GetMessage() *string {
	return s.Message
}

func (s *GetFileDetectResultResponseBodyResultList) GetResult() *int32 {
	return s.Result
}

func (s *GetFileDetectResultResponseBodyResultList) GetScore() *int32 {
	return s.Score
}

func (s *GetFileDetectResultResponseBodyResultList) GetVirusType() *string {
	return s.VirusType
}

func (s *GetFileDetectResultResponseBodyResultList) SetCode(v string) *GetFileDetectResultResponseBodyResultList {
	s.Code = &v
	return s
}

func (s *GetFileDetectResultResponseBodyResultList) SetCompress(v bool) *GetFileDetectResultResponseBodyResultList {
	s.Compress = &v
	return s
}

func (s *GetFileDetectResultResponseBodyResultList) SetExt(v string) *GetFileDetectResultResponseBodyResultList {
	s.Ext = &v
	return s
}

func (s *GetFileDetectResultResponseBodyResultList) SetHashKey(v string) *GetFileDetectResultResponseBodyResultList {
	s.HashKey = &v
	return s
}

func (s *GetFileDetectResultResponseBodyResultList) SetMessage(v string) *GetFileDetectResultResponseBodyResultList {
	s.Message = &v
	return s
}

func (s *GetFileDetectResultResponseBodyResultList) SetResult(v int32) *GetFileDetectResultResponseBodyResultList {
	s.Result = &v
	return s
}

func (s *GetFileDetectResultResponseBodyResultList) SetScore(v int32) *GetFileDetectResultResponseBodyResultList {
	s.Score = &v
	return s
}

func (s *GetFileDetectResultResponseBodyResultList) SetVirusType(v string) *GetFileDetectResultResponseBodyResultList {
	s.VirusType = &v
	return s
}

func (s *GetFileDetectResultResponseBodyResultList) Validate() error {
	return dara.Validate(s)
}
