// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCubecardWhitelistContentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCubecardWhitelistContentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCubecardWhitelistContentResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCubecardWhitelistContentResponseBody) *DeleteCubecardWhitelistContentResponse
	GetBody() *DeleteCubecardWhitelistContentResponseBody
}

type DeleteCubecardWhitelistContentResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCubecardWhitelistContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCubecardWhitelistContentResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCubecardWhitelistContentResponse) GoString() string {
	return s.String()
}

func (s *DeleteCubecardWhitelistContentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCubecardWhitelistContentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCubecardWhitelistContentResponse) GetBody() *DeleteCubecardWhitelistContentResponseBody {
	return s.Body
}

func (s *DeleteCubecardWhitelistContentResponse) SetHeaders(v map[string]*string) *DeleteCubecardWhitelistContentResponse {
	s.Headers = v
	return s
}

func (s *DeleteCubecardWhitelistContentResponse) SetStatusCode(v int32) *DeleteCubecardWhitelistContentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCubecardWhitelistContentResponse) SetBody(v *DeleteCubecardWhitelistContentResponseBody) *DeleteCubecardWhitelistContentResponse {
	s.Body = v
	return s
}

func (s *DeleteCubecardWhitelistContentResponse) Validate() error {
	return dara.Validate(s)
}
