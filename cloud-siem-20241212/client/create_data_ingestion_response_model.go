// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataIngestionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDataIngestionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDataIngestionResponse
	GetStatusCode() *int32
	SetBody(v *CreateDataIngestionResponseBody) *CreateDataIngestionResponse
	GetBody() *CreateDataIngestionResponseBody
}

type CreateDataIngestionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataIngestionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDataIngestionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDataIngestionResponse) GoString() string {
	return s.String()
}

func (s *CreateDataIngestionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDataIngestionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDataIngestionResponse) GetBody() *CreateDataIngestionResponseBody {
	return s.Body
}

func (s *CreateDataIngestionResponse) SetHeaders(v map[string]*string) *CreateDataIngestionResponse {
	s.Headers = v
	return s
}

func (s *CreateDataIngestionResponse) SetStatusCode(v int32) *CreateDataIngestionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataIngestionResponse) SetBody(v *CreateDataIngestionResponseBody) *CreateDataIngestionResponse {
	s.Body = v
	return s
}

func (s *CreateDataIngestionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
