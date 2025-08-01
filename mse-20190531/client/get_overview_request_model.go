// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOverviewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *GetOverviewRequest
	GetAcceptLanguage() *string
	SetPeriod(v int32) *GetOverviewRequest
	GetPeriod() *int32
	SetRegion(v string) *GetOverviewRequest
	GetRegion() *string
}

type GetOverviewRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The query time. Unit: days. For example, if you set this parameter to 30, the governance rules within the last 30 days are queried.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s GetOverviewRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOverviewRequest) GoString() string {
	return s.String()
}

func (s *GetOverviewRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *GetOverviewRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *GetOverviewRequest) GetRegion() *string {
	return s.Region
}

func (s *GetOverviewRequest) SetAcceptLanguage(v string) *GetOverviewRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *GetOverviewRequest) SetPeriod(v int32) *GetOverviewRequest {
	s.Period = &v
	return s
}

func (s *GetOverviewRequest) SetRegion(v string) *GetOverviewRequest {
	s.Region = &v
	return s
}

func (s *GetOverviewRequest) Validate() error {
	return dara.Validate(s)
}
