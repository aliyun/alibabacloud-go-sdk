// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainOrderPayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TrainOrderPayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TrainOrderPayResponse
	GetStatusCode() *int32
	SetBody(v *TrainOrderPayResponseBody) *TrainOrderPayResponse
	GetBody() *TrainOrderPayResponseBody
}

type TrainOrderPayResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TrainOrderPayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TrainOrderPayResponse) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderPayResponse) GoString() string {
	return s.String()
}

func (s *TrainOrderPayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TrainOrderPayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TrainOrderPayResponse) GetBody() *TrainOrderPayResponseBody {
	return s.Body
}

func (s *TrainOrderPayResponse) SetHeaders(v map[string]*string) *TrainOrderPayResponse {
	s.Headers = v
	return s
}

func (s *TrainOrderPayResponse) SetStatusCode(v int32) *TrainOrderPayResponse {
	s.StatusCode = &v
	return s
}

func (s *TrainOrderPayResponse) SetBody(v *TrainOrderPayResponseBody) *TrainOrderPayResponse {
	s.Body = v
	return s
}

func (s *TrainOrderPayResponse) Validate() error {
	return dara.Validate(s)
}
