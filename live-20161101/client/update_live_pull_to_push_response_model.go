// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLivePullToPushResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLivePullToPushResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLivePullToPushResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLivePullToPushResponseBody) *UpdateLivePullToPushResponse
	GetBody() *UpdateLivePullToPushResponseBody
}

type UpdateLivePullToPushResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLivePullToPushResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLivePullToPushResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLivePullToPushResponse) GoString() string {
	return s.String()
}

func (s *UpdateLivePullToPushResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLivePullToPushResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLivePullToPushResponse) GetBody() *UpdateLivePullToPushResponseBody {
	return s.Body
}

func (s *UpdateLivePullToPushResponse) SetHeaders(v map[string]*string) *UpdateLivePullToPushResponse {
	s.Headers = v
	return s
}

func (s *UpdateLivePullToPushResponse) SetStatusCode(v int32) *UpdateLivePullToPushResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLivePullToPushResponse) SetBody(v *UpdateLivePullToPushResponseBody) *UpdateLivePullToPushResponse {
	s.Body = v
	return s
}

func (s *UpdateLivePullToPushResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
