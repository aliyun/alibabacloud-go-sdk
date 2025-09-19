// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAttackAnalysisDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBase64(v string) *DescribeAttackAnalysisDataRequest
	GetBase64() *string
	SetCurrentPage(v int32) *DescribeAttackAnalysisDataRequest
	GetCurrentPage() *int32
	SetData(v string) *DescribeAttackAnalysisDataRequest
	GetData() *string
	SetEndTime(v int64) *DescribeAttackAnalysisDataRequest
	GetEndTime() *int64
	SetLang(v string) *DescribeAttackAnalysisDataRequest
	GetLang() *string
	SetPageSize(v int32) *DescribeAttackAnalysisDataRequest
	GetPageSize() *int32
	SetStartTime(v int64) *DescribeAttackAnalysisDataRequest
	GetStartTime() *int64
	SetType(v string) *DescribeAttackAnalysisDataRequest
	GetType() *string
}

type DescribeAttackAnalysisDataRequest struct {
	// Specifies whether to encode the value of the **client_url*	- field in the query results by using the Base64 algorithm. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	Base64 *string `json:"Base64,omitempty" xml:"Base64,omitempty"`
	// The number of the page to return. Pages start from page **1**.
	//
	// >  If the Type parameter is set to **DETAILS**, you must specify the CurrentPage parameter.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The condition that is used to filter attack events.
	//
	// >  The following list describes the valid values of crack_type:
	//
	// 	- 3: brute-force attack on MySQL
	//
	// 	- 4: FTP brute-force attack
	//
	// 	- 5: SSH brute-force attack
	//
	// 	- 6: RDP brute-force attack
	//
	// 	- 9: brute-force attack on Microsoft SQL Server
	//
	// 	- 101: intercepted attack on Java Struts 2
	//
	// 	- 102: intercepted attack on Redis
	//
	// 	- 103: communication with AntSword Webshell
	//
	// 	- 104: communication with China Chopper Webshell
	//
	// 	- 133: communication with XISE Webshell
	//
	// 	- sqli: SQL injection
	//
	// 	- codei: code execution
	//
	// 	- xss: cross-site scripting (XSS)
	//
	// 	- lfi: local file inclusion
	//
	// 	- rfi: remote file inclusion
	//
	// 	- webshell: trojan script
	//
	// 	- upload: vulnerability upload
	//
	// 	- path: directory traversal
	//
	// 	- bypass: unauthorized access
	//
	// 	- csrf: cross-site request forgery (CSRF)
	//
	// 	- crlf: carriage return line feed (CRLF)
	//
	// 	- other: others
	//
	// example:
	//
	// {"crack_type":"9"}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The timestamp when the attack stops. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1649040221
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries to return on each page.
	//
	// >  If the Type parameter is set to **DETAILS**, you must specify the PageSize parameter.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The timestamp at which the attack starts. By default, the statistics of the previous seven days are queried. Unit: seconds.
	//
	// >  The start time that you specify must be within the previous 40 days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1644027670
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The details of attack analysis. Valid values:
	//
	// 	- **TOTAL**: number of attacks
	//
	// 	- **TREND**: attack trend
	//
	// 	- **PIE_CHART**: distribution of attacks by type
	//
	// 	- **SOURCE_TOP**: top 5 attack sources
	//
	// 	- **CLIENT_TOP**: top 5 attacked assets
	//
	// 	- **DETAILS**: attack details
	//
	// >  If the Type parameter is set to **DETAILS**, you must specify the CurrentPage and PageSize parameters.
	//
	// This parameter is required.
	//
	// example:
	//
	// DETAILS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeAttackAnalysisDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAttackAnalysisDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeAttackAnalysisDataRequest) GetBase64() *string {
	return s.Base64
}

func (s *DescribeAttackAnalysisDataRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeAttackAnalysisDataRequest) GetData() *string {
	return s.Data
}

func (s *DescribeAttackAnalysisDataRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeAttackAnalysisDataRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeAttackAnalysisDataRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAttackAnalysisDataRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeAttackAnalysisDataRequest) GetType() *string {
	return s.Type
}

func (s *DescribeAttackAnalysisDataRequest) SetBase64(v string) *DescribeAttackAnalysisDataRequest {
	s.Base64 = &v
	return s
}

func (s *DescribeAttackAnalysisDataRequest) SetCurrentPage(v int32) *DescribeAttackAnalysisDataRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAttackAnalysisDataRequest) SetData(v string) *DescribeAttackAnalysisDataRequest {
	s.Data = &v
	return s
}

func (s *DescribeAttackAnalysisDataRequest) SetEndTime(v int64) *DescribeAttackAnalysisDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAttackAnalysisDataRequest) SetLang(v string) *DescribeAttackAnalysisDataRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAttackAnalysisDataRequest) SetPageSize(v int32) *DescribeAttackAnalysisDataRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAttackAnalysisDataRequest) SetStartTime(v int64) *DescribeAttackAnalysisDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeAttackAnalysisDataRequest) SetType(v string) *DescribeAttackAnalysisDataRequest {
	s.Type = &v
	return s
}

func (s *DescribeAttackAnalysisDataRequest) Validate() error {
	return dara.Validate(s)
}
