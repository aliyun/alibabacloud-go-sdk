// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveApplicationAccountFromUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveApplicationAccountFromUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveApplicationAccountFromUserResponse
	GetStatusCode() *int32
	SetBody(v *RemoveApplicationAccountFromUserResponseBody) *RemoveApplicationAccountFromUserResponse
	GetBody() *RemoveApplicationAccountFromUserResponseBody
}

type RemoveApplicationAccountFromUserResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveApplicationAccountFromUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveApplicationAccountFromUserResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveApplicationAccountFromUserResponse) GoString() string {
	return s.String()
}

func (s *RemoveApplicationAccountFromUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveApplicationAccountFromUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveApplicationAccountFromUserResponse) GetBody() *RemoveApplicationAccountFromUserResponseBody {
	return s.Body
}

func (s *RemoveApplicationAccountFromUserResponse) SetHeaders(v map[string]*string) *RemoveApplicationAccountFromUserResponse {
	s.Headers = v
	return s
}

func (s *RemoveApplicationAccountFromUserResponse) SetStatusCode(v int32) *RemoveApplicationAccountFromUserResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveApplicationAccountFromUserResponse) SetBody(v *RemoveApplicationAccountFromUserResponseBody) *RemoveApplicationAccountFromUserResponse {
	s.Body = v
	return s
}

func (s *RemoveApplicationAccountFromUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
