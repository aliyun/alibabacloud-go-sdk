// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApiEndpointsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetApiEndpointsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetApiEndpointsResponse
	GetStatusCode() *int32
	SetBody(v *GetApiEndpointsResponseBody) *GetApiEndpointsResponse
	GetBody() *GetApiEndpointsResponseBody
}

type GetApiEndpointsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApiEndpointsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApiEndpointsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetApiEndpointsResponse) GoString() string {
	return s.String()
}

func (s *GetApiEndpointsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetApiEndpointsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetApiEndpointsResponse) GetBody() *GetApiEndpointsResponseBody {
	return s.Body
}

func (s *GetApiEndpointsResponse) SetHeaders(v map[string]*string) *GetApiEndpointsResponse {
	s.Headers = v
	return s
}

func (s *GetApiEndpointsResponse) SetStatusCode(v int32) *GetApiEndpointsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApiEndpointsResponse) SetBody(v *GetApiEndpointsResponseBody) *GetApiEndpointsResponse {
	s.Body = v
	return s
}

func (s *GetApiEndpointsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
