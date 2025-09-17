// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetAppResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetAppResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetAppResourcesResponse
	GetStatusCode() *int32
	SetBody(v *ResetAppResourcesResponseBody) *ResetAppResourcesResponse
	GetBody() *ResetAppResourcesResponseBody
}

type ResetAppResourcesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetAppResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetAppResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetAppResourcesResponse) GoString() string {
	return s.String()
}

func (s *ResetAppResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetAppResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetAppResourcesResponse) GetBody() *ResetAppResourcesResponseBody {
	return s.Body
}

func (s *ResetAppResourcesResponse) SetHeaders(v map[string]*string) *ResetAppResourcesResponse {
	s.Headers = v
	return s
}

func (s *ResetAppResourcesResponse) SetStatusCode(v int32) *ResetAppResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetAppResourcesResponse) SetBody(v *ResetAppResourcesResponseBody) *ResetAppResourcesResponse {
	s.Body = v
	return s
}

func (s *ResetAppResourcesResponse) Validate() error {
	return dara.Validate(s)
}
