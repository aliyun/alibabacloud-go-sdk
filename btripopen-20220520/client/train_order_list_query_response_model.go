// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainOrderListQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TrainOrderListQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TrainOrderListQueryResponse
	GetStatusCode() *int32
	SetBody(v *TrainOrderListQueryResponseBody) *TrainOrderListQueryResponse
	GetBody() *TrainOrderListQueryResponseBody
}

type TrainOrderListQueryResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TrainOrderListQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TrainOrderListQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderListQueryResponse) GoString() string {
	return s.String()
}

func (s *TrainOrderListQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TrainOrderListQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TrainOrderListQueryResponse) GetBody() *TrainOrderListQueryResponseBody {
	return s.Body
}

func (s *TrainOrderListQueryResponse) SetHeaders(v map[string]*string) *TrainOrderListQueryResponse {
	s.Headers = v
	return s
}

func (s *TrainOrderListQueryResponse) SetStatusCode(v int32) *TrainOrderListQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *TrainOrderListQueryResponse) SetBody(v *TrainOrderListQueryResponseBody) *TrainOrderListQueryResponse {
	s.Body = v
	return s
}

func (s *TrainOrderListQueryResponse) Validate() error {
	return dara.Validate(s)
}
