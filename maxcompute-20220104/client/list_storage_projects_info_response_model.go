// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStorageProjectsInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListStorageProjectsInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListStorageProjectsInfoResponse
	GetStatusCode() *int32
	SetBody(v *ListStorageProjectsInfoResponseBody) *ListStorageProjectsInfoResponse
	GetBody() *ListStorageProjectsInfoResponseBody
}

type ListStorageProjectsInfoResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListStorageProjectsInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListStorageProjectsInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ListStorageProjectsInfoResponse) GoString() string {
	return s.String()
}

func (s *ListStorageProjectsInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListStorageProjectsInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListStorageProjectsInfoResponse) GetBody() *ListStorageProjectsInfoResponseBody {
	return s.Body
}

func (s *ListStorageProjectsInfoResponse) SetHeaders(v map[string]*string) *ListStorageProjectsInfoResponse {
	s.Headers = v
	return s
}

func (s *ListStorageProjectsInfoResponse) SetStatusCode(v int32) *ListStorageProjectsInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ListStorageProjectsInfoResponse) SetBody(v *ListStorageProjectsInfoResponseBody) *ListStorageProjectsInfoResponse {
	s.Body = v
	return s
}

func (s *ListStorageProjectsInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
