// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainNoInfoSearchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TrainNoInfoSearchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TrainNoInfoSearchResponse
	GetStatusCode() *int32
	SetBody(v *TrainNoInfoSearchResponseBody) *TrainNoInfoSearchResponse
	GetBody() *TrainNoInfoSearchResponseBody
}

type TrainNoInfoSearchResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TrainNoInfoSearchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TrainNoInfoSearchResponse) String() string {
	return dara.Prettify(s)
}

func (s TrainNoInfoSearchResponse) GoString() string {
	return s.String()
}

func (s *TrainNoInfoSearchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TrainNoInfoSearchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TrainNoInfoSearchResponse) GetBody() *TrainNoInfoSearchResponseBody {
	return s.Body
}

func (s *TrainNoInfoSearchResponse) SetHeaders(v map[string]*string) *TrainNoInfoSearchResponse {
	s.Headers = v
	return s
}

func (s *TrainNoInfoSearchResponse) SetStatusCode(v int32) *TrainNoInfoSearchResponse {
	s.StatusCode = &v
	return s
}

func (s *TrainNoInfoSearchResponse) SetBody(v *TrainNoInfoSearchResponseBody) *TrainNoInfoSearchResponse {
	s.Body = v
	return s
}

func (s *TrainNoInfoSearchResponse) Validate() error {
	return dara.Validate(s)
}
