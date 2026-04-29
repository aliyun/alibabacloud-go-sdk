// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveApplicationEnvironmentVariablesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveApplicationEnvironmentVariablesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveApplicationEnvironmentVariablesResponse
	GetStatusCode() *int32
	SetBody(v *RemoveApplicationEnvironmentVariablesResponseBody) *RemoveApplicationEnvironmentVariablesResponse
	GetBody() *RemoveApplicationEnvironmentVariablesResponseBody
}

type RemoveApplicationEnvironmentVariablesResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveApplicationEnvironmentVariablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveApplicationEnvironmentVariablesResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveApplicationEnvironmentVariablesResponse) GoString() string {
	return s.String()
}

func (s *RemoveApplicationEnvironmentVariablesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveApplicationEnvironmentVariablesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveApplicationEnvironmentVariablesResponse) GetBody() *RemoveApplicationEnvironmentVariablesResponseBody {
	return s.Body
}

func (s *RemoveApplicationEnvironmentVariablesResponse) SetHeaders(v map[string]*string) *RemoveApplicationEnvironmentVariablesResponse {
	s.Headers = v
	return s
}

func (s *RemoveApplicationEnvironmentVariablesResponse) SetStatusCode(v int32) *RemoveApplicationEnvironmentVariablesResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveApplicationEnvironmentVariablesResponse) SetBody(v *RemoveApplicationEnvironmentVariablesResponseBody) *RemoveApplicationEnvironmentVariablesResponse {
	s.Body = v
	return s
}

func (s *RemoveApplicationEnvironmentVariablesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
