// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMassPushResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MassPushResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MassPushResponse
	GetStatusCode() *int32
	SetBody(v *MassPushResponseBody) *MassPushResponse
	GetBody() *MassPushResponseBody
}

type MassPushResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MassPushResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MassPushResponse) String() string {
	return dara.Prettify(s)
}

func (s MassPushResponse) GoString() string {
	return s.String()
}

func (s *MassPushResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MassPushResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MassPushResponse) GetBody() *MassPushResponseBody {
	return s.Body
}

func (s *MassPushResponse) SetHeaders(v map[string]*string) *MassPushResponse {
	s.Headers = v
	return s
}

func (s *MassPushResponse) SetStatusCode(v int32) *MassPushResponse {
	s.StatusCode = &v
	return s
}

func (s *MassPushResponse) SetBody(v *MassPushResponseBody) *MassPushResponse {
	s.Body = v
	return s
}

func (s *MassPushResponse) Validate() error {
	return dara.Validate(s)
}
