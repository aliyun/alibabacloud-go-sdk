// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeConfirmResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChangeConfirmResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChangeConfirmResponse
	GetStatusCode() *int32
	SetBody(v *ChangeConfirmResponseBody) *ChangeConfirmResponse
	GetBody() *ChangeConfirmResponseBody
}

type ChangeConfirmResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeConfirmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeConfirmResponse) String() string {
	return dara.Prettify(s)
}

func (s ChangeConfirmResponse) GoString() string {
	return s.String()
}

func (s *ChangeConfirmResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChangeConfirmResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChangeConfirmResponse) GetBody() *ChangeConfirmResponseBody {
	return s.Body
}

func (s *ChangeConfirmResponse) SetHeaders(v map[string]*string) *ChangeConfirmResponse {
	s.Headers = v
	return s
}

func (s *ChangeConfirmResponse) SetStatusCode(v int32) *ChangeConfirmResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeConfirmResponse) SetBody(v *ChangeConfirmResponseBody) *ChangeConfirmResponse {
	s.Body = v
	return s
}

func (s *ChangeConfirmResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
