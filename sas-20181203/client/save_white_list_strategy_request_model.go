// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveWhiteListStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *SaveWhiteListStrategyRequest
	GetLang() *string
	SetSourceIp(v string) *SaveWhiteListStrategyRequest
	GetSourceIp() *string
	SetStrategyId(v int64) *SaveWhiteListStrategyRequest
	GetStrategyId() *int64
	SetStrategyName(v string) *SaveWhiteListStrategyRequest
	GetStrategyName() *string
	SetStudyTime(v int32) *SaveWhiteListStrategyRequest
	GetStudyTime() *int32
}

type SaveWhiteListStrategyRequest struct {
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
	// The source IP address of the request. You do not need to specify this parameter. It is automatically obtained by the system.
	//
	// example:
	//
	// 124.89.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The ID of the application whitelist policy.
	//
	// >  You can call the [DescribeWhiteListStrategyList](~~DescribeWhiteListStrategyList~~) operation to query the ID.
	//
	// example:
	//
	// 8494
	StrategyId *int64 `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	// The name of the application whitelist policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	StrategyName *string `json:"StrategyName,omitempty" xml:"StrategyName,omitempty"`
	// The duration of intelligent learning. Unit: hours.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	StudyTime *int32 `json:"StudyTime,omitempty" xml:"StudyTime,omitempty"`
}

func (s SaveWhiteListStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveWhiteListStrategyRequest) GoString() string {
	return s.String()
}

func (s *SaveWhiteListStrategyRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveWhiteListStrategyRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *SaveWhiteListStrategyRequest) GetStrategyId() *int64 {
	return s.StrategyId
}

func (s *SaveWhiteListStrategyRequest) GetStrategyName() *string {
	return s.StrategyName
}

func (s *SaveWhiteListStrategyRequest) GetStudyTime() *int32 {
	return s.StudyTime
}

func (s *SaveWhiteListStrategyRequest) SetLang(v string) *SaveWhiteListStrategyRequest {
	s.Lang = &v
	return s
}

func (s *SaveWhiteListStrategyRequest) SetSourceIp(v string) *SaveWhiteListStrategyRequest {
	s.SourceIp = &v
	return s
}

func (s *SaveWhiteListStrategyRequest) SetStrategyId(v int64) *SaveWhiteListStrategyRequest {
	s.StrategyId = &v
	return s
}

func (s *SaveWhiteListStrategyRequest) SetStrategyName(v string) *SaveWhiteListStrategyRequest {
	s.StrategyName = &v
	return s
}

func (s *SaveWhiteListStrategyRequest) SetStudyTime(v int32) *SaveWhiteListStrategyRequest {
	s.StudyTime = &v
	return s
}

func (s *SaveWhiteListStrategyRequest) Validate() error {
	return dara.Validate(s)
}
