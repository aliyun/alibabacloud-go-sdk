// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBaselineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBaselineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBaselineResponse
	GetStatusCode() *int32
	SetBody(v *DeleteBaselineResponseBody) *DeleteBaselineResponse
	GetBody() *DeleteBaselineResponseBody
}

type DeleteBaselineResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBaselineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBaselineResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBaselineResponse) GoString() string {
	return s.String()
}

func (s *DeleteBaselineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBaselineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBaselineResponse) GetBody() *DeleteBaselineResponseBody {
	return s.Body
}

func (s *DeleteBaselineResponse) SetHeaders(v map[string]*string) *DeleteBaselineResponse {
	s.Headers = v
	return s
}

func (s *DeleteBaselineResponse) SetStatusCode(v int32) *DeleteBaselineResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBaselineResponse) SetBody(v *DeleteBaselineResponseBody) *DeleteBaselineResponse {
	s.Body = v
	return s
}

func (s *DeleteBaselineResponse) Validate() error {
	return dara.Validate(s)
}
