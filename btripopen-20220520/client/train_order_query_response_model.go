// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainOrderQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TrainOrderQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TrainOrderQueryResponse
	GetStatusCode() *int32
	SetBody(v *TrainOrderQueryResponseBody) *TrainOrderQueryResponse
	GetBody() *TrainOrderQueryResponseBody
}

type TrainOrderQueryResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TrainOrderQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TrainOrderQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderQueryResponse) GoString() string {
	return s.String()
}

func (s *TrainOrderQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TrainOrderQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TrainOrderQueryResponse) GetBody() *TrainOrderQueryResponseBody {
	return s.Body
}

func (s *TrainOrderQueryResponse) SetHeaders(v map[string]*string) *TrainOrderQueryResponse {
	s.Headers = v
	return s
}

func (s *TrainOrderQueryResponse) SetStatusCode(v int32) *TrainOrderQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *TrainOrderQueryResponse) SetBody(v *TrainOrderQueryResponseBody) *TrainOrderQueryResponse {
	s.Body = v
	return s
}

func (s *TrainOrderQueryResponse) Validate() error {
	return dara.Validate(s)
}
