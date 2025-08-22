// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDcdnUserConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDcdnUserConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDcdnUserConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDcdnUserConfigResponseBody) *DeleteDcdnUserConfigResponse
	GetBody() *DeleteDcdnUserConfigResponseBody
}

type DeleteDcdnUserConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDcdnUserConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDcdnUserConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDcdnUserConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteDcdnUserConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDcdnUserConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDcdnUserConfigResponse) GetBody() *DeleteDcdnUserConfigResponseBody {
	return s.Body
}

func (s *DeleteDcdnUserConfigResponse) SetHeaders(v map[string]*string) *DeleteDcdnUserConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteDcdnUserConfigResponse) SetStatusCode(v int32) *DeleteDcdnUserConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDcdnUserConfigResponse) SetBody(v *DeleteDcdnUserConfigResponseBody) *DeleteDcdnUserConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteDcdnUserConfigResponse) Validate() error {
	return dara.Validate(s)
}
