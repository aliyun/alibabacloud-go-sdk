// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBackupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBackupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBackupsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteBackupsResponseBody) *DeleteBackupsResponse
	GetBody() *DeleteBackupsResponseBody
}

type DeleteBackupsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBackupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBackupsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBackupsResponse) GoString() string {
	return s.String()
}

func (s *DeleteBackupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBackupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBackupsResponse) GetBody() *DeleteBackupsResponseBody {
	return s.Body
}

func (s *DeleteBackupsResponse) SetHeaders(v map[string]*string) *DeleteBackupsResponse {
	s.Headers = v
	return s
}

func (s *DeleteBackupsResponse) SetStatusCode(v int32) *DeleteBackupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBackupsResponse) SetBody(v *DeleteBackupsResponseBody) *DeleteBackupsResponse {
	s.Body = v
	return s
}

func (s *DeleteBackupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
