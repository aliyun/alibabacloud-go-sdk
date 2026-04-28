// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveFaceGroupFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveFaceGroupFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveFaceGroupFileResponse
	GetStatusCode() *int32
}

type RemoveFaceGroupFileResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s RemoveFaceGroupFileResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveFaceGroupFileResponse) GoString() string {
	return s.String()
}

func (s *RemoveFaceGroupFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveFaceGroupFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveFaceGroupFileResponse) SetHeaders(v map[string]*string) *RemoveFaceGroupFileResponse {
	s.Headers = v
	return s
}

func (s *RemoveFaceGroupFileResponse) SetStatusCode(v int32) *RemoveFaceGroupFileResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveFaceGroupFileResponse) Validate() error {
	return dara.Validate(s)
}
