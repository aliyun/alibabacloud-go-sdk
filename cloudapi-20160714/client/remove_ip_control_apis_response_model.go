// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveIpControlApisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveIpControlApisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveIpControlApisResponse
	GetStatusCode() *int32
	SetBody(v *RemoveIpControlApisResponseBody) *RemoveIpControlApisResponse
	GetBody() *RemoveIpControlApisResponseBody
}

type RemoveIpControlApisResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveIpControlApisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveIpControlApisResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveIpControlApisResponse) GoString() string {
	return s.String()
}

func (s *RemoveIpControlApisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveIpControlApisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveIpControlApisResponse) GetBody() *RemoveIpControlApisResponseBody {
	return s.Body
}

func (s *RemoveIpControlApisResponse) SetHeaders(v map[string]*string) *RemoveIpControlApisResponse {
	s.Headers = v
	return s
}

func (s *RemoveIpControlApisResponse) SetStatusCode(v int32) *RemoveIpControlApisResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveIpControlApisResponse) SetBody(v *RemoveIpControlApisResponseBody) *RemoveIpControlApisResponse {
	s.Body = v
	return s
}

func (s *RemoveIpControlApisResponse) Validate() error {
	return dara.Validate(s)
}
