// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTrainTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTrainTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTrainTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateTrainTaskResponseBody) *CreateTrainTaskResponse
	GetBody() *CreateTrainTaskResponseBody
}

type CreateTrainTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTrainTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTrainTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTrainTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateTrainTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTrainTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTrainTaskResponse) GetBody() *CreateTrainTaskResponseBody {
	return s.Body
}

func (s *CreateTrainTaskResponse) SetHeaders(v map[string]*string) *CreateTrainTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateTrainTaskResponse) SetStatusCode(v int32) *CreateTrainTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTrainTaskResponse) SetBody(v *CreateTrainTaskResponseBody) *CreateTrainTaskResponse {
	s.Body = v
	return s
}

func (s *CreateTrainTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
