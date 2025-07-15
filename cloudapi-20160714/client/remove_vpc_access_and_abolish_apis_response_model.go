// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveVpcAccessAndAbolishApisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveVpcAccessAndAbolishApisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveVpcAccessAndAbolishApisResponse
	GetStatusCode() *int32
	SetBody(v *RemoveVpcAccessAndAbolishApisResponseBody) *RemoveVpcAccessAndAbolishApisResponse
	GetBody() *RemoveVpcAccessAndAbolishApisResponseBody
}

type RemoveVpcAccessAndAbolishApisResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveVpcAccessAndAbolishApisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveVpcAccessAndAbolishApisResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveVpcAccessAndAbolishApisResponse) GoString() string {
	return s.String()
}

func (s *RemoveVpcAccessAndAbolishApisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveVpcAccessAndAbolishApisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveVpcAccessAndAbolishApisResponse) GetBody() *RemoveVpcAccessAndAbolishApisResponseBody {
	return s.Body
}

func (s *RemoveVpcAccessAndAbolishApisResponse) SetHeaders(v map[string]*string) *RemoveVpcAccessAndAbolishApisResponse {
	s.Headers = v
	return s
}

func (s *RemoveVpcAccessAndAbolishApisResponse) SetStatusCode(v int32) *RemoveVpcAccessAndAbolishApisResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveVpcAccessAndAbolishApisResponse) SetBody(v *RemoveVpcAccessAndAbolishApisResponseBody) *RemoveVpcAccessAndAbolishApisResponse {
	s.Body = v
	return s
}

func (s *RemoveVpcAccessAndAbolishApisResponse) Validate() error {
	return dara.Validate(s)
}
