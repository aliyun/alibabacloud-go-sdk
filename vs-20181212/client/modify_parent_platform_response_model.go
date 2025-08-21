// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyParentPlatformResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyParentPlatformResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyParentPlatformResponse
	GetStatusCode() *int32
	SetBody(v *ModifyParentPlatformResponseBody) *ModifyParentPlatformResponse
	GetBody() *ModifyParentPlatformResponseBody
}

type ModifyParentPlatformResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyParentPlatformResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyParentPlatformResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyParentPlatformResponse) GoString() string {
	return s.String()
}

func (s *ModifyParentPlatformResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyParentPlatformResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyParentPlatformResponse) GetBody() *ModifyParentPlatformResponseBody {
	return s.Body
}

func (s *ModifyParentPlatformResponse) SetHeaders(v map[string]*string) *ModifyParentPlatformResponse {
	s.Headers = v
	return s
}

func (s *ModifyParentPlatformResponse) SetStatusCode(v int32) *ModifyParentPlatformResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyParentPlatformResponse) SetBody(v *ModifyParentPlatformResponseBody) *ModifyParentPlatformResponse {
	s.Body = v
	return s
}

func (s *ModifyParentPlatformResponse) Validate() error {
	return dara.Validate(s)
}
