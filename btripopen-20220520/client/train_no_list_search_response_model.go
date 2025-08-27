// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainNoListSearchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TrainNoListSearchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TrainNoListSearchResponse
	GetStatusCode() *int32
	SetBody(v *TrainNoListSearchResponseBody) *TrainNoListSearchResponse
	GetBody() *TrainNoListSearchResponseBody
}

type TrainNoListSearchResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TrainNoListSearchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TrainNoListSearchResponse) String() string {
	return dara.Prettify(s)
}

func (s TrainNoListSearchResponse) GoString() string {
	return s.String()
}

func (s *TrainNoListSearchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TrainNoListSearchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TrainNoListSearchResponse) GetBody() *TrainNoListSearchResponseBody {
	return s.Body
}

func (s *TrainNoListSearchResponse) SetHeaders(v map[string]*string) *TrainNoListSearchResponse {
	s.Headers = v
	return s
}

func (s *TrainNoListSearchResponse) SetStatusCode(v int32) *TrainNoListSearchResponse {
	s.StatusCode = &v
	return s
}

func (s *TrainNoListSearchResponse) SetBody(v *TrainNoListSearchResponseBody) *TrainNoListSearchResponse {
	s.Body = v
	return s
}

func (s *TrainNoListSearchResponse) Validate() error {
	return dara.Validate(s)
}
