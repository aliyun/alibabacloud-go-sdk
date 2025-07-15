// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveTrafficControlApisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveTrafficControlApisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveTrafficControlApisResponse
	GetStatusCode() *int32
	SetBody(v *RemoveTrafficControlApisResponseBody) *RemoveTrafficControlApisResponse
	GetBody() *RemoveTrafficControlApisResponseBody
}

type RemoveTrafficControlApisResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveTrafficControlApisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveTrafficControlApisResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveTrafficControlApisResponse) GoString() string {
	return s.String()
}

func (s *RemoveTrafficControlApisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveTrafficControlApisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveTrafficControlApisResponse) GetBody() *RemoveTrafficControlApisResponseBody {
	return s.Body
}

func (s *RemoveTrafficControlApisResponse) SetHeaders(v map[string]*string) *RemoveTrafficControlApisResponse {
	s.Headers = v
	return s
}

func (s *RemoveTrafficControlApisResponse) SetStatusCode(v int32) *RemoveTrafficControlApisResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveTrafficControlApisResponse) SetBody(v *RemoveTrafficControlApisResponseBody) *RemoveTrafficControlApisResponse {
	s.Body = v
	return s
}

func (s *RemoveTrafficControlApisResponse) Validate() error {
	return dara.Validate(s)
}
