// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemovePersonalNumbersFromUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemovePersonalNumbersFromUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemovePersonalNumbersFromUserResponse
	GetStatusCode() *int32
	SetBody(v *RemovePersonalNumbersFromUserResponseBody) *RemovePersonalNumbersFromUserResponse
	GetBody() *RemovePersonalNumbersFromUserResponseBody
}

type RemovePersonalNumbersFromUserResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemovePersonalNumbersFromUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemovePersonalNumbersFromUserResponse) String() string {
	return dara.Prettify(s)
}

func (s RemovePersonalNumbersFromUserResponse) GoString() string {
	return s.String()
}

func (s *RemovePersonalNumbersFromUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemovePersonalNumbersFromUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemovePersonalNumbersFromUserResponse) GetBody() *RemovePersonalNumbersFromUserResponseBody {
	return s.Body
}

func (s *RemovePersonalNumbersFromUserResponse) SetHeaders(v map[string]*string) *RemovePersonalNumbersFromUserResponse {
	s.Headers = v
	return s
}

func (s *RemovePersonalNumbersFromUserResponse) SetStatusCode(v int32) *RemovePersonalNumbersFromUserResponse {
	s.StatusCode = &v
	return s
}

func (s *RemovePersonalNumbersFromUserResponse) SetBody(v *RemovePersonalNumbersFromUserResponseBody) *RemovePersonalNumbersFromUserResponse {
	s.Body = v
	return s
}

func (s *RemovePersonalNumbersFromUserResponse) Validate() error {
	return dara.Validate(s)
}
