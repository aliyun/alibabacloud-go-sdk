// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLivePullToPushResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLivePullToPushResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLivePullToPushResponse
	GetStatusCode() *int32
	SetBody(v *CreateLivePullToPushResponseBody) *CreateLivePullToPushResponse
	GetBody() *CreateLivePullToPushResponseBody
}

type CreateLivePullToPushResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLivePullToPushResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLivePullToPushResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLivePullToPushResponse) GoString() string {
	return s.String()
}

func (s *CreateLivePullToPushResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLivePullToPushResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLivePullToPushResponse) GetBody() *CreateLivePullToPushResponseBody {
	return s.Body
}

func (s *CreateLivePullToPushResponse) SetHeaders(v map[string]*string) *CreateLivePullToPushResponse {
	s.Headers = v
	return s
}

func (s *CreateLivePullToPushResponse) SetStatusCode(v int32) *CreateLivePullToPushResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLivePullToPushResponse) SetBody(v *CreateLivePullToPushResponseBody) *CreateLivePullToPushResponse {
	s.Body = v
	return s
}

func (s *CreateLivePullToPushResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
