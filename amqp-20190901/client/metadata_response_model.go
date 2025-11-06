// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMetadataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MetadataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MetadataResponse
	GetStatusCode() *int32
	SetBody(v *MetadataResponseBody) *MetadataResponse
	GetBody() *MetadataResponseBody
}

type MetadataResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MetadataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MetadataResponse) String() string {
	return dara.Prettify(s)
}

func (s MetadataResponse) GoString() string {
	return s.String()
}

func (s *MetadataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MetadataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MetadataResponse) GetBody() *MetadataResponseBody {
	return s.Body
}

func (s *MetadataResponse) SetHeaders(v map[string]*string) *MetadataResponse {
	s.Headers = v
	return s
}

func (s *MetadataResponse) SetStatusCode(v int32) *MetadataResponse {
	s.StatusCode = &v
	return s
}

func (s *MetadataResponse) SetBody(v *MetadataResponseBody) *MetadataResponse {
	s.Body = v
	return s
}

func (s *MetadataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
