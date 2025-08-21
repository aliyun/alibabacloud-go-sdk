// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveApmResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveApmResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveApmResponse
	GetStatusCode() *int32
	SetBody(v *RemoveApmResponseBody) *RemoveApmResponse
	GetBody() *RemoveApmResponseBody
}

type RemoveApmResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveApmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveApmResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveApmResponse) GoString() string {
	return s.String()
}

func (s *RemoveApmResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveApmResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveApmResponse) GetBody() *RemoveApmResponseBody {
	return s.Body
}

func (s *RemoveApmResponse) SetHeaders(v map[string]*string) *RemoveApmResponse {
	s.Headers = v
	return s
}

func (s *RemoveApmResponse) SetStatusCode(v int32) *RemoveApmResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveApmResponse) SetBody(v *RemoveApmResponseBody) *RemoveApmResponse {
	s.Body = v
	return s
}

func (s *RemoveApmResponse) Validate() error {
	return dara.Validate(s)
}
