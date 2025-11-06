// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportDefinitionAsynchronousResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImportDefinitionAsynchronousResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImportDefinitionAsynchronousResponse
	GetStatusCode() *int32
	SetBody(v *ImportDefinitionAsynchronousResponseBody) *ImportDefinitionAsynchronousResponse
	GetBody() *ImportDefinitionAsynchronousResponseBody
}

type ImportDefinitionAsynchronousResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportDefinitionAsynchronousResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportDefinitionAsynchronousResponse) String() string {
	return dara.Prettify(s)
}

func (s ImportDefinitionAsynchronousResponse) GoString() string {
	return s.String()
}

func (s *ImportDefinitionAsynchronousResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImportDefinitionAsynchronousResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImportDefinitionAsynchronousResponse) GetBody() *ImportDefinitionAsynchronousResponseBody {
	return s.Body
}

func (s *ImportDefinitionAsynchronousResponse) SetHeaders(v map[string]*string) *ImportDefinitionAsynchronousResponse {
	s.Headers = v
	return s
}

func (s *ImportDefinitionAsynchronousResponse) SetStatusCode(v int32) *ImportDefinitionAsynchronousResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportDefinitionAsynchronousResponse) SetBody(v *ImportDefinitionAsynchronousResponseBody) *ImportDefinitionAsynchronousResponse {
	s.Body = v
	return s
}

func (s *ImportDefinitionAsynchronousResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
