// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTableLevelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTableLevelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTableLevelResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTableLevelResponseBody) *DeleteTableLevelResponse
	GetBody() *DeleteTableLevelResponseBody
}

type DeleteTableLevelResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTableLevelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTableLevelResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTableLevelResponse) GoString() string {
	return s.String()
}

func (s *DeleteTableLevelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTableLevelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTableLevelResponse) GetBody() *DeleteTableLevelResponseBody {
	return s.Body
}

func (s *DeleteTableLevelResponse) SetHeaders(v map[string]*string) *DeleteTableLevelResponse {
	s.Headers = v
	return s
}

func (s *DeleteTableLevelResponse) SetStatusCode(v int32) *DeleteTableLevelResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTableLevelResponse) SetBody(v *DeleteTableLevelResponseBody) *DeleteTableLevelResponse {
	s.Body = v
	return s
}

func (s *DeleteTableLevelResponse) Validate() error {
	return dara.Validate(s)
}
