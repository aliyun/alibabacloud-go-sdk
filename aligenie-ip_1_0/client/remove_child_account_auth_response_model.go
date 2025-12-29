// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveChildAccountAuthResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveChildAccountAuthResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveChildAccountAuthResponse
	GetStatusCode() *int32
	SetBody(v *RemoveChildAccountAuthResponseBody) *RemoveChildAccountAuthResponse
	GetBody() *RemoveChildAccountAuthResponseBody
}

type RemoveChildAccountAuthResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveChildAccountAuthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveChildAccountAuthResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveChildAccountAuthResponse) GoString() string {
	return s.String()
}

func (s *RemoveChildAccountAuthResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveChildAccountAuthResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveChildAccountAuthResponse) GetBody() *RemoveChildAccountAuthResponseBody {
	return s.Body
}

func (s *RemoveChildAccountAuthResponse) SetHeaders(v map[string]*string) *RemoveChildAccountAuthResponse {
	s.Headers = v
	return s
}

func (s *RemoveChildAccountAuthResponse) SetStatusCode(v int32) *RemoveChildAccountAuthResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveChildAccountAuthResponse) SetBody(v *RemoveChildAccountAuthResponseBody) *RemoveChildAccountAuthResponse {
	s.Body = v
	return s
}

func (s *RemoveChildAccountAuthResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
