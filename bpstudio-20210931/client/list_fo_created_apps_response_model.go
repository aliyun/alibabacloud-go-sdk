// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFoCreatedAppsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFoCreatedAppsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFoCreatedAppsResponse
	GetStatusCode() *int32
	SetBody(v *ListFoCreatedAppsResponseBody) *ListFoCreatedAppsResponse
	GetBody() *ListFoCreatedAppsResponseBody
}

type ListFoCreatedAppsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFoCreatedAppsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFoCreatedAppsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFoCreatedAppsResponse) GoString() string {
	return s.String()
}

func (s *ListFoCreatedAppsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFoCreatedAppsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFoCreatedAppsResponse) GetBody() *ListFoCreatedAppsResponseBody {
	return s.Body
}

func (s *ListFoCreatedAppsResponse) SetHeaders(v map[string]*string) *ListFoCreatedAppsResponse {
	s.Headers = v
	return s
}

func (s *ListFoCreatedAppsResponse) SetStatusCode(v int32) *ListFoCreatedAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFoCreatedAppsResponse) SetBody(v *ListFoCreatedAppsResponseBody) *ListFoCreatedAppsResponse {
	s.Body = v
	return s
}

func (s *ListFoCreatedAppsResponse) Validate() error {
	return dara.Validate(s)
}
