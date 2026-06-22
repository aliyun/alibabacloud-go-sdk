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
	// The page number of the page to return. Minimum value: **1**. Default value: **1**, which indicates that the first page is returned.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The end time in timestamp format.
	//
	// example:
	//
	// 1675058931215
	EndTimeStamp *int64 `json:"EndTimeStamp,omitempty" xml:"EndTimeStamp,omitempty"`
	// The language of the request and response. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The maximum number of entries to return on each page in a paged query. Default value: 20. If you leave this parameter empty, 20 entries are returned by default. > Specify a value for the paging parameter.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The list of risk levels.
	RiskLevelList []*string `json:"RiskLevelList,omitempty" xml:"RiskLevelList,omitempty" type:"Repeated"`
	// The IP address of the attack source.
	//
	// This parameter is required.
	//
	// example:
	//
	// 47.92.139.**
	SrcIp *string `json:"SrcIp,omitempty" xml:"SrcIp,omitempty"`
	// The start time in timestamp format.
	//
	// example:
	//
	// 1681624877761
	StartTimeStamp *int64 `json:"StartTimeStamp,omitempty" xml:"StartTimeStamp,omitempty"`
	// The type of attack source statistics information. Valid values:
	//
	// - **TOP_ATTACKED_AGENT**: the top 5 probes by number of attacks.
	//
	// - **TOP_ATTACKED_IP**: the top 5 IP addresses by number of attacks.
	//
	// - **ATTACK_EVENT_TYPE**: the intrusion event type.
	//
	// - **ATTACK_HONEYPOT_TYPE**: the type of the compromised honeypot.
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
