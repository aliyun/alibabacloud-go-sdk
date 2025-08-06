// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPlayConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPlayConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPlayConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetPlayConfigResponseBody) *GetPlayConfigResponse
	GetBody() *GetPlayConfigResponseBody
}

type GetPlayConfigResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPlayConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPlayConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPlayConfigResponse) GoString() string {
	return s.String()
}

func (s *GetPlayConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPlayConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPlayConfigResponse) GetBody() *GetPlayConfigResponseBody {
	return s.Body
}

func (s *GetPlayConfigResponse) SetHeaders(v map[string]*string) *GetPlayConfigResponse {
	s.Headers = v
	return s
}

func (s *GetPlayConfigResponse) SetStatusCode(v int32) *GetPlayConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPlayConfigResponse) SetBody(v *GetPlayConfigResponseBody) *GetPlayConfigResponse {
	s.Body = v
	return s
}

func (s *GetPlayConfigResponse) Validate() error {
	return dara.Validate(s)
}
