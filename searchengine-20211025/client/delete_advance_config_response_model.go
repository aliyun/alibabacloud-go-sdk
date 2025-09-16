// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAdvanceConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAdvanceConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAdvanceConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAdvanceConfigResponseBody) *DeleteAdvanceConfigResponse
	GetBody() *DeleteAdvanceConfigResponseBody
}

type DeleteAdvanceConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAdvanceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAdvanceConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAdvanceConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteAdvanceConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAdvanceConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAdvanceConfigResponse) GetBody() *DeleteAdvanceConfigResponseBody {
	return s.Body
}

func (s *DeleteAdvanceConfigResponse) SetHeaders(v map[string]*string) *DeleteAdvanceConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteAdvanceConfigResponse) SetStatusCode(v int32) *DeleteAdvanceConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAdvanceConfigResponse) SetBody(v *DeleteAdvanceConfigResponseBody) *DeleteAdvanceConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteAdvanceConfigResponse) Validate() error {
	return dara.Validate(s)
}
