// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDcdnUserConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetDcdnUserConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetDcdnUserConfigResponse
	GetStatusCode() *int32
	SetBody(v *SetDcdnUserConfigResponseBody) *SetDcdnUserConfigResponse
	GetBody() *SetDcdnUserConfigResponseBody
}

type SetDcdnUserConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDcdnUserConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDcdnUserConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s SetDcdnUserConfigResponse) GoString() string {
	return s.String()
}

func (s *SetDcdnUserConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetDcdnUserConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetDcdnUserConfigResponse) GetBody() *SetDcdnUserConfigResponseBody {
	return s.Body
}

func (s *SetDcdnUserConfigResponse) SetHeaders(v map[string]*string) *SetDcdnUserConfigResponse {
	s.Headers = v
	return s
}

func (s *SetDcdnUserConfigResponse) SetStatusCode(v int32) *SetDcdnUserConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDcdnUserConfigResponse) SetBody(v *SetDcdnUserConfigResponseBody) *SetDcdnUserConfigResponse {
	s.Body = v
	return s
}

func (s *SetDcdnUserConfigResponse) Validate() error {
	return dara.Validate(s)
}
