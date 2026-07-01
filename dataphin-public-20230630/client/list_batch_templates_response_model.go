// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBatchTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBatchTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBatchTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *ListBatchTemplatesResponseBody) *ListBatchTemplatesResponse
	GetBody() *ListBatchTemplatesResponseBody
}

type ListBatchTemplatesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBatchTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBatchTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBatchTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListBatchTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBatchTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBatchTemplatesResponse) GetBody() *ListBatchTemplatesResponseBody {
	return s.Body
}

func (s *ListBatchTemplatesResponse) SetHeaders(v map[string]*string) *ListBatchTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListBatchTemplatesResponse) SetStatusCode(v int32) *ListBatchTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBatchTemplatesResponse) SetBody(v *ListBatchTemplatesResponseBody) *ListBatchTemplatesResponse {
	s.Body = v
	return s
}

func (s *ListBatchTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
