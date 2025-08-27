// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAirportSearchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *AirportSearchRequest
	GetKeyword() *string
	SetType(v int32) *AirportSearchRequest
	GetType() *int32
}

type AirportSearchRequest struct {
	// This parameter is required.
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// example:
	//
	// 0
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s AirportSearchRequest) String() string {
	return dara.Prettify(s)
}

func (s AirportSearchRequest) GoString() string {
	return s.String()
}

func (s *AirportSearchRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *AirportSearchRequest) GetType() *int32 {
	return s.Type
}

func (s *AirportSearchRequest) SetKeyword(v string) *AirportSearchRequest {
	s.Keyword = &v
	return s
}

func (s *AirportSearchRequest) SetType(v int32) *AirportSearchRequest {
	s.Type = &v
	return s
}

func (s *AirportSearchRequest) Validate() error {
	return dara.Validate(s)
}
