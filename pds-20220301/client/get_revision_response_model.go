// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRevisionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRevisionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRevisionResponse
	GetStatusCode() *int32
	SetBody(v *Revision) *GetRevisionResponse
	GetBody() *Revision
}

type GetRevisionResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Revision          `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRevisionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRevisionResponse) GoString() string {
	return s.String()
}

func (s *GetRevisionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRevisionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRevisionResponse) GetBody() *Revision {
	return s.Body
}

func (s *GetRevisionResponse) SetHeaders(v map[string]*string) *GetRevisionResponse {
	s.Headers = v
	return s
}

func (s *GetRevisionResponse) SetStatusCode(v int32) *GetRevisionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRevisionResponse) SetBody(v *Revision) *GetRevisionResponse {
	s.Body = v
	return s
}

func (s *GetRevisionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
