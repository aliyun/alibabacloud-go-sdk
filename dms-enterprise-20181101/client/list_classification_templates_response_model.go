// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClassificationTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListClassificationTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListClassificationTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *ListClassificationTemplatesResponseBody) *ListClassificationTemplatesResponse
	GetBody() *ListClassificationTemplatesResponseBody
}

type ListClassificationTemplatesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClassificationTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClassificationTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListClassificationTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListClassificationTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListClassificationTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListClassificationTemplatesResponse) GetBody() *ListClassificationTemplatesResponseBody {
	return s.Body
}

func (s *ListClassificationTemplatesResponse) SetHeaders(v map[string]*string) *ListClassificationTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListClassificationTemplatesResponse) SetStatusCode(v int32) *ListClassificationTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClassificationTemplatesResponse) SetBody(v *ListClassificationTemplatesResponseBody) *ListClassificationTemplatesResponse {
	s.Body = v
	return s
}

func (s *ListClassificationTemplatesResponse) Validate() error {
	return dara.Validate(s)
}
