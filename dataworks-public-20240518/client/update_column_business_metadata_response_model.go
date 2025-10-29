// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateColumnBusinessMetadataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateColumnBusinessMetadataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateColumnBusinessMetadataResponse
	GetStatusCode() *int32
	SetBody(v *UpdateColumnBusinessMetadataResponseBody) *UpdateColumnBusinessMetadataResponse
	GetBody() *UpdateColumnBusinessMetadataResponseBody
}

type UpdateColumnBusinessMetadataResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateColumnBusinessMetadataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateColumnBusinessMetadataResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateColumnBusinessMetadataResponse) GoString() string {
	return s.String()
}

func (s *UpdateColumnBusinessMetadataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateColumnBusinessMetadataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateColumnBusinessMetadataResponse) GetBody() *UpdateColumnBusinessMetadataResponseBody {
	return s.Body
}

func (s *UpdateColumnBusinessMetadataResponse) SetHeaders(v map[string]*string) *UpdateColumnBusinessMetadataResponse {
	s.Headers = v
	return s
}

func (s *UpdateColumnBusinessMetadataResponse) SetStatusCode(v int32) *UpdateColumnBusinessMetadataResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateColumnBusinessMetadataResponse) SetBody(v *UpdateColumnBusinessMetadataResponseBody) *UpdateColumnBusinessMetadataResponse {
	s.Body = v
	return s
}

func (s *UpdateColumnBusinessMetadataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
