// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDcdnIpaSpecificConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDcdnIpaSpecificConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDcdnIpaSpecificConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDcdnIpaSpecificConfigResponseBody) *DeleteDcdnIpaSpecificConfigResponse
	GetBody() *DeleteDcdnIpaSpecificConfigResponseBody
}

type DeleteDcdnIpaSpecificConfigResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDcdnIpaSpecificConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDcdnIpaSpecificConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDcdnIpaSpecificConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteDcdnIpaSpecificConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDcdnIpaSpecificConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDcdnIpaSpecificConfigResponse) GetBody() *DeleteDcdnIpaSpecificConfigResponseBody {
	return s.Body
}

func (s *DeleteDcdnIpaSpecificConfigResponse) SetHeaders(v map[string]*string) *DeleteDcdnIpaSpecificConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteDcdnIpaSpecificConfigResponse) SetStatusCode(v int32) *DeleteDcdnIpaSpecificConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDcdnIpaSpecificConfigResponse) SetBody(v *DeleteDcdnIpaSpecificConfigResponseBody) *DeleteDcdnIpaSpecificConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteDcdnIpaSpecificConfigResponse) Validate() error {
	return dara.Validate(s)
}
