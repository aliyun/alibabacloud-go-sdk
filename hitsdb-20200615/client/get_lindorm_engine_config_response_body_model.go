// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLindormEngineConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetLindormEngineConfigResponseBody
	GetAccessDeniedDetail() *string
	SetEngineConfigs(v []*GetLindormEngineConfigResponseBodyEngineConfigs) *GetLindormEngineConfigResponseBody
	GetEngineConfigs() []*GetLindormEngineConfigResponseBodyEngineConfigs
	SetRequestId(v string) *GetLindormEngineConfigResponseBody
	GetRequestId() *string
}

type GetLindormEngineConfigResponseBody struct {
	AccessDeniedDetail *string                                            `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	EngineConfigs      []*GetLindormEngineConfigResponseBodyEngineConfigs `json:"EngineConfigs,omitempty" xml:"EngineConfigs,omitempty" type:"Repeated"`
	RequestId          *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetLindormEngineConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLindormEngineConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetLindormEngineConfigResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetLindormEngineConfigResponseBody) GetEngineConfigs() []*GetLindormEngineConfigResponseBodyEngineConfigs {
	return s.EngineConfigs
}

func (s *GetLindormEngineConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLindormEngineConfigResponseBody) SetAccessDeniedDetail(v string) *GetLindormEngineConfigResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetLindormEngineConfigResponseBody) SetEngineConfigs(v []*GetLindormEngineConfigResponseBodyEngineConfigs) *GetLindormEngineConfigResponseBody {
	s.EngineConfigs = v
	return s
}

func (s *GetLindormEngineConfigResponseBody) SetRequestId(v string) *GetLindormEngineConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLindormEngineConfigResponseBody) Validate() error {
	if s.EngineConfigs != nil {
		for _, item := range s.EngineConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetLindormEngineConfigResponseBodyEngineConfigs struct {
	ConfigFiles []*GetLindormEngineConfigResponseBodyEngineConfigsConfigFiles `json:"ConfigFiles,omitempty" xml:"ConfigFiles,omitempty" type:"Repeated"`
	Engine      *string                                                       `json:"Engine,omitempty" xml:"Engine,omitempty"`
}

func (s GetLindormEngineConfigResponseBodyEngineConfigs) String() string {
	return dara.Prettify(s)
}

func (s GetLindormEngineConfigResponseBodyEngineConfigs) GoString() string {
	return s.String()
}

func (s *GetLindormEngineConfigResponseBodyEngineConfigs) GetConfigFiles() []*GetLindormEngineConfigResponseBodyEngineConfigsConfigFiles {
	return s.ConfigFiles
}

func (s *GetLindormEngineConfigResponseBodyEngineConfigs) GetEngine() *string {
	return s.Engine
}

func (s *GetLindormEngineConfigResponseBodyEngineConfigs) SetConfigFiles(v []*GetLindormEngineConfigResponseBodyEngineConfigsConfigFiles) *GetLindormEngineConfigResponseBodyEngineConfigs {
	s.ConfigFiles = v
	return s
}

func (s *GetLindormEngineConfigResponseBodyEngineConfigs) SetEngine(v string) *GetLindormEngineConfigResponseBodyEngineConfigs {
	s.Engine = &v
	return s
}

func (s *GetLindormEngineConfigResponseBodyEngineConfigs) Validate() error {
	if s.ConfigFiles != nil {
		for _, item := range s.ConfigFiles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetLindormEngineConfigResponseBodyEngineConfigsConfigFiles struct {
	ConfigItems []*GetLindormEngineConfigResponseBodyEngineConfigsConfigFilesConfigItems `json:"ConfigItems,omitempty" xml:"ConfigItems,omitempty" type:"Repeated"`
	FileName    *string                                                                  `json:"FileName,omitempty" xml:"FileName,omitempty"`
}

func (s GetLindormEngineConfigResponseBodyEngineConfigsConfigFiles) String() string {
	return dara.Prettify(s)
}

func (s GetLindormEngineConfigResponseBodyEngineConfigsConfigFiles) GoString() string {
	return s.String()
}

func (s *GetLindormEngineConfigResponseBodyEngineConfigsConfigFiles) GetConfigItems() []*GetLindormEngineConfigResponseBodyEngineConfigsConfigFilesConfigItems {
	return s.ConfigItems
}

func (s *GetLindormEngineConfigResponseBodyEngineConfigsConfigFiles) GetFileName() *string {
	return s.FileName
}

func (s *GetLindormEngineConfigResponseBodyEngineConfigsConfigFiles) SetConfigItems(v []*GetLindormEngineConfigResponseBodyEngineConfigsConfigFilesConfigItems) *GetLindormEngineConfigResponseBodyEngineConfigsConfigFiles {
	s.ConfigItems = v
	return s
}

func (s *GetLindormEngineConfigResponseBodyEngineConfigsConfigFiles) SetFileName(v string) *GetLindormEngineConfigResponseBodyEngineConfigsConfigFiles {
	s.FileName = &v
	return s
}

func (s *GetLindormEngineConfigResponseBodyEngineConfigsConfigFiles) Validate() error {
	if s.ConfigItems != nil {
		for _, item := range s.ConfigItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetLindormEngineConfigResponseBodyEngineConfigsConfigFilesConfigItems struct {
	ConfigItemKey   *string `json:"ConfigItemKey,omitempty" xml:"ConfigItemKey,omitempty"`
	ConfigItemValue *string `json:"ConfigItemValue,omitempty" xml:"ConfigItemValue,omitempty"`
}

func (s GetLindormEngineConfigResponseBodyEngineConfigsConfigFilesConfigItems) String() string {
	return dara.Prettify(s)
}

func (s GetLindormEngineConfigResponseBodyEngineConfigsConfigFilesConfigItems) GoString() string {
	return s.String()
}

func (s *GetLindormEngineConfigResponseBodyEngineConfigsConfigFilesConfigItems) GetConfigItemKey() *string {
	return s.ConfigItemKey
}

func (s *GetLindormEngineConfigResponseBodyEngineConfigsConfigFilesConfigItems) GetConfigItemValue() *string {
	return s.ConfigItemValue
}

func (s *GetLindormEngineConfigResponseBodyEngineConfigsConfigFilesConfigItems) SetConfigItemKey(v string) *GetLindormEngineConfigResponseBodyEngineConfigsConfigFilesConfigItems {
	s.ConfigItemKey = &v
	return s
}

func (s *GetLindormEngineConfigResponseBodyEngineConfigsConfigFilesConfigItems) SetConfigItemValue(v string) *GetLindormEngineConfigResponseBodyEngineConfigsConfigFilesConfigItems {
	s.ConfigItemValue = &v
	return s
}

func (s *GetLindormEngineConfigResponseBodyEngineConfigsConfigFilesConfigItems) Validate() error {
	return dara.Validate(s)
}
