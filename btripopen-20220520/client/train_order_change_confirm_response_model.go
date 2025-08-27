// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainOrderChangeConfirmResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TrainOrderChangeConfirmResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TrainOrderChangeConfirmResponse
	GetStatusCode() *int32
	SetBody(v *TrainOrderChangeConfirmResponseBody) *TrainOrderChangeConfirmResponse
	GetBody() *TrainOrderChangeConfirmResponseBody
}

type TrainOrderChangeConfirmResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TrainOrderChangeConfirmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TrainOrderChangeConfirmResponse) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderChangeConfirmResponse) GoString() string {
	return s.String()
}

func (s *TrainOrderChangeConfirmResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TrainOrderChangeConfirmResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TrainOrderChangeConfirmResponse) GetBody() *TrainOrderChangeConfirmResponseBody {
	return s.Body
}

func (s *TrainOrderChangeConfirmResponse) SetHeaders(v map[string]*string) *TrainOrderChangeConfirmResponse {
	s.Headers = v
	return s
}

func (s *TrainOrderChangeConfirmResponse) SetStatusCode(v int32) *TrainOrderChangeConfirmResponse {
	s.StatusCode = &v
	return s
}

func (s *TrainOrderChangeConfirmResponse) SetBody(v *TrainOrderChangeConfirmResponseBody) *TrainOrderChangeConfirmResponse {
	s.Body = v
	return s
}

func (s *TrainOrderChangeConfirmResponse) Validate() error {
	return dara.Validate(s)
}
