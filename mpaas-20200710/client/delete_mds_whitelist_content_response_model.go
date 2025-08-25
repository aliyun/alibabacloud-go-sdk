// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMdsWhitelistContentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMdsWhitelistContentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMdsWhitelistContentResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMdsWhitelistContentResponseBody) *DeleteMdsWhitelistContentResponse
	GetBody() *DeleteMdsWhitelistContentResponseBody
}

type DeleteMdsWhitelistContentResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMdsWhitelistContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMdsWhitelistContentResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMdsWhitelistContentResponse) GoString() string {
	return s.String()
}

func (s *DeleteMdsWhitelistContentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMdsWhitelistContentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMdsWhitelistContentResponse) GetBody() *DeleteMdsWhitelistContentResponseBody {
	return s.Body
}

func (s *DeleteMdsWhitelistContentResponse) SetHeaders(v map[string]*string) *DeleteMdsWhitelistContentResponse {
	s.Headers = v
	return s
}

func (s *DeleteMdsWhitelistContentResponse) SetStatusCode(v int32) *DeleteMdsWhitelistContentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMdsWhitelistContentResponse) SetBody(v *DeleteMdsWhitelistContentResponseBody) *DeleteMdsWhitelistContentResponse {
	s.Body = v
	return s
}

func (s *DeleteMdsWhitelistContentResponse) Validate() error {
	return dara.Validate(s)
}
