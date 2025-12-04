// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessKeyLastUsedResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAccessKeyLastUsedResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAccessKeyLastUsedResourcesResponse
	GetStatusCode() *int32
	SetBody(v *GetAccessKeyLastUsedResourcesResponseBody) *GetAccessKeyLastUsedResourcesResponse
	GetBody() *GetAccessKeyLastUsedResourcesResponseBody
}

type GetAccessKeyLastUsedResourcesResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAccessKeyLastUsedResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAccessKeyLastUsedResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAccessKeyLastUsedResourcesResponse) GoString() string {
	return s.String()
}

func (s *GetAccessKeyLastUsedResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAccessKeyLastUsedResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAccessKeyLastUsedResourcesResponse) GetBody() *GetAccessKeyLastUsedResourcesResponseBody {
	return s.Body
}

func (s *GetAccessKeyLastUsedResourcesResponse) SetHeaders(v map[string]*string) *GetAccessKeyLastUsedResourcesResponse {
	s.Headers = v
	return s
}

func (s *GetAccessKeyLastUsedResourcesResponse) SetStatusCode(v int32) *GetAccessKeyLastUsedResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccessKeyLastUsedResourcesResponse) SetBody(v *GetAccessKeyLastUsedResourcesResponseBody) *GetAccessKeyLastUsedResourcesResponse {
	s.Body = v
	return s
}

func (s *GetAccessKeyLastUsedResourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
