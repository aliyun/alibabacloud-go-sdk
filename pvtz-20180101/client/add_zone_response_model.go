// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddZoneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddZoneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddZoneResponse
	GetStatusCode() *int32
	SetBody(v *AddZoneResponseBody) *AddZoneResponse
	GetBody() *AddZoneResponseBody
}

type AddZoneResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddZoneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddZoneResponse) String() string {
	return dara.Prettify(s)
}

func (s AddZoneResponse) GoString() string {
	return s.String()
}

func (s *AddZoneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddZoneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddZoneResponse) GetBody() *AddZoneResponseBody {
	return s.Body
}

func (s *AddZoneResponse) SetHeaders(v map[string]*string) *AddZoneResponse {
	s.Headers = v
	return s
}

func (s *AddZoneResponse) SetStatusCode(v int32) *AddZoneResponse {
	s.StatusCode = &v
	return s
}

func (s *AddZoneResponse) SetBody(v *AddZoneResponseBody) *AddZoneResponse {
	s.Body = v
	return s
}

func (s *AddZoneResponse) Validate() error {
	return dara.Validate(s)
}
