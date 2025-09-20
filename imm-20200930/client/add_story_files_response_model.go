// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddStoryFilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddStoryFilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddStoryFilesResponse
	GetStatusCode() *int32
	SetBody(v *AddStoryFilesResponseBody) *AddStoryFilesResponse
	GetBody() *AddStoryFilesResponseBody
}

type AddStoryFilesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddStoryFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddStoryFilesResponse) String() string {
	return dara.Prettify(s)
}

func (s AddStoryFilesResponse) GoString() string {
	return s.String()
}

func (s *AddStoryFilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddStoryFilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddStoryFilesResponse) GetBody() *AddStoryFilesResponseBody {
	return s.Body
}

func (s *AddStoryFilesResponse) SetHeaders(v map[string]*string) *AddStoryFilesResponse {
	s.Headers = v
	return s
}

func (s *AddStoryFilesResponse) SetStatusCode(v int32) *AddStoryFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *AddStoryFilesResponse) SetBody(v *AddStoryFilesResponseBody) *AddStoryFilesResponse {
	s.Body = v
	return s
}

func (s *AddStoryFilesResponse) Validate() error {
	return dara.Validate(s)
}
