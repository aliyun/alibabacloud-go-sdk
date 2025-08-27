// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainOrderDetailQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TrainOrderDetailQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TrainOrderDetailQueryResponse
	GetStatusCode() *int32
	SetBody(v *TrainOrderDetailQueryResponseBody) *TrainOrderDetailQueryResponse
	GetBody() *TrainOrderDetailQueryResponseBody
}

type TrainOrderDetailQueryResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TrainOrderDetailQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TrainOrderDetailQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderDetailQueryResponse) GoString() string {
	return s.String()
}

func (s *TrainOrderDetailQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TrainOrderDetailQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TrainOrderDetailQueryResponse) GetBody() *TrainOrderDetailQueryResponseBody {
	return s.Body
}

func (s *TrainOrderDetailQueryResponse) SetHeaders(v map[string]*string) *TrainOrderDetailQueryResponse {
	s.Headers = v
	return s
}

func (s *TrainOrderDetailQueryResponse) SetStatusCode(v int32) *TrainOrderDetailQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *TrainOrderDetailQueryResponse) SetBody(v *TrainOrderDetailQueryResponseBody) *TrainOrderDetailQueryResponse {
	s.Body = v
	return s
}

func (s *TrainOrderDetailQueryResponse) Validate() error {
	return dara.Validate(s)
}
