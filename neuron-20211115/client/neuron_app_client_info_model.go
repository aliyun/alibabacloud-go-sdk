// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNeuronAppClientInfo interface {
	dara.Model
	String() string
	GoString() string
	SetAppEntry(v string) *NeuronAppClientInfo
	GetAppEntry() *string
	SetAppId(v int64) *NeuronAppClientInfo
	GetAppId() *int64
	SetMobiInfo(v *MobiInfo) *NeuronAppClientInfo
	GetMobiInfo() *MobiInfo
	SetPluginList(v []*AppPluginInfo) *NeuronAppClientInfo
	GetPluginList() []*AppPluginInfo
	SetProductId(v int64) *NeuronAppClientInfo
	GetProductId() *int64
	SetSidebar(v string) *NeuronAppClientInfo
	GetSidebar() *string
	SetVersion(v string) *NeuronAppClientInfo
	GetVersion() *string
}

type NeuronAppClientInfo struct {
	// example:
	//
	// module_58mkf3jj::page_gbnxyuh5
	AppEntry *string `json:"appEntry,omitempty" xml:"appEntry,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	AppId      *int64           `json:"appId,omitempty" xml:"appId,omitempty"`
	MobiInfo   *MobiInfo        `json:"mobiInfo,omitempty" xml:"mobiInfo,omitempty"`
	PluginList []*AppPluginInfo `json:"pluginList,omitempty" xml:"pluginList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	ProductId *int64 `json:"productId,omitempty" xml:"productId,omitempty"`
	// example:
	//
	// {}
	Sidebar *string `json:"sidebar,omitempty" xml:"sidebar,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1.0.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s NeuronAppClientInfo) String() string {
	return dara.Prettify(s)
}

func (s NeuronAppClientInfo) GoString() string {
	return s.String()
}

func (s *NeuronAppClientInfo) GetAppEntry() *string {
	return s.AppEntry
}

func (s *NeuronAppClientInfo) GetAppId() *int64 {
	return s.AppId
}

func (s *NeuronAppClientInfo) GetMobiInfo() *MobiInfo {
	return s.MobiInfo
}

func (s *NeuronAppClientInfo) GetPluginList() []*AppPluginInfo {
	return s.PluginList
}

func (s *NeuronAppClientInfo) GetProductId() *int64 {
	return s.ProductId
}

func (s *NeuronAppClientInfo) GetSidebar() *string {
	return s.Sidebar
}

func (s *NeuronAppClientInfo) GetVersion() *string {
	return s.Version
}

func (s *NeuronAppClientInfo) SetAppEntry(v string) *NeuronAppClientInfo {
	s.AppEntry = &v
	return s
}

func (s *NeuronAppClientInfo) SetAppId(v int64) *NeuronAppClientInfo {
	s.AppId = &v
	return s
}

func (s *NeuronAppClientInfo) SetMobiInfo(v *MobiInfo) *NeuronAppClientInfo {
	s.MobiInfo = v
	return s
}

func (s *NeuronAppClientInfo) SetPluginList(v []*AppPluginInfo) *NeuronAppClientInfo {
	s.PluginList = v
	return s
}

func (s *NeuronAppClientInfo) SetProductId(v int64) *NeuronAppClientInfo {
	s.ProductId = &v
	return s
}

func (s *NeuronAppClientInfo) SetSidebar(v string) *NeuronAppClientInfo {
	s.Sidebar = &v
	return s
}

func (s *NeuronAppClientInfo) SetVersion(v string) *NeuronAppClientInfo {
	s.Version = &v
	return s
}

func (s *NeuronAppClientInfo) Validate() error {
	if s.MobiInfo != nil {
		if err := s.MobiInfo.Validate(); err != nil {
			return err
		}
	}
	if s.PluginList != nil {
		for _, item := range s.PluginList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
