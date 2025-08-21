// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseARMServerInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReleaseARMServerInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReleaseARMServerInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ReleaseARMServerInstanceResponseBody) *ReleaseARMServerInstanceResponse
	GetBody() *ReleaseARMServerInstanceResponseBody
}

type ReleaseARMServerInstanceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseARMServerInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseARMServerInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ReleaseARMServerInstanceResponse) GoString() string {
	return s.String()
}

func (s *ReleaseARMServerInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReleaseARMServerInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReleaseARMServerInstanceResponse) GetBody() *ReleaseARMServerInstanceResponseBody {
	return s.Body
}

func (s *ReleaseARMServerInstanceResponse) SetHeaders(v map[string]*string) *ReleaseARMServerInstanceResponse {
	s.Headers = v
	return s
}

func (s *ReleaseARMServerInstanceResponse) SetStatusCode(v int32) *ReleaseARMServerInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseARMServerInstanceResponse) SetBody(v *ReleaseARMServerInstanceResponseBody) *ReleaseARMServerInstanceResponse {
	s.Body = v
	return s
}

func (s *ReleaseARMServerInstanceResponse) Validate() error {
	return dara.Validate(s)
}
