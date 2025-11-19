// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestoreMediaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RestoreMediaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RestoreMediaResponse
	GetStatusCode() *int32
	SetBody(v *RestoreMediaResponseBody) *RestoreMediaResponse
	GetBody() *RestoreMediaResponseBody
}

type RestoreMediaResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestoreMediaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestoreMediaResponse) String() string {
	return dara.Prettify(s)
}

func (s RestoreMediaResponse) GoString() string {
	return s.String()
}

func (s *RestoreMediaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RestoreMediaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RestoreMediaResponse) GetBody() *RestoreMediaResponseBody {
	return s.Body
}

func (s *RestoreMediaResponse) SetHeaders(v map[string]*string) *RestoreMediaResponse {
	s.Headers = v
	return s
}

func (s *RestoreMediaResponse) SetStatusCode(v int32) *RestoreMediaResponse {
	s.StatusCode = &v
	return s
}

func (s *RestoreMediaResponse) SetBody(v *RestoreMediaResponseBody) *RestoreMediaResponse {
	s.Body = v
	return s
}

func (s *RestoreMediaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
