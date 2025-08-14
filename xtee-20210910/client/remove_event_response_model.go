// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveEventResponse
	GetStatusCode() *int32
	SetBody(v *RemoveEventResponseBody) *RemoveEventResponse
	GetBody() *RemoveEventResponseBody
}

type RemoveEventResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveEventResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveEventResponse) GoString() string {
	return s.String()
}

func (s *RemoveEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveEventResponse) GetBody() *RemoveEventResponseBody {
	return s.Body
}

func (s *RemoveEventResponse) SetHeaders(v map[string]*string) *RemoveEventResponse {
	s.Headers = v
	return s
}

func (s *RemoveEventResponse) SetStatusCode(v int32) *RemoveEventResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveEventResponse) SetBody(v *RemoveEventResponseBody) *RemoveEventResponse {
	s.Body = v
	return s
}

func (s *RemoveEventResponse) Validate() error {
	return dara.Validate(s)
}
