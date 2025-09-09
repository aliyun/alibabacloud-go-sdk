// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveInstanceAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveInstanceAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveInstanceAccountResponse
	GetStatusCode() *int32
	SetBody(v *RemoveInstanceAccountResponseBody) *RemoveInstanceAccountResponse
	GetBody() *RemoveInstanceAccountResponseBody
}

type RemoveInstanceAccountResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveInstanceAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveInstanceAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveInstanceAccountResponse) GoString() string {
	return s.String()
}

func (s *RemoveInstanceAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveInstanceAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveInstanceAccountResponse) GetBody() *RemoveInstanceAccountResponseBody {
	return s.Body
}

func (s *RemoveInstanceAccountResponse) SetHeaders(v map[string]*string) *RemoveInstanceAccountResponse {
	s.Headers = v
	return s
}

func (s *RemoveInstanceAccountResponse) SetStatusCode(v int32) *RemoveInstanceAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveInstanceAccountResponse) SetBody(v *RemoveInstanceAccountResponseBody) *RemoveInstanceAccountResponse {
	s.Body = v
	return s
}

func (s *RemoveInstanceAccountResponse) Validate() error {
	return dara.Validate(s)
}
