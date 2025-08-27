// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainFeeCalculateRefundResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TrainFeeCalculateRefundResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TrainFeeCalculateRefundResponse
	GetStatusCode() *int32
	SetBody(v *TrainFeeCalculateRefundResponseBody) *TrainFeeCalculateRefundResponse
	GetBody() *TrainFeeCalculateRefundResponseBody
}

type TrainFeeCalculateRefundResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TrainFeeCalculateRefundResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TrainFeeCalculateRefundResponse) String() string {
	return dara.Prettify(s)
}

func (s TrainFeeCalculateRefundResponse) GoString() string {
	return s.String()
}

func (s *TrainFeeCalculateRefundResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TrainFeeCalculateRefundResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TrainFeeCalculateRefundResponse) GetBody() *TrainFeeCalculateRefundResponseBody {
	return s.Body
}

func (s *TrainFeeCalculateRefundResponse) SetHeaders(v map[string]*string) *TrainFeeCalculateRefundResponse {
	s.Headers = v
	return s
}

func (s *TrainFeeCalculateRefundResponse) SetStatusCode(v int32) *TrainFeeCalculateRefundResponse {
	s.StatusCode = &v
	return s
}

func (s *TrainFeeCalculateRefundResponse) SetBody(v *TrainFeeCalculateRefundResponseBody) *TrainFeeCalculateRefundResponse {
	s.Body = v
	return s
}

func (s *TrainFeeCalculateRefundResponse) Validate() error {
	return dara.Validate(s)
}
