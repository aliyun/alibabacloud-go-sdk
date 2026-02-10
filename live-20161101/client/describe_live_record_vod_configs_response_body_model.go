// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveRecordVodConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLiveRecordVodConfigs(v *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigs) *DescribeLiveRecordVodConfigsResponseBody
	GetLiveRecordVodConfigs() *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigs
	SetPageNum(v int32) *DescribeLiveRecordVodConfigsResponseBody
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeLiveRecordVodConfigsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeLiveRecordVodConfigsResponseBody
	GetRequestId() *string
	SetTotal(v string) *DescribeLiveRecordVodConfigsResponseBody
	GetTotal() *string
}

type DescribeLiveRecordVodConfigsResponseBody struct {
	LiveRecordVodConfigs *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigs `json:"LiveRecordVodConfigs,omitempty" xml:"LiveRecordVodConfigs,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5056369B-D337-499E-B8B7-B761BD37B08A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	Total *string `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeLiveRecordVodConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveRecordVodConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveRecordVodConfigsResponseBody) GetLiveRecordVodConfigs() *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigs {
	return s.LiveRecordVodConfigs
}

func (s *DescribeLiveRecordVodConfigsResponseBody) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeLiveRecordVodConfigsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLiveRecordVodConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveRecordVodConfigsResponseBody) GetTotal() *string {
	return s.Total
}

func (s *DescribeLiveRecordVodConfigsResponseBody) SetLiveRecordVodConfigs(v *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigs) *DescribeLiveRecordVodConfigsResponseBody {
	s.LiveRecordVodConfigs = v
	return s
}

func (s *DescribeLiveRecordVodConfigsResponseBody) SetPageNum(v int32) *DescribeLiveRecordVodConfigsResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribeLiveRecordVodConfigsResponseBody) SetPageSize(v int32) *DescribeLiveRecordVodConfigsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveRecordVodConfigsResponseBody) SetRequestId(v string) *DescribeLiveRecordVodConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveRecordVodConfigsResponseBody) SetTotal(v string) *DescribeLiveRecordVodConfigsResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeLiveRecordVodConfigsResponseBody) Validate() error {
	if s.LiveRecordVodConfigs != nil {
		if err := s.LiveRecordVodConfigs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigs struct {
	LiveRecordVodConfig []*DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig `json:"LiveRecordVodConfig,omitempty" xml:"LiveRecordVodConfig,omitempty" type:"Repeated"`
}

func (s DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigs) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigs) GoString() string {
	return s.String()
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigs) GetLiveRecordVodConfig() []*DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig {
	return s.LiveRecordVodConfig
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigs) SetLiveRecordVodConfig(v []*DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig) *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigs {
	s.LiveRecordVodConfig = v
	return s
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigs) Validate() error {
	if s.LiveRecordVodConfig != nil {
		for _, item := range s.LiveRecordVodConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig struct {
	AppName                    *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AutoCompose                *string `json:"AutoCompose,omitempty" xml:"AutoCompose,omitempty"`
	ComposeVodTranscodeGroupId *string `json:"ComposeVodTranscodeGroupId,omitempty" xml:"ComposeVodTranscodeGroupId,omitempty"`
	CreateTime                 *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CycleDuration              *int32  `json:"CycleDuration,omitempty" xml:"CycleDuration,omitempty"`
	DomainName                 *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OnDemand                   *int32  `json:"OnDemand,omitempty" xml:"OnDemand,omitempty"`
	StorageLocation            *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
	StreamName                 *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	VodTranscodeGroupId        *string `json:"VodTranscodeGroupId,omitempty" xml:"VodTranscodeGroupId,omitempty"`
}

func (s DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig) GoString() string {
	return s.String()
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig) GetAppName() *string {
	return s.AppName
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig) GetAutoCompose() *string {
	return s.AutoCompose
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig) GetComposeVodTranscodeGroupId() *string {
	return s.ComposeVodTranscodeGroupId
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig) GetCycleDuration() *int32 {
	return s.CycleDuration
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig) GetOnDemand() *int32 {
	return s.OnDemand
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig) GetStorageLocation() *string {
	return s.StorageLocation
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig) GetVodTranscodeGroupId() *string {
	return s.VodTranscodeGroupId
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig) SetAppName(v string) *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig {
	s.AppName = &v
	return s
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig) SetAutoCompose(v string) *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig {
	s.AutoCompose = &v
	return s
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig) SetComposeVodTranscodeGroupId(v string) *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig {
	s.ComposeVodTranscodeGroupId = &v
	return s
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig) SetCreateTime(v string) *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig {
	s.CreateTime = &v
	return s
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig) SetCycleDuration(v int32) *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig {
	s.CycleDuration = &v
	return s
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig) SetDomainName(v string) *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig) SetOnDemand(v int32) *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig {
	s.OnDemand = &v
	return s
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig) SetStorageLocation(v string) *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig {
	s.StorageLocation = &v
	return s
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig) SetStreamName(v string) *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig) SetVodTranscodeGroupId(v string) *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig {
	s.VodTranscodeGroupId = &v
	return s
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig) Validate() error {
	return dara.Validate(s)
}
