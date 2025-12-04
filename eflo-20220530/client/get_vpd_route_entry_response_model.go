// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVpdRouteEntryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVpdRouteEntryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVpdRouteEntryResponse
	GetStatusCode() *int32
	SetBody(v *GetVpdRouteEntryResponseBody) *GetVpdRouteEntryResponse
	GetBody() *GetVpdRouteEntryResponseBody
}

type GetVpdRouteEntryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVpdRouteEntryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVpdRouteEntryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVpdRouteEntryResponse) GoString() string {
	return s.String()
}

func (s *GetVpdRouteEntryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVpdRouteEntryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVpdRouteEntryResponse) GetBody() *GetVpdRouteEntryResponseBody {
	return s.Body
}

func (s *GetVpdRouteEntryResponse) SetHeaders(v map[string]*string) *GetVpdRouteEntryResponse {
	s.Headers = v
	return s
}

func (s *GetVpdRouteEntryResponse) SetStatusCode(v int32) *GetVpdRouteEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVpdRouteEntryResponse) SetBody(v *GetVpdRouteEntryResponseBody) *GetVpdRouteEntryResponse {
	s.Body = v
	return s
}

func (s *GetVpdRouteEntryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
