// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iArchiveFilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ArchiveFilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ArchiveFilesResponse
	GetStatusCode() *int32
	SetBody(v *ArchiveFilesResponseBody) *ArchiveFilesResponse
	GetBody() *ArchiveFilesResponseBody
}

type ArchiveFilesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ArchiveFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ArchiveFilesResponse) String() string {
	return dara.Prettify(s)
}

func (s ArchiveFilesResponse) GoString() string {
	return s.String()
}

func (s *ArchiveFilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ArchiveFilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ArchiveFilesResponse) GetBody() *ArchiveFilesResponseBody {
	return s.Body
}

func (s *ArchiveFilesResponse) SetHeaders(v map[string]*string) *ArchiveFilesResponse {
	s.Headers = v
	return s
}

func (s *ArchiveFilesResponse) SetStatusCode(v int32) *ArchiveFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *ArchiveFilesResponse) SetBody(v *ArchiveFilesResponseBody) *ArchiveFilesResponse {
	s.Body = v
	return s
}

func (s *ArchiveFilesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
