// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveApplicationResponse
	GetStatusCode() *int32
	SetBody(v *RemoveApplicationResponseBody) *RemoveApplicationResponse
	GetBody() *RemoveApplicationResponseBody
}

type RemoveApplicationResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveApplicationResponse) GoString() string {
	return s.String()
}

func (s *RemoveApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveApplicationResponse) GetBody() *RemoveApplicationResponseBody {
	return s.Body
}

func (s *RemoveApplicationResponse) SetHeaders(v map[string]*string) *RemoveApplicationResponse {
	s.Headers = v
	return s
}

func (s *RemoveApplicationResponse) SetStatusCode(v int32) *RemoveApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveApplicationResponse) SetBody(v *RemoveApplicationResponseBody) *RemoveApplicationResponse {
	s.Body = v
	return s
}

func (s *RemoveApplicationResponse) Validate() error {
	return dara.Validate(s)
}
