// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMapFromHavanaBindIdShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *MapFromHavanaBindIdShrinkRequest
	GetAppName() *string
	SetHavanaBindId(v string) *MapFromHavanaBindIdShrinkRequest
	GetHavanaBindId() *string
	SetHavanaBindStationsShrink(v string) *MapFromHavanaBindIdShrinkRequest
	GetHavanaBindStationsShrink() *string
}

type MapFromHavanaBindIdShrinkRequest struct {
	// This parameter is required.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	HavanaBindId *string `json:"HavanaBindId,omitempty" xml:"HavanaBindId,omitempty"`
	// This parameter is required.
	HavanaBindStationsShrink *string `json:"HavanaBindStations,omitempty" xml:"HavanaBindStations,omitempty"`
}

func (s MapFromHavanaBindIdShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s MapFromHavanaBindIdShrinkRequest) GoString() string {
	return s.String()
}

func (s *MapFromHavanaBindIdShrinkRequest) GetAppName() *string {
	return s.AppName
}

func (s *MapFromHavanaBindIdShrinkRequest) GetHavanaBindId() *string {
	return s.HavanaBindId
}

func (s *MapFromHavanaBindIdShrinkRequest) GetHavanaBindStationsShrink() *string {
	return s.HavanaBindStationsShrink
}

func (s *MapFromHavanaBindIdShrinkRequest) SetAppName(v string) *MapFromHavanaBindIdShrinkRequest {
	s.AppName = &v
	return s
}

func (s *MapFromHavanaBindIdShrinkRequest) SetHavanaBindId(v string) *MapFromHavanaBindIdShrinkRequest {
	s.HavanaBindId = &v
	return s
}

func (s *MapFromHavanaBindIdShrinkRequest) SetHavanaBindStationsShrink(v string) *MapFromHavanaBindIdShrinkRequest {
	s.HavanaBindStationsShrink = &v
	return s
}

func (s *MapFromHavanaBindIdShrinkRequest) Validate() error {
	return dara.Validate(s)
}
