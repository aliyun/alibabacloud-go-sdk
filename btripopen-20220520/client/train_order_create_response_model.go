// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainOrderCreateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TrainOrderCreateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TrainOrderCreateResponse
	GetStatusCode() *int32
	SetBody(v *TrainOrderCreateResponseBody) *TrainOrderCreateResponse
	GetBody() *TrainOrderCreateResponseBody
}

type TrainOrderCreateResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TrainOrderCreateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TrainOrderCreateResponse) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderCreateResponse) GoString() string {
	return s.String()
}

func (s *TrainOrderCreateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TrainOrderCreateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TrainOrderCreateResponse) GetBody() *TrainOrderCreateResponseBody {
	return s.Body
}

func (s *TrainOrderCreateResponse) SetHeaders(v map[string]*string) *TrainOrderCreateResponse {
	s.Headers = v
	return s
}

func (s *TrainOrderCreateResponse) SetStatusCode(v int32) *TrainOrderCreateResponse {
	s.StatusCode = &v
	return s
}

func (s *TrainOrderCreateResponse) SetBody(v *TrainOrderCreateResponseBody) *TrainOrderCreateResponse {
	s.Body = v
	return s
}

func (s *TrainOrderCreateResponse) Validate() error {
	return dara.Validate(s)
}
