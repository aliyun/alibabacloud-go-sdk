// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMapFromHavanaBindIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *MapFromHavanaBindIdRequest
	GetAppName() *string
	SetHavanaBindId(v string) *MapFromHavanaBindIdRequest
	GetHavanaBindId() *string
	SetHavanaBindStations(v map[string]interface{}) *MapFromHavanaBindIdRequest
	GetHavanaBindStations() map[string]interface{}
}

type MapFromHavanaBindIdRequest struct {
	// This parameter is required.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	HavanaBindId *string `json:"HavanaBindId,omitempty" xml:"HavanaBindId,omitempty"`
	// This parameter is required.
	HavanaBindStations map[string]interface{} `json:"HavanaBindStations,omitempty" xml:"HavanaBindStations,omitempty"`
}

func (s MapFromHavanaBindIdRequest) String() string {
	return dara.Prettify(s)
}

func (s MapFromHavanaBindIdRequest) GoString() string {
	return s.String()
}

func (s *MapFromHavanaBindIdRequest) GetAppName() *string {
	return s.AppName
}

func (s *MapFromHavanaBindIdRequest) GetHavanaBindId() *string {
	return s.HavanaBindId
}

func (s *MapFromHavanaBindIdRequest) GetHavanaBindStations() map[string]interface{} {
	return s.HavanaBindStations
}

func (s *MapFromHavanaBindIdRequest) SetAppName(v string) *MapFromHavanaBindIdRequest {
	s.AppName = &v
	return s
}

func (s *MapFromHavanaBindIdRequest) SetHavanaBindId(v string) *MapFromHavanaBindIdRequest {
	s.HavanaBindId = &v
	return s
}

func (s *MapFromHavanaBindIdRequest) SetHavanaBindStations(v map[string]interface{}) *MapFromHavanaBindIdRequest {
	s.HavanaBindStations = v
	return s
}

func (s *MapFromHavanaBindIdRequest) Validate() error {
	return dara.Validate(s)
}
