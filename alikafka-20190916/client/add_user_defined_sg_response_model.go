// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserDefinedSgResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddUserDefinedSgResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddUserDefinedSgResponse
	GetStatusCode() *int32
	SetBody(v *AddUserDefinedSgResponseBody) *AddUserDefinedSgResponse
	GetBody() *AddUserDefinedSgResponseBody
}

type AddUserDefinedSgResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddUserDefinedSgResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddUserDefinedSgResponse) String() string {
	return dara.Prettify(s)
}

func (s AddUserDefinedSgResponse) GoString() string {
	return s.String()
}

func (s *AddUserDefinedSgResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddUserDefinedSgResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddUserDefinedSgResponse) GetBody() *AddUserDefinedSgResponseBody {
	return s.Body
}

func (s *AddUserDefinedSgResponse) SetHeaders(v map[string]*string) *AddUserDefinedSgResponse {
	s.Headers = v
	return s
}

func (s *AddUserDefinedSgResponse) SetStatusCode(v int32) *AddUserDefinedSgResponse {
	s.StatusCode = &v
	return s
}

func (s *AddUserDefinedSgResponse) SetBody(v *AddUserDefinedSgResponseBody) *AddUserDefinedSgResponse {
	s.Body = v
	return s
}

func (s *AddUserDefinedSgResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
