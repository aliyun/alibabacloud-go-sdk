// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLabelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLabelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLabelResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLabelResponseBody) *DeleteLabelResponse
	GetBody() *DeleteLabelResponseBody
}

type DeleteLabelResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLabelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLabelResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLabelResponse) GoString() string {
	return s.String()
}

func (s *DeleteLabelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLabelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLabelResponse) GetBody() *DeleteLabelResponseBody {
	return s.Body
}

func (s *DeleteLabelResponse) SetHeaders(v map[string]*string) *DeleteLabelResponse {
	s.Headers = v
	return s
}

func (s *DeleteLabelResponse) SetStatusCode(v int32) *DeleteLabelResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLabelResponse) SetBody(v *DeleteLabelResponseBody) *DeleteLabelResponse {
	s.Body = v
	return s
}

func (s *DeleteLabelResponse) Validate() error {
	return dara.Validate(s)
}
