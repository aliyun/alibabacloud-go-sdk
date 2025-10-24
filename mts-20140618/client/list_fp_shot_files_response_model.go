// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFpShotFilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFpShotFilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFpShotFilesResponse
	GetStatusCode() *int32
	SetBody(v *ListFpShotFilesResponseBody) *ListFpShotFilesResponse
	GetBody() *ListFpShotFilesResponseBody
}

type ListFpShotFilesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFpShotFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFpShotFilesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFpShotFilesResponse) GoString() string {
	return s.String()
}

func (s *ListFpShotFilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFpShotFilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFpShotFilesResponse) GetBody() *ListFpShotFilesResponseBody {
	return s.Body
}

func (s *ListFpShotFilesResponse) SetHeaders(v map[string]*string) *ListFpShotFilesResponse {
	s.Headers = v
	return s
}

func (s *ListFpShotFilesResponse) SetStatusCode(v int32) *ListFpShotFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFpShotFilesResponse) SetBody(v *ListFpShotFilesResponseBody) *ListFpShotFilesResponse {
	s.Body = v
	return s
}

func (s *ListFpShotFilesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
