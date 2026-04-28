// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestoreRevisionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RestoreRevisionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RestoreRevisionResponse
	GetStatusCode() *int32
	SetBody(v *Revision) *RestoreRevisionResponse
	GetBody() *Revision
}

type RestoreRevisionResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Revision          `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestoreRevisionResponse) String() string {
	return dara.Prettify(s)
}

func (s RestoreRevisionResponse) GoString() string {
	return s.String()
}

func (s *RestoreRevisionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RestoreRevisionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RestoreRevisionResponse) GetBody() *Revision {
	return s.Body
}

func (s *RestoreRevisionResponse) SetHeaders(v map[string]*string) *RestoreRevisionResponse {
	s.Headers = v
	return s
}

func (s *RestoreRevisionResponse) SetStatusCode(v int32) *RestoreRevisionResponse {
	s.StatusCode = &v
	return s
}

func (s *RestoreRevisionResponse) SetBody(v *Revision) *RestoreRevisionResponse {
	s.Body = v
	return s
}

func (s *RestoreRevisionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
