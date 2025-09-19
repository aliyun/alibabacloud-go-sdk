// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHoneypotEventTrendRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTimeStamp(v int64) *GetHoneypotEventTrendRequest
	GetEndTimeStamp() *int64
	SetLang(v string) *GetHoneypotEventTrendRequest
	GetLang() *string
	SetRiskLevelList(v []*string) *GetHoneypotEventTrendRequest
	GetRiskLevelList() []*string
	SetSrcIp(v string) *GetHoneypotEventTrendRequest
	GetSrcIp() *string
	SetStartTimeStamp(v int64) *GetHoneypotEventTrendRequest
	GetStartTimeStamp() *int64
}

type GetHoneypotEventTrendRequest struct {
	// End time, timestamp format.
	//
	// example:
	//
	// 1687831329169
	EndTimeStamp *int64 `json:"EndTimeStamp,omitempty" xml:"EndTimeStamp,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The risk levels of the alert events.
	RiskLevelList []*string `json:"RiskLevelList,omitempty" xml:"RiskLevelList,omitempty" type:"Repeated"`
	// The source IP address of the attack.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.91.254.***
	SrcIp *string `json:"SrcIp,omitempty" xml:"SrcIp,omitempty"`
	// Start time, timestamp format.
	//
	// example:
	//
	// 1683516557757
	StartTimeStamp *int64 `json:"StartTimeStamp,omitempty" xml:"StartTimeStamp,omitempty"`
}

func (s GetHoneypotEventTrendRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHoneypotEventTrendRequest) GoString() string {
	return s.String()
}

func (s *GetHoneypotEventTrendRequest) GetEndTimeStamp() *int64 {
	return s.EndTimeStamp
}

func (s *GetHoneypotEventTrendRequest) GetLang() *string {
	return s.Lang
}

func (s *GetHoneypotEventTrendRequest) GetRiskLevelList() []*string {
	return s.RiskLevelList
}

func (s *GetHoneypotEventTrendRequest) GetSrcIp() *string {
	return s.SrcIp
}

func (s *GetHoneypotEventTrendRequest) GetStartTimeStamp() *int64 {
	return s.StartTimeStamp
}

func (s *GetHoneypotEventTrendRequest) SetEndTimeStamp(v int64) *GetHoneypotEventTrendRequest {
	s.EndTimeStamp = &v
	return s
}

func (s *GetHoneypotEventTrendRequest) SetLang(v string) *GetHoneypotEventTrendRequest {
	s.Lang = &v
	return s
}

func (s *GetHoneypotEventTrendRequest) SetRiskLevelList(v []*string) *GetHoneypotEventTrendRequest {
	s.RiskLevelList = v
	return s
}

func (s *GetHoneypotEventTrendRequest) SetSrcIp(v string) *GetHoneypotEventTrendRequest {
	s.SrcIp = &v
	return s
}

func (s *GetHoneypotEventTrendRequest) SetStartTimeStamp(v int64) *GetHoneypotEventTrendRequest {
	s.StartTimeStamp = &v
	return s
}

func (s *GetHoneypotEventTrendRequest) Validate() error {
	return dara.Validate(s)
}
