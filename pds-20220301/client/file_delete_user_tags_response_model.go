// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFileDeleteUserTagsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FileDeleteUserTagsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FileDeleteUserTagsResponse
	GetStatusCode() *int32
}

type FileDeleteUserTagsResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s FileDeleteUserTagsResponse) String() string {
	return dara.Prettify(s)
}

func (s FileDeleteUserTagsResponse) GoString() string {
	return s.String()
}

func (s *FileDeleteUserTagsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FileDeleteUserTagsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FileDeleteUserTagsResponse) SetHeaders(v map[string]*string) *FileDeleteUserTagsResponse {
	s.Headers = v
	return s
}

func (s *FileDeleteUserTagsResponse) SetStatusCode(v int32) *FileDeleteUserTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *FileDeleteUserTagsResponse) Validate() error {
	return dara.Validate(s)
}
