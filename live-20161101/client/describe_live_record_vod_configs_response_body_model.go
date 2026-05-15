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
	AppName                    *string                                                                                          `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AutoCompose                *string                                                                                          `json:"AutoCompose,omitempty" xml:"AutoCompose,omitempty"`
	ComposeVodTranscodeGroupId *string                                                                                          `json:"ComposeVodTranscodeGroupId,omitempty" xml:"ComposeVodTranscodeGroupId,omitempty"`
	CreateTime                 *string                                                                                          `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CycleDuration              *int32                                                                                           `json:"CycleDuration,omitempty" xml:"CycleDuration,omitempty"`
	DelayTime                  *int32                                                                                           `json:"DelayTime,omitempty" xml:"DelayTime,omitempty"`
	DomainName                 *string                                                                                          `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	FormatConfig               *bool                                                                                            `json:"FormatConfig,omitempty" xml:"FormatConfig,omitempty"`
	OnDemand                   *int32                                                                                           `json:"OnDemand,omitempty" xml:"OnDemand,omitempty"`
	RecordContent              *string                                                                                          `json:"RecordContent,omitempty" xml:"RecordContent,omitempty"`
	RecordFormatList           *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfigRecordFormatList `json:"RecordFormatList,omitempty" xml:"RecordFormatList,omitempty" type:"Struct"`
	SpaceId                    *string                                                                                          `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	StorageLocation            *string                                                                                          `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
	StreamName                 *string                                                                                          `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	TranscodeTemplates         *string                                                                                          `json:"TranscodeTemplates,omitempty" xml:"TranscodeTemplates,omitempty"`
	VodTranscodeGroupId        *string                                                                                          `json:"VodTranscodeGroupId,omitempty" xml:"VodTranscodeGroupId,omitempty"`
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

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig) GetDelayTime() *int32 {
	return s.DelayTime
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig) GetFormatConfig() *bool {
	return s.FormatConfig
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig) GetOnDemand() *int32 {
	return s.OnDemand
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig) GetRecordContent() *string {
	return s.RecordContent
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig) GetRecordFormatList() *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfigRecordFormatList {
	return s.RecordFormatList
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig) GetSpaceId() *string {
	return s.SpaceId
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig) GetStorageLocation() *string {
	return s.StorageLocation
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig) GetTranscodeTemplates() *string {
	return s.TranscodeTemplates
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

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig) SetDelayTime(v int32) *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig {
	s.DelayTime = &v
	return s
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig) SetDomainName(v string) *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig) SetFormatConfig(v bool) *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig {
	s.FormatConfig = &v
	return s
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig) SetOnDemand(v int32) *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig {
	s.OnDemand = &v
	return s
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig) SetRecordContent(v string) *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig {
	s.RecordContent = &v
	return s
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig) SetRecordFormatList(v *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfigRecordFormatList) *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig {
	s.RecordFormatList = v
	return s
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig) SetSpaceId(v string) *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig {
	s.SpaceId = &v
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

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig) SetTranscodeTemplates(v string) *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig {
	s.TranscodeTemplates = &v
	return s
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig) SetVodTranscodeGroupId(v string) *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig {
	s.VodTranscodeGroupId = &v
	return s
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfig) Validate() error {
	if s.RecordFormatList != nil {
		if err := s.RecordFormatList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfigRecordFormatList struct {
	RecordFormat []*DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfigRecordFormatListRecordFormat `json:"RecordFormat,omitempty" xml:"RecordFormat,omitempty" type:"Repeated"`
}

func (s DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfigRecordFormatList) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfigRecordFormatList) GoString() string {
	return s.String()
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfigRecordFormatList) GetRecordFormat() []*DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfigRecordFormatListRecordFormat {
	return s.RecordFormat
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfigRecordFormatList) SetRecordFormat(v []*DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfigRecordFormatListRecordFormat) *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfigRecordFormatList {
	s.RecordFormat = v
	return s
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfigRecordFormatList) Validate() error {
	if s.RecordFormat != nil {
		for _, item := range s.RecordFormat {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfigRecordFormatListRecordFormat struct {
	AutoCompose       *string `json:"AutoCompose,omitempty" xml:"AutoCompose,omitempty"`
	Format            *string `json:"Format,omitempty" xml:"Format,omitempty"`
	ProcessMethod     *string `json:"ProcessMethod,omitempty" xml:"ProcessMethod,omitempty"`
	ProcessTemplateId *string `json:"ProcessTemplateId,omitempty" xml:"ProcessTemplateId,omitempty"`
	SliceDuration     *int32  `json:"SliceDuration,omitempty" xml:"SliceDuration,omitempty"`
	Tags              *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	VideoProcess      *string `json:"VideoProcess,omitempty" xml:"VideoProcess,omitempty"`
}

func (s DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfigRecordFormatListRecordFormat) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfigRecordFormatListRecordFormat) GoString() string {
	return s.String()
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfigRecordFormatListRecordFormat) GetAutoCompose() *string {
	return s.AutoCompose
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfigRecordFormatListRecordFormat) GetFormat() *string {
	return s.Format
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfigRecordFormatListRecordFormat) GetProcessMethod() *string {
	return s.ProcessMethod
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfigRecordFormatListRecordFormat) GetProcessTemplateId() *string {
	return s.ProcessTemplateId
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfigRecordFormatListRecordFormat) GetSliceDuration() *int32 {
	return s.SliceDuration
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfigRecordFormatListRecordFormat) GetTags() *string {
	return s.Tags
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfigRecordFormatListRecordFormat) GetVideoProcess() *string {
	return s.VideoProcess
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfigRecordFormatListRecordFormat) SetAutoCompose(v string) *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfigRecordFormatListRecordFormat {
	s.AutoCompose = &v
	return s
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfigRecordFormatListRecordFormat) SetFormat(v string) *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfigRecordFormatListRecordFormat {
	s.Format = &v
	return s
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfigRecordFormatListRecordFormat) SetProcessMethod(v string) *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfigRecordFormatListRecordFormat {
	s.ProcessMethod = &v
	return s
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfigRecordFormatListRecordFormat) SetProcessTemplateId(v string) *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfigRecordFormatListRecordFormat {
	s.ProcessTemplateId = &v
	return s
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfigRecordFormatListRecordFormat) SetSliceDuration(v int32) *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfigRecordFormatListRecordFormat {
	s.SliceDuration = &v
	return s
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfigRecordFormatListRecordFormat) SetTags(v string) *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfigRecordFormatListRecordFormat {
	s.Tags = &v
	return s
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfigRecordFormatListRecordFormat) SetVideoProcess(v string) *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfigRecordFormatListRecordFormat {
	s.VideoProcess = &v
	return s
}

func (s *DescribeLiveRecordVodConfigsResponseBodyLiveRecordVodConfigsLiveRecordVodConfigRecordFormatListRecordFormat) Validate() error {
	return dara.Validate(s)
}
