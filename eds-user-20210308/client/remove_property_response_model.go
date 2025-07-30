// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemovePropertyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemovePropertyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemovePropertyResponse
	GetStatusCode() *int32
	SetBody(v *RemovePropertyResponseBody) *RemovePropertyResponse
	GetBody() *RemovePropertyResponseBody
}

type RemovePropertyResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemovePropertyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemovePropertyResponse) String() string {
	return dara.Prettify(s)
}

func (s RemovePropertyResponse) GoString() string {
	return s.String()
}

func (s *RemovePropertyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemovePropertyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemovePropertyResponse) GetBody() *RemovePropertyResponseBody {
	return s.Body
}

func (s *RemovePropertyResponse) SetHeaders(v map[string]*string) *RemovePropertyResponse {
	s.Headers = v
	return s
}

func (s *RemovePropertyResponse) SetStatusCode(v int32) *RemovePropertyResponse {
	s.StatusCode = &v
	return s
}

func (s *RemovePropertyResponse) SetBody(v *RemovePropertyResponseBody) *RemovePropertyResponse {
	s.Body = v
	return s
}

func (s *RemovePropertyResponse) Validate() error {
	return dara.Validate(s)
}
