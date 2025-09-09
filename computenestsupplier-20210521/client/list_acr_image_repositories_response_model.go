// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAcrImageRepositoriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAcrImageRepositoriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAcrImageRepositoriesResponse
	GetStatusCode() *int32
	SetBody(v *ListAcrImageRepositoriesResponseBody) *ListAcrImageRepositoriesResponse
	GetBody() *ListAcrImageRepositoriesResponseBody
}

type ListAcrImageRepositoriesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAcrImageRepositoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAcrImageRepositoriesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAcrImageRepositoriesResponse) GoString() string {
	return s.String()
}

func (s *ListAcrImageRepositoriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAcrImageRepositoriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAcrImageRepositoriesResponse) GetBody() *ListAcrImageRepositoriesResponseBody {
	return s.Body
}

func (s *ListAcrImageRepositoriesResponse) SetHeaders(v map[string]*string) *ListAcrImageRepositoriesResponse {
	s.Headers = v
	return s
}

func (s *ListAcrImageRepositoriesResponse) SetStatusCode(v int32) *ListAcrImageRepositoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAcrImageRepositoriesResponse) SetBody(v *ListAcrImageRepositoriesResponseBody) *ListAcrImageRepositoriesResponse {
	s.Body = v
	return s
}

func (s *ListAcrImageRepositoriesResponse) Validate() error {
	return dara.Validate(s)
}
