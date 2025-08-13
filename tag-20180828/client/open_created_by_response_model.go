// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenCreatedByResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenCreatedByResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenCreatedByResponse
	GetStatusCode() *int32
	SetBody(v *OpenCreatedByResponseBody) *OpenCreatedByResponse
	GetBody() *OpenCreatedByResponseBody
}

type OpenCreatedByResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenCreatedByResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenCreatedByResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenCreatedByResponse) GoString() string {
	return s.String()
}

func (s *OpenCreatedByResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenCreatedByResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenCreatedByResponse) GetBody() *OpenCreatedByResponseBody {
	return s.Body
}

func (s *OpenCreatedByResponse) SetHeaders(v map[string]*string) *OpenCreatedByResponse {
	s.Headers = v
	return s
}

func (s *OpenCreatedByResponse) SetStatusCode(v int32) *OpenCreatedByResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenCreatedByResponse) SetBody(v *OpenCreatedByResponseBody) *OpenCreatedByResponse {
	s.Body = v
	return s
}

func (s *OpenCreatedByResponse) Validate() error {
	return dara.Validate(s)
}
