// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAttackEventInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttackInstance(v string) *ListAttackEventInfoRequest
	GetAttackInstance() *string
	SetAttackType(v string) *ListAttackEventInfoRequest
	GetAttackType() *string
	SetCurrentPage(v int32) *ListAttackEventInfoRequest
	GetCurrentPage() *int32
	SetDstPort(v string) *ListAttackEventInfoRequest
	GetDstPort() *string
	SetEndTime(v int64) *ListAttackEventInfoRequest
	GetEndTime() *int64
	SetLang(v string) *ListAttackEventInfoRequest
	GetLang() *string
	SetPageSize(v int32) *ListAttackEventInfoRequest
	GetPageSize() *int32
	SetSrcIp(v string) *ListAttackEventInfoRequest
	GetSrcIp() *string
	SetStartTime(v int64) *ListAttackEventInfoRequest
	GetStartTime() *int64
}

type ListAttackEventInfoRequest struct {
	// Attacked asset. You can fill in the machine instance name, public IP, or private IP.
	//
	// example:
	//
	// instance_**
	AttackInstance *string `json:"AttackInstance,omitempty" xml:"AttackInstance,omitempty"`
	// Attack type. Values:
	//
	// - 9: SQL Server brute force attack
	//
	// - 5: SSH brute force attack
	//
	// - 6: RDP brute force attack
	//
	// - 101: Java Struts2 attack interception
	//
	// - 102: Redis attack interception
	//
	// - 103: Chinese AntSword WebShell communication
	//
	// - 104: Chinese Chopper WebShell communication
	//
	// - 133: XISE WebShell communication
	//
	// - 161: WebShell upload
	//
	// - 209: PHP WebShell upload
	//
	// - 210: JSP WebShell upload
	//
	// - 211: ASP WebShell upload
	//
	// - 215: Special suffix WebShell upload
	//
	// - ai_webshell: Intelligent defense for WebShell upload
	//
	// - java_common_rce: Java common RCE vulnerability interception
	//
	// - alinet_webrce: Adaptive web attack defense
	//
	// example:
	//
	// 9
	AttackType *string `json:"AttackType,omitempty" xml:"AttackType,omitempty"`
	// Page number to display in a paginated query.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Target port of the attack.
	//
	// example:
	//
	// 9085
	DstPort *string `json:"DstPort,omitempty" xml:"DstPort,omitempty"`
	// Timestamp of the end time.
	//
	// example:
	//
	// 1753152532550
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Set the language type for request and response messages. Default is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Maximum number of records per page in a paginated query.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Source IP of the attack.
	//
	// example:
	//
	// 185.237.96.***
	SrcIp *string `json:"SrcIp,omitempty" xml:"SrcIp,omitempty"`
	// Timestamp of the start time.
	//
	// example:
	//
	// 1752547732549
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListAttackEventInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAttackEventInfoRequest) GoString() string {
	return s.String()
}

func (s *ListAttackEventInfoRequest) GetAttackInstance() *string {
	return s.AttackInstance
}

func (s *ListAttackEventInfoRequest) GetAttackType() *string {
	return s.AttackType
}

func (s *ListAttackEventInfoRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListAttackEventInfoRequest) GetDstPort() *string {
	return s.DstPort
}

func (s *ListAttackEventInfoRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListAttackEventInfoRequest) GetLang() *string {
	return s.Lang
}

func (s *ListAttackEventInfoRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAttackEventInfoRequest) GetSrcIp() *string {
	return s.SrcIp
}

func (s *ListAttackEventInfoRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListAttackEventInfoRequest) SetAttackInstance(v string) *ListAttackEventInfoRequest {
	s.AttackInstance = &v
	return s
}

func (s *ListAttackEventInfoRequest) SetAttackType(v string) *ListAttackEventInfoRequest {
	s.AttackType = &v
	return s
}

func (s *ListAttackEventInfoRequest) SetCurrentPage(v int32) *ListAttackEventInfoRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListAttackEventInfoRequest) SetDstPort(v string) *ListAttackEventInfoRequest {
	s.DstPort = &v
	return s
}

func (s *ListAttackEventInfoRequest) SetEndTime(v int64) *ListAttackEventInfoRequest {
	s.EndTime = &v
	return s
}

func (s *ListAttackEventInfoRequest) SetLang(v string) *ListAttackEventInfoRequest {
	s.Lang = &v
	return s
}

func (s *ListAttackEventInfoRequest) SetPageSize(v int32) *ListAttackEventInfoRequest {
	s.PageSize = &v
	return s
}

func (s *ListAttackEventInfoRequest) SetSrcIp(v string) *ListAttackEventInfoRequest {
	s.SrcIp = &v
	return s
}

func (s *ListAttackEventInfoRequest) SetStartTime(v int64) *ListAttackEventInfoRequest {
	s.StartTime = &v
	return s
}

func (s *ListAttackEventInfoRequest) Validate() error {
	return dara.Validate(s)
}
