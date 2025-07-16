// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSpaceDirectoriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSpaceDirectoriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSpaceDirectoriesResponse
	GetStatusCode() *int32
	SetBody(v *GetSpaceDirectoriesResponseBody) *GetSpaceDirectoriesResponse
	GetBody() *GetSpaceDirectoriesResponseBody
}

type GetSpaceDirectoriesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSpaceDirectoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSpaceDirectoriesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSpaceDirectoriesResponse) GoString() string {
	return s.String()
}

func (s *GetSpaceDirectoriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSpaceDirectoriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSpaceDirectoriesResponse) GetBody() *GetSpaceDirectoriesResponseBody {
	return s.Body
}

func (s *GetSpaceDirectoriesResponse) SetHeaders(v map[string]*string) *GetSpaceDirectoriesResponse {
	s.Headers = v
	return s
}

func (s *GetSpaceDirectoriesResponse) SetStatusCode(v int32) *GetSpaceDirectoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSpaceDirectoriesResponse) SetBody(v *GetSpaceDirectoriesResponseBody) *GetSpaceDirectoriesResponse {
	s.Body = v
	return s
}

func (s *GetSpaceDirectoriesResponse) Validate() error {
	return dara.Validate(s)
}
