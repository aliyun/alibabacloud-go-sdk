// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartAppResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartAppResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartAppResourcesResponse
	GetStatusCode() *int32
	SetBody(v *StartAppResourcesResponseBody) *StartAppResourcesResponse
	GetBody() *StartAppResourcesResponseBody
}

type StartAppResourcesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartAppResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartAppResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s StartAppResourcesResponse) GoString() string {
	return s.String()
}

func (s *StartAppResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartAppResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartAppResourcesResponse) GetBody() *StartAppResourcesResponseBody {
	return s.Body
}

func (s *StartAppResourcesResponse) SetHeaders(v map[string]*string) *StartAppResourcesResponse {
	s.Headers = v
	return s
}

func (s *StartAppResourcesResponse) SetStatusCode(v int32) *StartAppResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *StartAppResourcesResponse) SetBody(v *StartAppResourcesResponseBody) *StartAppResourcesResponse {
	s.Body = v
	return s
}

func (s *StartAppResourcesResponse) Validate() error {
	return dara.Validate(s)
}
