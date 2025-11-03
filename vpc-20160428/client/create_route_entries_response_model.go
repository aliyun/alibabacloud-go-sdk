// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRouteEntriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRouteEntriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRouteEntriesResponse
	GetStatusCode() *int32
	SetBody(v *CreateRouteEntriesResponseBody) *CreateRouteEntriesResponse
	GetBody() *CreateRouteEntriesResponseBody
}

type CreateRouteEntriesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRouteEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRouteEntriesResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRouteEntriesResponse) GoString() string {
	return s.String()
}

func (s *CreateRouteEntriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRouteEntriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRouteEntriesResponse) GetBody() *CreateRouteEntriesResponseBody {
	return s.Body
}

func (s *CreateRouteEntriesResponse) SetHeaders(v map[string]*string) *CreateRouteEntriesResponse {
	s.Headers = v
	return s
}

func (s *CreateRouteEntriesResponse) SetStatusCode(v int32) *CreateRouteEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRouteEntriesResponse) SetBody(v *CreateRouteEntriesResponseBody) *CreateRouteEntriesResponse {
	s.Body = v
	return s
}

func (s *CreateRouteEntriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
