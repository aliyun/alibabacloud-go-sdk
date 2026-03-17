// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHoneypotAttackStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *GetHoneypotAttackStatisticsRequest
	GetCurrentPage() *int32
	SetEndTimeStamp(v int64) *GetHoneypotAttackStatisticsRequest
	GetEndTimeStamp() *int64
	SetLang(v string) *GetHoneypotAttackStatisticsRequest
	GetLang() *string
	SetPageSize(v int32) *GetHoneypotAttackStatisticsRequest
	GetPageSize() *int32
	SetRiskLevelList(v []*string) *GetHoneypotAttackStatisticsRequest
	GetRiskLevelList() []*string
	SetSrcIp(v string) *GetHoneypotAttackStatisticsRequest
	GetSrcIp() *string
	SetStartTimeStamp(v int64) *GetHoneypotAttackStatisticsRequest
	GetStartTimeStamp() *int64
	SetStatisticsType(v string) *GetHoneypotAttackStatisticsRequest
	GetStatisticsType() *string
}

type GetHoneypotAttackStatisticsRequest struct {
	// Set the page number from which to start displaying the query results. The starting value is **1**. The default value is **1**, indicating that the display starts from the **1st*	- page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// End time, in timestamp format.
	//
	// example:
	//
	// 1675058931215
	EndTimeStamp *int64 `json:"EndTimeStamp,omitempty" xml:"EndTimeStamp,omitempty"`
	// Sets the language type for requests and received messages, default is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Specifies the maximum number of data entries displayed per page for paginated queries. The default number of entries displayed per page is 20. If the pagesize parameter is empty, 20 entries will be returned by default. It is recommended that the pagesize value should not be empty.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// List of risk levels
	RiskLevelList []*string `json:"RiskLevelList,omitempty" xml:"RiskLevelList,omitempty" type:"Repeated"`
	// Attacker\\"s IP
	//
	// This parameter is required.
	//
	// example:
	//
	// 47.92.139.**
	SrcIp *string `json:"SrcIp,omitempty" xml:"SrcIp,omitempty"`
	// Start time, in timestamp format.
	//
	// example:
	//
	// 1681624877761
	StartTimeStamp *int64 `json:"StartTimeStamp,omitempty" xml:"StartTimeStamp,omitempty"`
	// The type of attack source statistics. Values:
	//
	// - **TOP_ATTACKED_AGENT**: Top 5 most attacked probes.
	//
	// - **TOP_ATTACKED_IP**: Top 5 most attacked IP addresses.
	//
	//  - **ATTACK_EVENT_TYPE**: Type of intrusion event.
	//
	// - **ATTACK_HONEYPOT_TYPE**: Type of compromised honeypot.
	//
	// This parameter is required.
	//
	// example:
	//
	// TOP_ATTACKED_IP
	StatisticsType *string `json:"StatisticsType,omitempty" xml:"StatisticsType,omitempty"`
}

func (s GetHoneypotAttackStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHoneypotAttackStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetHoneypotAttackStatisticsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetHoneypotAttackStatisticsRequest) GetEndTimeStamp() *int64 {
	return s.EndTimeStamp
}

func (s *GetHoneypotAttackStatisticsRequest) GetLang() *string {
	return s.Lang
}

func (s *GetHoneypotAttackStatisticsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetHoneypotAttackStatisticsRequest) GetRiskLevelList() []*string {
	return s.RiskLevelList
}

func (s *GetHoneypotAttackStatisticsRequest) GetSrcIp() *string {
	return s.SrcIp
}

func (s *GetHoneypotAttackStatisticsRequest) GetStartTimeStamp() *int64 {
	return s.StartTimeStamp
}

func (s *GetHoneypotAttackStatisticsRequest) GetStatisticsType() *string {
	return s.StatisticsType
}

func (s *GetHoneypotAttackStatisticsRequest) SetCurrentPage(v int32) *GetHoneypotAttackStatisticsRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetHoneypotAttackStatisticsRequest) SetEndTimeStamp(v int64) *GetHoneypotAttackStatisticsRequest {
	s.EndTimeStamp = &v
	return s
}

func (s *GetHoneypotAttackStatisticsRequest) SetLang(v string) *GetHoneypotAttackStatisticsRequest {
	s.Lang = &v
	return s
}

func (s *GetHoneypotAttackStatisticsRequest) SetPageSize(v int32) *GetHoneypotAttackStatisticsRequest {
	s.PageSize = &v
	return s
}

func (s *GetHoneypotAttackStatisticsRequest) SetRiskLevelList(v []*string) *GetHoneypotAttackStatisticsRequest {
	s.RiskLevelList = v
	return s
}

func (s *GetHoneypotAttackStatisticsRequest) SetSrcIp(v string) *GetHoneypotAttackStatisticsRequest {
	s.SrcIp = &v
	return s
}

func (s *GetHoneypotAttackStatisticsRequest) SetStartTimeStamp(v int64) *GetHoneypotAttackStatisticsRequest {
	s.StartTimeStamp = &v
	return s
}

func (s *GetHoneypotAttackStatisticsRequest) SetStatisticsType(v string) *GetHoneypotAttackStatisticsRequest {
	s.StatisticsType = &v
	return s
}

func (s *GetHoneypotAttackStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
