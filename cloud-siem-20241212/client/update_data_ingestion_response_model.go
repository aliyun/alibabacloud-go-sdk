// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataIngestionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDataIngestionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDataIngestionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDataIngestionResponseBody) *UpdateDataIngestionResponse
	GetBody() *UpdateDataIngestionResponseBody
}

type UpdateDataIngestionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDataIngestionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDataIngestionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataIngestionResponse) GoString() string {
	return s.String()
}

func (s *UpdateDataIngestionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDataIngestionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDataIngestionResponse) GetBody() *UpdateDataIngestionResponseBody {
	return s.Body
}

func (s *UpdateDataIngestionResponse) SetHeaders(v map[string]*string) *UpdateDataIngestionResponse {
	s.Headers = v
	return s
}

func (s *UpdateDataIngestionResponse) SetStatusCode(v int32) *UpdateDataIngestionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDataIngestionResponse) SetBody(v *UpdateDataIngestionResponseBody) *UpdateDataIngestionResponse {
	s.Body = v
	return s
}

func (s *UpdateDataIngestionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
