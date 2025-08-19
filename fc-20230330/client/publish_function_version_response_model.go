// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishFunctionVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PublishFunctionVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PublishFunctionVersionResponse
	GetStatusCode() *int32
	SetBody(v *Version) *PublishFunctionVersionResponse
	GetBody() *Version
}

type PublishFunctionVersionResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Version           `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishFunctionVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s PublishFunctionVersionResponse) GoString() string {
	return s.String()
}

func (s *PublishFunctionVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PublishFunctionVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PublishFunctionVersionResponse) GetBody() *Version {
	return s.Body
}

func (s *PublishFunctionVersionResponse) SetHeaders(v map[string]*string) *PublishFunctionVersionResponse {
	s.Headers = v
	return s
}

func (s *PublishFunctionVersionResponse) SetStatusCode(v int32) *PublishFunctionVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishFunctionVersionResponse) SetBody(v *Version) *PublishFunctionVersionResponse {
	s.Body = v
	return s
}

func (s *PublishFunctionVersionResponse) Validate() error {
	return dara.Validate(s)
}
