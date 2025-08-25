// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMcubeWhitelistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMcubeWhitelistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMcubeWhitelistResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMcubeWhitelistResponseBody) *DeleteMcubeWhitelistResponse
	GetBody() *DeleteMcubeWhitelistResponseBody
}

type DeleteMcubeWhitelistResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMcubeWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMcubeWhitelistResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMcubeWhitelistResponse) GoString() string {
	return s.String()
}

func (s *DeleteMcubeWhitelistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMcubeWhitelistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMcubeWhitelistResponse) GetBody() *DeleteMcubeWhitelistResponseBody {
	return s.Body
}

func (s *DeleteMcubeWhitelistResponse) SetHeaders(v map[string]*string) *DeleteMcubeWhitelistResponse {
	s.Headers = v
	return s
}

func (s *DeleteMcubeWhitelistResponse) SetStatusCode(v int32) *DeleteMcubeWhitelistResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMcubeWhitelistResponse) SetBody(v *DeleteMcubeWhitelistResponseBody) *DeleteMcubeWhitelistResponse {
	s.Body = v
	return s
}

func (s *DeleteMcubeWhitelistResponse) Validate() error {
	return dara.Validate(s)
}
