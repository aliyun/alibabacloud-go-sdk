// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRunLabelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRunLabelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRunLabelResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRunLabelResponseBody) *DeleteRunLabelResponse
	GetBody() *DeleteRunLabelResponseBody
}

type DeleteRunLabelResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRunLabelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRunLabelResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRunLabelResponse) GoString() string {
	return s.String()
}

func (s *DeleteRunLabelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRunLabelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRunLabelResponse) GetBody() *DeleteRunLabelResponseBody {
	return s.Body
}

func (s *DeleteRunLabelResponse) SetHeaders(v map[string]*string) *DeleteRunLabelResponse {
	s.Headers = v
	return s
}

func (s *DeleteRunLabelResponse) SetStatusCode(v int32) *DeleteRunLabelResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRunLabelResponse) SetBody(v *DeleteRunLabelResponseBody) *DeleteRunLabelResponse {
	s.Body = v
	return s
}

func (s *DeleteRunLabelResponse) Validate() error {
	return dara.Validate(s)
}
