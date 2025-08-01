// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetResourcesResponse
	GetStatusCode() *int32
	SetBody(v *GetResourcesResponseBody) *GetResourcesResponse
	GetBody() *GetResourcesResponseBody
}

type GetResourcesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetResourcesResponse) GoString() string {
	return s.String()
}

func (s *GetResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetResourcesResponse) GetBody() *GetResourcesResponseBody {
	return s.Body
}

func (s *GetResourcesResponse) SetHeaders(v map[string]*string) *GetResourcesResponse {
	s.Headers = v
	return s
}

func (s *GetResourcesResponse) SetStatusCode(v int32) *GetResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourcesResponse) SetBody(v *GetResourcesResponseBody) *GetResourcesResponse {
	s.Body = v
	return s
}

func (s *GetResourcesResponse) Validate() error {
	return dara.Validate(s)
}
