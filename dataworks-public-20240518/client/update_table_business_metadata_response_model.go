// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTableBusinessMetadataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTableBusinessMetadataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTableBusinessMetadataResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTableBusinessMetadataResponseBody) *UpdateTableBusinessMetadataResponse
	GetBody() *UpdateTableBusinessMetadataResponseBody
}

type UpdateTableBusinessMetadataResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTableBusinessMetadataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTableBusinessMetadataResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTableBusinessMetadataResponse) GoString() string {
	return s.String()
}

func (s *UpdateTableBusinessMetadataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTableBusinessMetadataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTableBusinessMetadataResponse) GetBody() *UpdateTableBusinessMetadataResponseBody {
	return s.Body
}

func (s *UpdateTableBusinessMetadataResponse) SetHeaders(v map[string]*string) *UpdateTableBusinessMetadataResponse {
	s.Headers = v
	return s
}

func (s *UpdateTableBusinessMetadataResponse) SetStatusCode(v int32) *UpdateTableBusinessMetadataResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTableBusinessMetadataResponse) SetBody(v *UpdateTableBusinessMetadataResponseBody) *UpdateTableBusinessMetadataResponse {
	s.Body = v
	return s
}

func (s *UpdateTableBusinessMetadataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
