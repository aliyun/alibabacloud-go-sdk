// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseCreatedByResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloseCreatedByResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloseCreatedByResponse
	GetStatusCode() *int32
	SetBody(v *CloseCreatedByResponseBody) *CloseCreatedByResponse
	GetBody() *CloseCreatedByResponseBody
}

type CloseCreatedByResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloseCreatedByResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloseCreatedByResponse) String() string {
	return dara.Prettify(s)
}

func (s CloseCreatedByResponse) GoString() string {
	return s.String()
}

func (s *CloseCreatedByResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloseCreatedByResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloseCreatedByResponse) GetBody() *CloseCreatedByResponseBody {
	return s.Body
}

func (s *CloseCreatedByResponse) SetHeaders(v map[string]*string) *CloseCreatedByResponse {
	s.Headers = v
	return s
}

func (s *CloseCreatedByResponse) SetStatusCode(v int32) *CloseCreatedByResponse {
	s.StatusCode = &v
	return s
}

func (s *CloseCreatedByResponse) SetBody(v *CloseCreatedByResponseBody) *CloseCreatedByResponse {
	s.Body = v
	return s
}

func (s *CloseCreatedByResponse) Validate() error {
	return dara.Validate(s)
}
