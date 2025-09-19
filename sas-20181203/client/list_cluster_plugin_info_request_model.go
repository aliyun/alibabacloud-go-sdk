// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterPluginInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterIds(v []*string) *ListClusterPluginInfoRequest
	GetClusterIds() []*string
	SetLang(v string) *ListClusterPluginInfoRequest
	GetLang() *string
	SetPluginName(v string) *ListClusterPluginInfoRequest
	GetPluginName() *string
}

type ListClusterPluginInfoRequest struct {
	// The IDs of the clusters.
	//
	// This parameter is required.
	ClusterIds []*string `json:"ClusterIds,omitempty" xml:"ClusterIds,omitempty" type:"Repeated"`
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
	// The name of the plug-in.
	//
	// example:
	//
	// alihips
	PluginName *string `json:"PluginName,omitempty" xml:"PluginName,omitempty"`
}

func (s ListClusterPluginInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ListClusterPluginInfoRequest) GoString() string {
	return s.String()
}

func (s *ListClusterPluginInfoRequest) GetClusterIds() []*string {
	return s.ClusterIds
}

func (s *ListClusterPluginInfoRequest) GetLang() *string {
	return s.Lang
}

func (s *ListClusterPluginInfoRequest) GetPluginName() *string {
	return s.PluginName
}

func (s *ListClusterPluginInfoRequest) SetClusterIds(v []*string) *ListClusterPluginInfoRequest {
	s.ClusterIds = v
	return s
}

func (s *ListClusterPluginInfoRequest) SetLang(v string) *ListClusterPluginInfoRequest {
	s.Lang = &v
	return s
}

func (s *ListClusterPluginInfoRequest) SetPluginName(v string) *ListClusterPluginInfoRequest {
	s.PluginName = &v
	return s
}

func (s *ListClusterPluginInfoRequest) Validate() error {
	return dara.Validate(s)
}
