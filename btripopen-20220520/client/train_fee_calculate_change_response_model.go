// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainFeeCalculateChangeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TrainFeeCalculateChangeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TrainFeeCalculateChangeResponse
	GetStatusCode() *int32
	SetBody(v *TrainFeeCalculateChangeResponseBody) *TrainFeeCalculateChangeResponse
	GetBody() *TrainFeeCalculateChangeResponseBody
}

type TrainFeeCalculateChangeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TrainFeeCalculateChangeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TrainFeeCalculateChangeResponse) String() string {
	return dara.Prettify(s)
}

func (s TrainFeeCalculateChangeResponse) GoString() string {
	return s.String()
}

func (s *TrainFeeCalculateChangeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TrainFeeCalculateChangeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TrainFeeCalculateChangeResponse) GetBody() *TrainFeeCalculateChangeResponseBody {
	return s.Body
}

func (s *TrainFeeCalculateChangeResponse) SetHeaders(v map[string]*string) *TrainFeeCalculateChangeResponse {
	s.Headers = v
	return s
}

func (s *TrainFeeCalculateChangeResponse) SetStatusCode(v int32) *TrainFeeCalculateChangeResponse {
	s.StatusCode = &v
	return s
}

func (s *TrainFeeCalculateChangeResponse) SetBody(v *TrainFeeCalculateChangeResponseBody) *TrainFeeCalculateChangeResponse {
	s.Body = v
	return s
}

func (s *TrainFeeCalculateChangeResponse) Validate() error {
	return dara.Validate(s)
}
