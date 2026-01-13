// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadApiCallDailyDetailCmd interface {
	dara.Model
	String() string
	GoString() string
	SetApiName(v string) *DownloadApiCallDailyDetailCmd
	GetApiName() *string
	SetEndTime(v string) *DownloadApiCallDailyDetailCmd
	GetEndTime() *string
	SetEngineType(v string) *DownloadApiCallDailyDetailCmd
	GetEngineType() *string
	SetStartTime(v string) *DownloadApiCallDailyDetailCmd
	GetStartTime() *string
}

type DownloadApiCallDailyDetailCmd struct {
	ApiName    *string `json:"apiName,omitempty" xml:"apiName,omitempty"`
	EndTime    *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	EngineType *string `json:"engineType,omitempty" xml:"engineType,omitempty"`
	StartTime  *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s DownloadApiCallDailyDetailCmd) String() string {
	return dara.Prettify(s)
}

func (s DownloadApiCallDailyDetailCmd) GoString() string {
	return s.String()
}

func (s *DownloadApiCallDailyDetailCmd) GetApiName() *string {
	return s.ApiName
}

func (s *DownloadApiCallDailyDetailCmd) GetEndTime() *string {
	return s.EndTime
}

func (s *DownloadApiCallDailyDetailCmd) GetEngineType() *string {
	return s.EngineType
}

func (s *DownloadApiCallDailyDetailCmd) GetStartTime() *string {
	return s.StartTime
}

func (s *DownloadApiCallDailyDetailCmd) SetApiName(v string) *DownloadApiCallDailyDetailCmd {
	s.ApiName = &v
	return s
}

func (s *DownloadApiCallDailyDetailCmd) SetEndTime(v string) *DownloadApiCallDailyDetailCmd {
	s.EndTime = &v
	return s
}

func (s *DownloadApiCallDailyDetailCmd) SetEngineType(v string) *DownloadApiCallDailyDetailCmd {
	s.EngineType = &v
	return s
}

func (s *DownloadApiCallDailyDetailCmd) SetStartTime(v string) *DownloadApiCallDailyDetailCmd {
	s.StartTime = &v
	return s
}

func (s *DownloadApiCallDailyDetailCmd) Validate() error {
	return dara.Validate(s)
}
