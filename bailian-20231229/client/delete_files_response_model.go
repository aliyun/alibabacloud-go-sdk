// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFilesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteFilesResponseBody) *DeleteFilesResponse
	GetBody() *DeleteFilesResponseBody
}

type DeleteFilesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFilesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFilesResponse) GoString() string {
	return s.String()
}

func (s *DeleteFilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFilesResponse) GetBody() *DeleteFilesResponseBody {
	return s.Body
}

func (s *DeleteFilesResponse) SetHeaders(v map[string]*string) *DeleteFilesResponse {
	s.Headers = v
	return s
}

func (s *DeleteFilesResponse) SetStatusCode(v int32) *DeleteFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFilesResponse) SetBody(v *DeleteFilesResponseBody) *DeleteFilesResponse {
	s.Body = v
	return s
}

func (s *DeleteFilesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
