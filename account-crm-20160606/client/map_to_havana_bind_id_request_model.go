// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMapToHavanaBindIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *MapToHavanaBindIdRequest
	GetAppName() *string
	SetHavanaBindStations(v map[string]interface{}) *MapToHavanaBindIdRequest
	GetHavanaBindStations() map[string]interface{}
	SetPk(v string) *MapToHavanaBindIdRequest
	GetPk() *string
}

type MapToHavanaBindIdRequest struct {
	// This parameter is required.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	HavanaBindStations map[string]interface{} `json:"HavanaBindStations,omitempty" xml:"HavanaBindStations,omitempty"`
	// This parameter is required.
	Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
}

func (s MapToHavanaBindIdRequest) String() string {
	return dara.Prettify(s)
}

func (s MapToHavanaBindIdRequest) GoString() string {
	return s.String()
}

func (s *MapToHavanaBindIdRequest) GetAppName() *string {
	return s.AppName
}

func (s *MapToHavanaBindIdRequest) GetHavanaBindStations() map[string]interface{} {
	return s.HavanaBindStations
}

func (s *MapToHavanaBindIdRequest) GetPk() *string {
	return s.Pk
}

func (s *MapToHavanaBindIdRequest) SetAppName(v string) *MapToHavanaBindIdRequest {
	s.AppName = &v
	return s
}

func (s *MapToHavanaBindIdRequest) SetHavanaBindStations(v map[string]interface{}) *MapToHavanaBindIdRequest {
	s.HavanaBindStations = v
	return s
}

func (s *MapToHavanaBindIdRequest) SetPk(v string) *MapToHavanaBindIdRequest {
	s.Pk = &v
	return s
}

func (s *MapToHavanaBindIdRequest) Validate() error {
	return dara.Validate(s)
}
