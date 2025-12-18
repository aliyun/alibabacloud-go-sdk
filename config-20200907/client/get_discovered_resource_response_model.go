// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDiscoveredResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDiscoveredResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDiscoveredResourceResponse
	GetStatusCode() *int32
	SetBody(v *GetDiscoveredResourceResponseBody) *GetDiscoveredResourceResponse
	GetBody() *GetDiscoveredResourceResponseBody
}

type GetDiscoveredResourceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDiscoveredResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDiscoveredResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDiscoveredResourceResponse) GoString() string {
	return s.String()
}

func (s *GetDiscoveredResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDiscoveredResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDiscoveredResourceResponse) GetBody() *GetDiscoveredResourceResponseBody {
	return s.Body
}

func (s *GetDiscoveredResourceResponse) SetHeaders(v map[string]*string) *GetDiscoveredResourceResponse {
	s.Headers = v
	return s
}

func (s *GetDiscoveredResourceResponse) SetStatusCode(v int32) *GetDiscoveredResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDiscoveredResourceResponse) SetBody(v *GetDiscoveredResourceResponseBody) *GetDiscoveredResourceResponse {
	s.Body = v
	return s
}

func (s *GetDiscoveredResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
