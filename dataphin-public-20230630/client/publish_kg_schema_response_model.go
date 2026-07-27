// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishKgSchemaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PublishKgSchemaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PublishKgSchemaResponse
	GetStatusCode() *int32
	SetBody(v *PublishKgSchemaResponseBody) *PublishKgSchemaResponse
	GetBody() *PublishKgSchemaResponseBody
}

type PublishKgSchemaResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishKgSchemaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishKgSchemaResponse) String() string {
	return dara.Prettify(s)
}

func (s PublishKgSchemaResponse) GoString() string {
	return s.String()
}

func (s *PublishKgSchemaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PublishKgSchemaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PublishKgSchemaResponse) GetBody() *PublishKgSchemaResponseBody {
	return s.Body
}

func (s *PublishKgSchemaResponse) SetHeaders(v map[string]*string) *PublishKgSchemaResponse {
	s.Headers = v
	return s
}

func (s *PublishKgSchemaResponse) SetStatusCode(v int32) *PublishKgSchemaResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishKgSchemaResponse) SetBody(v *PublishKgSchemaResponseBody) *PublishKgSchemaResponse {
	s.Body = v
	return s
}

func (s *PublishKgSchemaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
