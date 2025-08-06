// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFilterConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilterName(v string) *AddFilterConfigsRequest
	GetFilterName() *string
	SetItemConfigs(v string) *AddFilterConfigsRequest
	GetItemConfigs() *string
	SetType(v string) *AddFilterConfigsRequest
	GetType() *string
}

type AddFilterConfigsRequest struct {
	FilterName  *string `json:"FilterName,omitempty" xml:"FilterName,omitempty"`
	ItemConfigs *string `json:"ItemConfigs,omitempty" xml:"ItemConfigs,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AddFilterConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s AddFilterConfigsRequest) GoString() string {
	return s.String()
}

func (s *AddFilterConfigsRequest) GetFilterName() *string {
	return s.FilterName
}

func (s *AddFilterConfigsRequest) GetItemConfigs() *string {
	return s.ItemConfigs
}

func (s *AddFilterConfigsRequest) GetType() *string {
	return s.Type
}

func (s *AddFilterConfigsRequest) SetFilterName(v string) *AddFilterConfigsRequest {
	s.FilterName = &v
	return s
}

func (s *AddFilterConfigsRequest) SetItemConfigs(v string) *AddFilterConfigsRequest {
	s.ItemConfigs = &v
	return s
}

func (s *AddFilterConfigsRequest) SetType(v string) *AddFilterConfigsRequest {
	s.Type = &v
	return s
}

func (s *AddFilterConfigsRequest) Validate() error {
	return dara.Validate(s)
}
