// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourceInstanceLabelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteResourceInstanceLabelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteResourceInstanceLabelResponse
	GetStatusCode() *int32
	SetBody(v *DeleteResourceInstanceLabelResponseBody) *DeleteResourceInstanceLabelResponse
	GetBody() *DeleteResourceInstanceLabelResponseBody
}

type DeleteResourceInstanceLabelResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteResourceInstanceLabelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteResourceInstanceLabelResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceInstanceLabelResponse) GoString() string {
	return s.String()
}

func (s *DeleteResourceInstanceLabelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteResourceInstanceLabelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteResourceInstanceLabelResponse) GetBody() *DeleteResourceInstanceLabelResponseBody {
	return s.Body
}

func (s *DeleteResourceInstanceLabelResponse) SetHeaders(v map[string]*string) *DeleteResourceInstanceLabelResponse {
	s.Headers = v
	return s
}

func (s *DeleteResourceInstanceLabelResponse) SetStatusCode(v int32) *DeleteResourceInstanceLabelResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteResourceInstanceLabelResponse) SetBody(v *DeleteResourceInstanceLabelResponseBody) *DeleteResourceInstanceLabelResponse {
	s.Body = v
	return s
}

func (s *DeleteResourceInstanceLabelResponse) Validate() error {
	return dara.Validate(s)
}
