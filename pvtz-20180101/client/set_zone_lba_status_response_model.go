// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetZoneLbaStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetZoneLbaStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetZoneLbaStatusResponse
	GetStatusCode() *int32
	SetBody(v *SetZoneLbaStatusResponseBody) *SetZoneLbaStatusResponse
	GetBody() *SetZoneLbaStatusResponseBody
}

type SetZoneLbaStatusResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetZoneLbaStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetZoneLbaStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s SetZoneLbaStatusResponse) GoString() string {
	return s.String()
}

func (s *SetZoneLbaStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetZoneLbaStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetZoneLbaStatusResponse) GetBody() *SetZoneLbaStatusResponseBody {
	return s.Body
}

func (s *SetZoneLbaStatusResponse) SetHeaders(v map[string]*string) *SetZoneLbaStatusResponse {
	s.Headers = v
	return s
}

func (s *SetZoneLbaStatusResponse) SetStatusCode(v int32) *SetZoneLbaStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *SetZoneLbaStatusResponse) SetBody(v *SetZoneLbaStatusResponseBody) *SetZoneLbaStatusResponse {
	s.Body = v
	return s
}

func (s *SetZoneLbaStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
