// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveSnapshotDetectPornConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLiveSnapshotDetectPornConfigList(v *DescribeLiveSnapshotDetectPornConfigResponseBodyLiveSnapshotDetectPornConfigList) *DescribeLiveSnapshotDetectPornConfigResponseBody
	GetLiveSnapshotDetectPornConfigList() *DescribeLiveSnapshotDetectPornConfigResponseBodyLiveSnapshotDetectPornConfigList
	SetOrder(v string) *DescribeLiveSnapshotDetectPornConfigResponseBody
	GetOrder() *string
	SetPageNum(v int32) *DescribeLiveSnapshotDetectPornConfigResponseBody
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeLiveSnapshotDetectPornConfigResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeLiveSnapshotDetectPornConfigResponseBody
	GetRequestId() *string
	SetTotalNum(v int32) *DescribeLiveSnapshotDetectPornConfigResponseBody
	GetTotalNum() *int32
	SetTotalPage(v int32) *DescribeLiveSnapshotDetectPornConfigResponseBody
	GetTotalPage() *int32
}

type DescribeLiveSnapshotDetectPornConfigResponseBody struct {
	LiveSnapshotDetectPornConfigList *DescribeLiveSnapshotDetectPornConfigResponseBodyLiveSnapshotDetectPornConfigList `json:"LiveSnapshotDetectPornConfigList,omitempty" xml:"LiveSnapshotDetectPornConfigList,omitempty" type:"Struct"`
	// The sort order.
	//
	// example:
	//
	// Asc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
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
	// 2
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries that meet the specified conditions.
	//
	// example:
	//
	// 6
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 11
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s DescribeLiveSnapshotDetectPornConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveSnapshotDetectPornConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveSnapshotDetectPornConfigResponseBody) GetLiveSnapshotDetectPornConfigList() *DescribeLiveSnapshotDetectPornConfigResponseBodyLiveSnapshotDetectPornConfigList {
	return s.LiveSnapshotDetectPornConfigList
}

func (s *DescribeLiveSnapshotDetectPornConfigResponseBody) GetOrder() *string {
	return s.Order
}

func (s *DescribeLiveSnapshotDetectPornConfigResponseBody) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeLiveSnapshotDetectPornConfigResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLiveSnapshotDetectPornConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveSnapshotDetectPornConfigResponseBody) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *DescribeLiveSnapshotDetectPornConfigResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeLiveSnapshotDetectPornConfigResponseBody) SetLiveSnapshotDetectPornConfigList(v *DescribeLiveSnapshotDetectPornConfigResponseBodyLiveSnapshotDetectPornConfigList) *DescribeLiveSnapshotDetectPornConfigResponseBody {
	s.LiveSnapshotDetectPornConfigList = v
	return s
}

func (s *DescribeLiveSnapshotDetectPornConfigResponseBody) SetOrder(v string) *DescribeLiveSnapshotDetectPornConfigResponseBody {
	s.Order = &v
	return s
}

func (s *DescribeLiveSnapshotDetectPornConfigResponseBody) SetPageNum(v int32) *DescribeLiveSnapshotDetectPornConfigResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribeLiveSnapshotDetectPornConfigResponseBody) SetPageSize(v int32) *DescribeLiveSnapshotDetectPornConfigResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveSnapshotDetectPornConfigResponseBody) SetRequestId(v string) *DescribeLiveSnapshotDetectPornConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveSnapshotDetectPornConfigResponseBody) SetTotalNum(v int32) *DescribeLiveSnapshotDetectPornConfigResponseBody {
	s.TotalNum = &v
	return s
}

func (s *DescribeLiveSnapshotDetectPornConfigResponseBody) SetTotalPage(v int32) *DescribeLiveSnapshotDetectPornConfigResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeLiveSnapshotDetectPornConfigResponseBody) Validate() error {
	if s.LiveSnapshotDetectPornConfigList != nil {
		if err := s.LiveSnapshotDetectPornConfigList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLiveSnapshotDetectPornConfigResponseBodyLiveSnapshotDetectPornConfigList struct {
	LiveSnapshotDetectPornConfig []*DescribeLiveSnapshotDetectPornConfigResponseBodyLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfig `json:"LiveSnapshotDetectPornConfig,omitempty" xml:"LiveSnapshotDetectPornConfig,omitempty" type:"Repeated"`
}

func (s DescribeLiveSnapshotDetectPornConfigResponseBodyLiveSnapshotDetectPornConfigList) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveSnapshotDetectPornConfigResponseBodyLiveSnapshotDetectPornConfigList) GoString() string {
	return s.String()
}

func (s *DescribeLiveSnapshotDetectPornConfigResponseBodyLiveSnapshotDetectPornConfigList) GetLiveSnapshotDetectPornConfig() []*DescribeLiveSnapshotDetectPornConfigResponseBodyLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfig {
	return s.LiveSnapshotDetectPornConfig
}

func (s *DescribeLiveSnapshotDetectPornConfigResponseBodyLiveSnapshotDetectPornConfigList) SetLiveSnapshotDetectPornConfig(v []*DescribeLiveSnapshotDetectPornConfigResponseBodyLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfig) *DescribeLiveSnapshotDetectPornConfigResponseBodyLiveSnapshotDetectPornConfigList {
	s.LiveSnapshotDetectPornConfig = v
	return s
}

func (s *DescribeLiveSnapshotDetectPornConfigResponseBodyLiveSnapshotDetectPornConfigList) Validate() error {
	if s.LiveSnapshotDetectPornConfig != nil {
		for _, item := range s.LiveSnapshotDetectPornConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLiveSnapshotDetectPornConfigResponseBodyLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfig struct {
	AppName     *string                                                                                                             `json:"AppName,omitempty" xml:"AppName,omitempty"`
	DomainName  *string                                                                                                             `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Interval    *int32                                                                                                              `json:"Interval,omitempty" xml:"Interval,omitempty"`
	OssBucket   *string                                                                                                             `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssEndpoint *string                                                                                                             `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
	OssObject   *string                                                                                                             `json:"OssObject,omitempty" xml:"OssObject,omitempty"`
	Scenes      *DescribeLiveSnapshotDetectPornConfigResponseBodyLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfigScenes `json:"Scenes,omitempty" xml:"Scenes,omitempty" type:"Struct"`
}

func (s DescribeLiveSnapshotDetectPornConfigResponseBodyLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveSnapshotDetectPornConfigResponseBodyLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfig) GoString() string {
	return s.String()
}

func (s *DescribeLiveSnapshotDetectPornConfigResponseBodyLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfig) GetAppName() *string {
	return s.AppName
}

func (s *DescribeLiveSnapshotDetectPornConfigResponseBodyLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfig) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveSnapshotDetectPornConfigResponseBodyLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfig) GetInterval() *int32 {
	return s.Interval
}

func (s *DescribeLiveSnapshotDetectPornConfigResponseBodyLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfig) GetOssBucket() *string {
	return s.OssBucket
}

func (s *DescribeLiveSnapshotDetectPornConfigResponseBodyLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfig) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *DescribeLiveSnapshotDetectPornConfigResponseBodyLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfig) GetOssObject() *string {
	return s.OssObject
}

func (s *DescribeLiveSnapshotDetectPornConfigResponseBodyLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfig) GetScenes() *DescribeLiveSnapshotDetectPornConfigResponseBodyLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfigScenes {
	return s.Scenes
}

func (s *DescribeLiveSnapshotDetectPornConfigResponseBodyLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfig) SetAppName(v string) *DescribeLiveSnapshotDetectPornConfigResponseBodyLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfig {
	s.AppName = &v
	return s
}

func (s *DescribeLiveSnapshotDetectPornConfigResponseBodyLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfig) SetDomainName(v string) *DescribeLiveSnapshotDetectPornConfigResponseBodyLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfig {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveSnapshotDetectPornConfigResponseBodyLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfig) SetInterval(v int32) *DescribeLiveSnapshotDetectPornConfigResponseBodyLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfig {
	s.Interval = &v
	return s
}

func (s *DescribeLiveSnapshotDetectPornConfigResponseBodyLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfig) SetOssBucket(v string) *DescribeLiveSnapshotDetectPornConfigResponseBodyLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfig {
	s.OssBucket = &v
	return s
}

func (s *DescribeLiveSnapshotDetectPornConfigResponseBodyLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfig) SetOssEndpoint(v string) *DescribeLiveSnapshotDetectPornConfigResponseBodyLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfig {
	s.OssEndpoint = &v
	return s
}

func (s *DescribeLiveSnapshotDetectPornConfigResponseBodyLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfig) SetOssObject(v string) *DescribeLiveSnapshotDetectPornConfigResponseBodyLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfig {
	s.OssObject = &v
	return s
}

func (s *DescribeLiveSnapshotDetectPornConfigResponseBodyLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfig) SetScenes(v *DescribeLiveSnapshotDetectPornConfigResponseBodyLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfigScenes) *DescribeLiveSnapshotDetectPornConfigResponseBodyLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfig {
	s.Scenes = v
	return s
}

func (s *DescribeLiveSnapshotDetectPornConfigResponseBodyLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfig) Validate() error {
	if s.Scenes != nil {
		if err := s.Scenes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLiveSnapshotDetectPornConfigResponseBodyLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfigScenes struct {
	Scene []*string `json:"scene,omitempty" xml:"scene,omitempty" type:"Repeated"`
}

func (s DescribeLiveSnapshotDetectPornConfigResponseBodyLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfigScenes) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveSnapshotDetectPornConfigResponseBodyLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfigScenes) GoString() string {
	return s.String()
}

func (s *DescribeLiveSnapshotDetectPornConfigResponseBodyLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfigScenes) GetScene() []*string {
	return s.Scene
}

func (s *DescribeLiveSnapshotDetectPornConfigResponseBodyLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfigScenes) SetScene(v []*string) *DescribeLiveSnapshotDetectPornConfigResponseBodyLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfigScenes {
	s.Scene = v
	return s
}

func (s *DescribeLiveSnapshotDetectPornConfigResponseBodyLiveSnapshotDetectPornConfigListLiveSnapshotDetectPornConfigScenes) Validate() error {
	return dara.Validate(s)
}
