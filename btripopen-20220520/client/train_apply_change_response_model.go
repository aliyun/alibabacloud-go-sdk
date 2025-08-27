// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainApplyChangeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TrainApplyChangeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TrainApplyChangeResponse
	GetStatusCode() *int32
	SetBody(v *TrainApplyChangeResponseBody) *TrainApplyChangeResponse
	GetBody() *TrainApplyChangeResponseBody
}

type TrainApplyChangeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TrainApplyChangeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TrainApplyChangeResponse) String() string {
	return dara.Prettify(s)
}

func (s TrainApplyChangeResponse) GoString() string {
	return s.String()
}

func (s *TrainApplyChangeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TrainApplyChangeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TrainApplyChangeResponse) GetBody() *TrainApplyChangeResponseBody {
	return s.Body
}

func (s *TrainApplyChangeResponse) SetHeaders(v map[string]*string) *TrainApplyChangeResponse {
	s.Headers = v
	return s
}

func (s *TrainApplyChangeResponse) SetStatusCode(v int32) *TrainApplyChangeResponse {
	s.StatusCode = &v
	return s
}

func (s *TrainApplyChangeResponse) SetBody(v *TrainApplyChangeResponseBody) *TrainApplyChangeResponse {
	s.Body = v
	return s
}

func (s *TrainApplyChangeResponse) Validate() error {
	return dara.Validate(s)
}
