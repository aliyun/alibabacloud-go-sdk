// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantDataServiceApiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GrantDataServiceApiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GrantDataServiceApiResponse
	GetStatusCode() *int32
	SetBody(v *GrantDataServiceApiResponseBody) *GrantDataServiceApiResponse
	GetBody() *GrantDataServiceApiResponseBody
}

type GrantDataServiceApiResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GrantDataServiceApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GrantDataServiceApiResponse) String() string {
	return dara.Prettify(s)
}

func (s GrantDataServiceApiResponse) GoString() string {
	return s.String()
}

func (s *GrantDataServiceApiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GrantDataServiceApiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GrantDataServiceApiResponse) GetBody() *GrantDataServiceApiResponseBody {
	return s.Body
}

func (s *GrantDataServiceApiResponse) SetHeaders(v map[string]*string) *GrantDataServiceApiResponse {
	s.Headers = v
	return s
}

func (s *GrantDataServiceApiResponse) SetStatusCode(v int32) *GrantDataServiceApiResponse {
	s.StatusCode = &v
	return s
}

func (s *GrantDataServiceApiResponse) SetBody(v *GrantDataServiceApiResponseBody) *GrantDataServiceApiResponse {
	s.Body = v
	return s
}

func (s *GrantDataServiceApiResponse) Validate() error {
	return dara.Validate(s)
}
