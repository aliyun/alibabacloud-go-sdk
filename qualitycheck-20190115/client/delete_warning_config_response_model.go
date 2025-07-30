// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWarningConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteWarningConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteWarningConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteWarningConfigResponseBody) *DeleteWarningConfigResponse
	GetBody() *DeleteWarningConfigResponseBody
}

type DeleteWarningConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWarningConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWarningConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteWarningConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteWarningConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteWarningConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteWarningConfigResponse) GetBody() *DeleteWarningConfigResponseBody {
	return s.Body
}

func (s *DeleteWarningConfigResponse) SetHeaders(v map[string]*string) *DeleteWarningConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteWarningConfigResponse) SetStatusCode(v int32) *DeleteWarningConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWarningConfigResponse) SetBody(v *DeleteWarningConfigResponseBody) *DeleteWarningConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteWarningConfigResponse) Validate() error {
	return dara.Validate(s)
}
