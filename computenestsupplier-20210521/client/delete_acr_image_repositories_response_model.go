// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAcrImageRepositoriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAcrImageRepositoriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAcrImageRepositoriesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAcrImageRepositoriesResponseBody) *DeleteAcrImageRepositoriesResponse
	GetBody() *DeleteAcrImageRepositoriesResponseBody
}

type DeleteAcrImageRepositoriesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAcrImageRepositoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAcrImageRepositoriesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAcrImageRepositoriesResponse) GoString() string {
	return s.String()
}

func (s *DeleteAcrImageRepositoriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAcrImageRepositoriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAcrImageRepositoriesResponse) GetBody() *DeleteAcrImageRepositoriesResponseBody {
	return s.Body
}

func (s *DeleteAcrImageRepositoriesResponse) SetHeaders(v map[string]*string) *DeleteAcrImageRepositoriesResponse {
	s.Headers = v
	return s
}

func (s *DeleteAcrImageRepositoriesResponse) SetStatusCode(v int32) *DeleteAcrImageRepositoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAcrImageRepositoriesResponse) SetBody(v *DeleteAcrImageRepositoriesResponseBody) *DeleteAcrImageRepositoriesResponse {
	s.Body = v
	return s
}

func (s *DeleteAcrImageRepositoriesResponse) Validate() error {
	return dara.Validate(s)
}
