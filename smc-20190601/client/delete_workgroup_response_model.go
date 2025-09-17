// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkgroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteWorkgroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteWorkgroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteWorkgroupResponseBody) *DeleteWorkgroupResponse
	GetBody() *DeleteWorkgroupResponseBody
}

type DeleteWorkgroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWorkgroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWorkgroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkgroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteWorkgroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteWorkgroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteWorkgroupResponse) GetBody() *DeleteWorkgroupResponseBody {
	return s.Body
}

func (s *DeleteWorkgroupResponse) SetHeaders(v map[string]*string) *DeleteWorkgroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteWorkgroupResponse) SetStatusCode(v int32) *DeleteWorkgroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWorkgroupResponse) SetBody(v *DeleteWorkgroupResponseBody) *DeleteWorkgroupResponse {
	s.Body = v
	return s
}

func (s *DeleteWorkgroupResponse) Validate() error {
	return dara.Validate(s)
}
