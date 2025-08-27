// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainStopoverSearchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TrainStopoverSearchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TrainStopoverSearchResponse
	GetStatusCode() *int32
	SetBody(v *TrainStopoverSearchResponseBody) *TrainStopoverSearchResponse
	GetBody() *TrainStopoverSearchResponseBody
}

type TrainStopoverSearchResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TrainStopoverSearchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TrainStopoverSearchResponse) String() string {
	return dara.Prettify(s)
}

func (s TrainStopoverSearchResponse) GoString() string {
	return s.String()
}

func (s *TrainStopoverSearchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TrainStopoverSearchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TrainStopoverSearchResponse) GetBody() *TrainStopoverSearchResponseBody {
	return s.Body
}

func (s *TrainStopoverSearchResponse) SetHeaders(v map[string]*string) *TrainStopoverSearchResponse {
	s.Headers = v
	return s
}

func (s *TrainStopoverSearchResponse) SetStatusCode(v int32) *TrainStopoverSearchResponse {
	s.StatusCode = &v
	return s
}

func (s *TrainStopoverSearchResponse) SetBody(v *TrainStopoverSearchResponseBody) *TrainStopoverSearchResponse {
	s.Body = v
	return s
}

func (s *TrainStopoverSearchResponse) Validate() error {
	return dara.Validate(s)
}
