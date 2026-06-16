// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWebFetchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOutputFormat(v string) *WebFetchRequest
	GetOutputFormat() *string
	SetRegionId(v string) *WebFetchRequest
	GetRegionId() *string
	SetUrl(v string) *WebFetchRequest
	GetUrl() *string
}

type WebFetchRequest struct {
	OutputFormat *string `json:"OutputFormat,omitempty" xml:"OutputFormat,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s WebFetchRequest) String() string {
	return dara.Prettify(s)
}

func (s WebFetchRequest) GoString() string {
	return s.String()
}

func (s *WebFetchRequest) GetOutputFormat() *string {
	return s.OutputFormat
}

func (s *WebFetchRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *WebFetchRequest) GetUrl() *string {
	return s.Url
}

func (s *WebFetchRequest) SetOutputFormat(v string) *WebFetchRequest {
	s.OutputFormat = &v
	return s
}

func (s *WebFetchRequest) SetRegionId(v string) *WebFetchRequest {
	s.RegionId = &v
	return s
}

func (s *WebFetchRequest) SetUrl(v string) *WebFetchRequest {
	s.Url = &v
	return s
}

func (s *WebFetchRequest) Validate() error {
	return dara.Validate(s)
}
