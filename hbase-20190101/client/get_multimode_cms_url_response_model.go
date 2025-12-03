// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultimodeCmsUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMultimodeCmsUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMultimodeCmsUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetMultimodeCmsUrlResponseBody) *GetMultimodeCmsUrlResponse
	GetBody() *GetMultimodeCmsUrlResponseBody
}

type GetMultimodeCmsUrlResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMultimodeCmsUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMultimodeCmsUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMultimodeCmsUrlResponse) GoString() string {
	return s.String()
}

func (s *GetMultimodeCmsUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMultimodeCmsUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMultimodeCmsUrlResponse) GetBody() *GetMultimodeCmsUrlResponseBody {
	return s.Body
}

func (s *GetMultimodeCmsUrlResponse) SetHeaders(v map[string]*string) *GetMultimodeCmsUrlResponse {
	s.Headers = v
	return s
}

func (s *GetMultimodeCmsUrlResponse) SetStatusCode(v int32) *GetMultimodeCmsUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMultimodeCmsUrlResponse) SetBody(v *GetMultimodeCmsUrlResponseBody) *GetMultimodeCmsUrlResponse {
	s.Body = v
	return s
}

func (s *GetMultimodeCmsUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
