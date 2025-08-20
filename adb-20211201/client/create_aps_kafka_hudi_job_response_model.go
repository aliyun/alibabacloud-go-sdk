// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApsKafkaHudiJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateApsKafkaHudiJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateApsKafkaHudiJobResponse
	GetStatusCode() *int32
	SetBody(v *CreateApsKafkaHudiJobResponseBody) *CreateApsKafkaHudiJobResponse
	GetBody() *CreateApsKafkaHudiJobResponseBody
}

type CreateApsKafkaHudiJobResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateApsKafkaHudiJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateApsKafkaHudiJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateApsKafkaHudiJobResponse) GoString() string {
	return s.String()
}

func (s *CreateApsKafkaHudiJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateApsKafkaHudiJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateApsKafkaHudiJobResponse) GetBody() *CreateApsKafkaHudiJobResponseBody {
	return s.Body
}

func (s *CreateApsKafkaHudiJobResponse) SetHeaders(v map[string]*string) *CreateApsKafkaHudiJobResponse {
	s.Headers = v
	return s
}

func (s *CreateApsKafkaHudiJobResponse) SetStatusCode(v int32) *CreateApsKafkaHudiJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateApsKafkaHudiJobResponse) SetBody(v *CreateApsKafkaHudiJobResponseBody) *CreateApsKafkaHudiJobResponse {
	s.Body = v
	return s
}

func (s *CreateApsKafkaHudiJobResponse) Validate() error {
	return dara.Validate(s)
}
