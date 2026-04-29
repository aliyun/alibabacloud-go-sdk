// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarClawPluginsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DescribePolarClawPluginsRequest
	GetApplicationId() *string
	SetPluginList(v []*string) *DescribePolarClawPluginsRequest
	GetPluginList() []*string
}

type DescribePolarClawPluginsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string   `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	PluginList    []*string `json:"PluginList,omitempty" xml:"PluginList,omitempty" type:"Repeated"`
}

func (s DescribePolarClawPluginsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawPluginsRequest) GoString() string {
	return s.String()
}

func (s *DescribePolarClawPluginsRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DescribePolarClawPluginsRequest) GetPluginList() []*string {
	return s.PluginList
}

func (s *DescribePolarClawPluginsRequest) SetApplicationId(v string) *DescribePolarClawPluginsRequest {
	s.ApplicationId = &v
	return s
}

func (s *DescribePolarClawPluginsRequest) SetPluginList(v []*string) *DescribePolarClawPluginsRequest {
	s.PluginList = v
	return s
}

func (s *DescribePolarClawPluginsRequest) Validate() error {
	return dara.Validate(s)
}
