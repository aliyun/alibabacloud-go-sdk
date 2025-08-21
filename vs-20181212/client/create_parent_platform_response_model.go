// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateParentPlatformResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateParentPlatformResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateParentPlatformResponse
	GetStatusCode() *int32
	SetBody(v *CreateParentPlatformResponseBody) *CreateParentPlatformResponse
	GetBody() *CreateParentPlatformResponseBody
}

type CreateParentPlatformResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateParentPlatformResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateParentPlatformResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateParentPlatformResponse) GoString() string {
	return s.String()
}

func (s *CreateParentPlatformResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateParentPlatformResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateParentPlatformResponse) GetBody() *CreateParentPlatformResponseBody {
	return s.Body
}

func (s *CreateParentPlatformResponse) SetHeaders(v map[string]*string) *CreateParentPlatformResponse {
	s.Headers = v
	return s
}

func (s *CreateParentPlatformResponse) SetStatusCode(v int32) *CreateParentPlatformResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateParentPlatformResponse) SetBody(v *CreateParentPlatformResponseBody) *CreateParentPlatformResponse {
	s.Body = v
	return s
}

func (s *CreateParentPlatformResponse) Validate() error {
	return dara.Validate(s)
}
