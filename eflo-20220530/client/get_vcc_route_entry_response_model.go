// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVccRouteEntryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVccRouteEntryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVccRouteEntryResponse
	GetStatusCode() *int32
	SetBody(v *GetVccRouteEntryResponseBody) *GetVccRouteEntryResponse
	GetBody() *GetVccRouteEntryResponseBody
}

type GetVccRouteEntryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVccRouteEntryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVccRouteEntryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVccRouteEntryResponse) GoString() string {
	return s.String()
}

func (s *GetVccRouteEntryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVccRouteEntryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVccRouteEntryResponse) GetBody() *GetVccRouteEntryResponseBody {
	return s.Body
}

func (s *GetVccRouteEntryResponse) SetHeaders(v map[string]*string) *GetVccRouteEntryResponse {
	s.Headers = v
	return s
}

func (s *GetVccRouteEntryResponse) SetStatusCode(v int32) *GetVccRouteEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVccRouteEntryResponse) SetBody(v *GetVccRouteEntryResponseBody) *GetVccRouteEntryResponse {
	s.Body = v
	return s
}

func (s *GetVccRouteEntryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
