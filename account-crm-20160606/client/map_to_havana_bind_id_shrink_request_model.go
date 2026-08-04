// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMapToHavanaBindIdShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *MapToHavanaBindIdShrinkRequest
	GetAppName() *string
	SetHavanaBindStationsShrink(v string) *MapToHavanaBindIdShrinkRequest
	GetHavanaBindStationsShrink() *string
	SetPk(v string) *MapToHavanaBindIdShrinkRequest
	GetPk() *string
}

type MapToHavanaBindIdShrinkRequest struct {
	// This parameter is required.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	HavanaBindStationsShrink *string `json:"HavanaBindStations,omitempty" xml:"HavanaBindStations,omitempty"`
	// This parameter is required.
	Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
}

func (s MapToHavanaBindIdShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s MapToHavanaBindIdShrinkRequest) GoString() string {
	return s.String()
}

func (s *MapToHavanaBindIdShrinkRequest) GetAppName() *string {
	return s.AppName
}

func (s *MapToHavanaBindIdShrinkRequest) GetHavanaBindStationsShrink() *string {
	return s.HavanaBindStationsShrink
}

func (s *MapToHavanaBindIdShrinkRequest) GetPk() *string {
	return s.Pk
}

func (s *MapToHavanaBindIdShrinkRequest) SetAppName(v string) *MapToHavanaBindIdShrinkRequest {
	s.AppName = &v
	return s
}

func (s *MapToHavanaBindIdShrinkRequest) SetHavanaBindStationsShrink(v string) *MapToHavanaBindIdShrinkRequest {
	s.HavanaBindStationsShrink = &v
	return s
}

func (s *MapToHavanaBindIdShrinkRequest) SetPk(v string) *MapToHavanaBindIdShrinkRequest {
	s.Pk = &v
	return s
}

func (s *MapToHavanaBindIdShrinkRequest) Validate() error {
	return dara.Validate(s)
}
