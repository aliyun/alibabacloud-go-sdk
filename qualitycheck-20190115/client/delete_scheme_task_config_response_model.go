// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSchemeTaskConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSchemeTaskConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSchemeTaskConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSchemeTaskConfigResponseBody) *DeleteSchemeTaskConfigResponse
	GetBody() *DeleteSchemeTaskConfigResponseBody
}

type DeleteSchemeTaskConfigResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSchemeTaskConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSchemeTaskConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSchemeTaskConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteSchemeTaskConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSchemeTaskConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSchemeTaskConfigResponse) GetBody() *DeleteSchemeTaskConfigResponseBody {
	return s.Body
}

func (s *DeleteSchemeTaskConfigResponse) SetHeaders(v map[string]*string) *DeleteSchemeTaskConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteSchemeTaskConfigResponse) SetStatusCode(v int32) *DeleteSchemeTaskConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSchemeTaskConfigResponse) SetBody(v *DeleteSchemeTaskConfigResponseBody) *DeleteSchemeTaskConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteSchemeTaskConfigResponse) Validate() error {
	return dara.Validate(s)
}
