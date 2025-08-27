// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainStationSearchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *TrainStationSearchRequest
	GetKeyword() *string
}

type TrainStationSearchRequest struct {
	// This parameter is required.
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
}

func (s TrainStationSearchRequest) String() string {
	return dara.Prettify(s)
}

func (s TrainStationSearchRequest) GoString() string {
	return s.String()
}

func (s *TrainStationSearchRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *TrainStationSearchRequest) SetKeyword(v string) *TrainStationSearchRequest {
	s.Keyword = &v
	return s
}

func (s *TrainStationSearchRequest) Validate() error {
	return dara.Validate(s)
}
