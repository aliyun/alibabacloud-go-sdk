// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVccRouteEntryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVccRouteEntryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVccRouteEntryResponse
	GetStatusCode() *int32
	SetBody(v *CreateVccRouteEntryResponseBody) *CreateVccRouteEntryResponse
	GetBody() *CreateVccRouteEntryResponseBody
}

type CreateVccRouteEntryResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVccRouteEntryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVccRouteEntryResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVccRouteEntryResponse) GoString() string {
	return s.String()
}

func (s *CreateVccRouteEntryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVccRouteEntryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVccRouteEntryResponse) GetBody() *CreateVccRouteEntryResponseBody {
	return s.Body
}

func (s *CreateVccRouteEntryResponse) SetHeaders(v map[string]*string) *CreateVccRouteEntryResponse {
	s.Headers = v
	return s
}

func (s *CreateVccRouteEntryResponse) SetStatusCode(v int32) *CreateVccRouteEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVccRouteEntryResponse) SetBody(v *CreateVccRouteEntryResponseBody) *CreateVccRouteEntryResponse {
	s.Body = v
	return s
}

func (s *CreateVccRouteEntryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
