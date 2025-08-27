// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainStationSearchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TrainStationSearchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TrainStationSearchResponse
	GetStatusCode() *int32
	SetBody(v *TrainStationSearchResponseBody) *TrainStationSearchResponse
	GetBody() *TrainStationSearchResponseBody
}

type TrainStationSearchResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TrainStationSearchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TrainStationSearchResponse) String() string {
	return dara.Prettify(s)
}

func (s TrainStationSearchResponse) GoString() string {
	return s.String()
}

func (s *TrainStationSearchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TrainStationSearchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TrainStationSearchResponse) GetBody() *TrainStationSearchResponseBody {
	return s.Body
}

func (s *TrainStationSearchResponse) SetHeaders(v map[string]*string) *TrainStationSearchResponse {
	s.Headers = v
	return s
}

func (s *TrainStationSearchResponse) SetStatusCode(v int32) *TrainStationSearchResponse {
	s.StatusCode = &v
	return s
}

func (s *TrainStationSearchResponse) SetBody(v *TrainStationSearchResponseBody) *TrainStationSearchResponse {
	s.Body = v
	return s
}

func (s *TrainStationSearchResponse) Validate() error {
	return dara.Validate(s)
}
