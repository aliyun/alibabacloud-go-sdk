// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDNAFilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDNAFilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDNAFilesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDNAFilesResponseBody) *DeleteDNAFilesResponse
	GetBody() *DeleteDNAFilesResponseBody
}

type DeleteDNAFilesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDNAFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDNAFilesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDNAFilesResponse) GoString() string {
	return s.String()
}

func (s *DeleteDNAFilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDNAFilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDNAFilesResponse) GetBody() *DeleteDNAFilesResponseBody {
	return s.Body
}

func (s *DeleteDNAFilesResponse) SetHeaders(v map[string]*string) *DeleteDNAFilesResponse {
	s.Headers = v
	return s
}

func (s *DeleteDNAFilesResponse) SetStatusCode(v int32) *DeleteDNAFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDNAFilesResponse) SetBody(v *DeleteDNAFilesResponseBody) *DeleteDNAFilesResponse {
	s.Body = v
	return s
}

func (s *DeleteDNAFilesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
