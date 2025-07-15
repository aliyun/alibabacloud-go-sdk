// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportCustomCallTaggingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImportCustomCallTaggingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImportCustomCallTaggingResponse
	GetStatusCode() *int32
	SetBody(v *ImportCustomCallTaggingResponseBody) *ImportCustomCallTaggingResponse
	GetBody() *ImportCustomCallTaggingResponseBody
}

type ImportCustomCallTaggingResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportCustomCallTaggingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportCustomCallTaggingResponse) String() string {
	return dara.Prettify(s)
}

func (s ImportCustomCallTaggingResponse) GoString() string {
	return s.String()
}

func (s *ImportCustomCallTaggingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImportCustomCallTaggingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImportCustomCallTaggingResponse) GetBody() *ImportCustomCallTaggingResponseBody {
	return s.Body
}

func (s *ImportCustomCallTaggingResponse) SetHeaders(v map[string]*string) *ImportCustomCallTaggingResponse {
	s.Headers = v
	return s
}

func (s *ImportCustomCallTaggingResponse) SetStatusCode(v int32) *ImportCustomCallTaggingResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportCustomCallTaggingResponse) SetBody(v *ImportCustomCallTaggingResponseBody) *ImportCustomCallTaggingResponse {
	s.Body = v
	return s
}

func (s *ImportCustomCallTaggingResponse) Validate() error {
	return dara.Validate(s)
}
