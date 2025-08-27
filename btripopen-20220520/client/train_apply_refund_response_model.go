// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainApplyRefundResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TrainApplyRefundResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TrainApplyRefundResponse
	GetStatusCode() *int32
	SetBody(v *TrainApplyRefundResponseBody) *TrainApplyRefundResponse
	GetBody() *TrainApplyRefundResponseBody
}

type TrainApplyRefundResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TrainApplyRefundResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TrainApplyRefundResponse) String() string {
	return dara.Prettify(s)
}

func (s TrainApplyRefundResponse) GoString() string {
	return s.String()
}

func (s *TrainApplyRefundResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TrainApplyRefundResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TrainApplyRefundResponse) GetBody() *TrainApplyRefundResponseBody {
	return s.Body
}

func (s *TrainApplyRefundResponse) SetHeaders(v map[string]*string) *TrainApplyRefundResponse {
	s.Headers = v
	return s
}

func (s *TrainApplyRefundResponse) SetStatusCode(v int32) *TrainApplyRefundResponse {
	s.StatusCode = &v
	return s
}

func (s *TrainApplyRefundResponse) SetBody(v *TrainApplyRefundResponseBody) *TrainApplyRefundResponse {
	s.Body = v
	return s
}

func (s *TrainApplyRefundResponse) Validate() error {
	return dara.Validate(s)
}
