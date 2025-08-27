// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainOrderCancelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TrainOrderCancelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TrainOrderCancelResponse
	GetStatusCode() *int32
	SetBody(v *TrainOrderCancelResponseBody) *TrainOrderCancelResponse
	GetBody() *TrainOrderCancelResponseBody
}

type TrainOrderCancelResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TrainOrderCancelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TrainOrderCancelResponse) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderCancelResponse) GoString() string {
	return s.String()
}

func (s *TrainOrderCancelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TrainOrderCancelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TrainOrderCancelResponse) GetBody() *TrainOrderCancelResponseBody {
	return s.Body
}

func (s *TrainOrderCancelResponse) SetHeaders(v map[string]*string) *TrainOrderCancelResponse {
	s.Headers = v
	return s
}

func (s *TrainOrderCancelResponse) SetStatusCode(v int32) *TrainOrderCancelResponse {
	s.StatusCode = &v
	return s
}

func (s *TrainOrderCancelResponse) SetBody(v *TrainOrderCancelResponseBody) *TrainOrderCancelResponse {
	s.Body = v
	return s
}

func (s *TrainOrderCancelResponse) Validate() error {
	return dara.Validate(s)
}
