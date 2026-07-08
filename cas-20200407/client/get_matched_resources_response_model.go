// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMatchedResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMatchedResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMatchedResourcesResponse
	GetStatusCode() *int32
	SetBody(v *GetMatchedResourcesResponseBody) *GetMatchedResourcesResponse
	GetBody() *GetMatchedResourcesResponseBody
}

type GetMatchedResourcesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMatchedResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMatchedResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMatchedResourcesResponse) GoString() string {
	return s.String()
}

func (s *GetMatchedResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMatchedResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMatchedResourcesResponse) GetBody() *GetMatchedResourcesResponseBody {
	return s.Body
}

func (s *GetMatchedResourcesResponse) SetHeaders(v map[string]*string) *GetMatchedResourcesResponse {
	s.Headers = v
	return s
}

func (s *GetMatchedResourcesResponse) SetStatusCode(v int32) *GetMatchedResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMatchedResourcesResponse) SetBody(v *GetMatchedResourcesResponseBody) *GetMatchedResourcesResponse {
	s.Body = v
	return s
}

func (s *GetMatchedResourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
