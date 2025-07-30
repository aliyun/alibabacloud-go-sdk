// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveSubInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveSubInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveSubInstanceResponse
	GetStatusCode() *int32
	SetBody(v *RemoveSubInstanceResponseBody) *RemoveSubInstanceResponse
	GetBody() *RemoveSubInstanceResponseBody
}

type RemoveSubInstanceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveSubInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveSubInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveSubInstanceResponse) GoString() string {
	return s.String()
}

func (s *RemoveSubInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveSubInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveSubInstanceResponse) GetBody() *RemoveSubInstanceResponseBody {
	return s.Body
}

func (s *RemoveSubInstanceResponse) SetHeaders(v map[string]*string) *RemoveSubInstanceResponse {
	s.Headers = v
	return s
}

func (s *RemoveSubInstanceResponse) SetStatusCode(v int32) *RemoveSubInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveSubInstanceResponse) SetBody(v *RemoveSubInstanceResponseBody) *RemoveSubInstanceResponse {
	s.Body = v
	return s
}

func (s *RemoveSubInstanceResponse) Validate() error {
	return dara.Validate(s)
}
