// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomCallTaggingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCustomCallTaggingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCustomCallTaggingResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCustomCallTaggingResponseBody) *DeleteCustomCallTaggingResponse
	GetBody() *DeleteCustomCallTaggingResponseBody
}

type DeleteCustomCallTaggingResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustomCallTaggingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustomCallTaggingResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomCallTaggingResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomCallTaggingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCustomCallTaggingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCustomCallTaggingResponse) GetBody() *DeleteCustomCallTaggingResponseBody {
	return s.Body
}

func (s *DeleteCustomCallTaggingResponse) SetHeaders(v map[string]*string) *DeleteCustomCallTaggingResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomCallTaggingResponse) SetStatusCode(v int32) *DeleteCustomCallTaggingResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomCallTaggingResponse) SetBody(v *DeleteCustomCallTaggingResponseBody) *DeleteCustomCallTaggingResponse {
	s.Body = v
	return s
}

func (s *DeleteCustomCallTaggingResponse) Validate() error {
	return dara.Validate(s)
}
