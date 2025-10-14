// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataIngestionTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataIngestionTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataIngestionTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *ListDataIngestionTemplatesResponseBody) *ListDataIngestionTemplatesResponse
	GetBody() *ListDataIngestionTemplatesResponseBody
}

type ListDataIngestionTemplatesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataIngestionTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataIngestionTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataIngestionTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListDataIngestionTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataIngestionTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataIngestionTemplatesResponse) GetBody() *ListDataIngestionTemplatesResponseBody {
	return s.Body
}

func (s *ListDataIngestionTemplatesResponse) SetHeaders(v map[string]*string) *ListDataIngestionTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListDataIngestionTemplatesResponse) SetStatusCode(v int32) *ListDataIngestionTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataIngestionTemplatesResponse) SetBody(v *ListDataIngestionTemplatesResponseBody) *ListDataIngestionTemplatesResponse {
	s.Body = v
	return s
}

func (s *ListDataIngestionTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
