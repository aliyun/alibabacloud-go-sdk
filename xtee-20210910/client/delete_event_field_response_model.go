// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEventFieldResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEventFieldResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEventFieldResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEventFieldResponseBody) *DeleteEventFieldResponse
	GetBody() *DeleteEventFieldResponseBody
}

type DeleteEventFieldResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEventFieldResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEventFieldResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventFieldResponse) GoString() string {
	return s.String()
}

func (s *DeleteEventFieldResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEventFieldResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEventFieldResponse) GetBody() *DeleteEventFieldResponseBody {
	return s.Body
}

func (s *DeleteEventFieldResponse) SetHeaders(v map[string]*string) *DeleteEventFieldResponse {
	s.Headers = v
	return s
}

func (s *DeleteEventFieldResponse) SetStatusCode(v int32) *DeleteEventFieldResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEventFieldResponse) SetBody(v *DeleteEventFieldResponseBody) *DeleteEventFieldResponse {
	s.Body = v
	return s
}

func (s *DeleteEventFieldResponse) Validate() error {
	return dara.Validate(s)
}
