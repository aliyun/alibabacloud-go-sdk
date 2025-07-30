// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDtsJobPasswordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDtsJobPasswordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDtsJobPasswordResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDtsJobPasswordResponseBody) *ModifyDtsJobPasswordResponse
	GetBody() *ModifyDtsJobPasswordResponseBody
}

type ModifyDtsJobPasswordResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDtsJobPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDtsJobPasswordResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDtsJobPasswordResponse) GoString() string {
	return s.String()
}

func (s *ModifyDtsJobPasswordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDtsJobPasswordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDtsJobPasswordResponse) GetBody() *ModifyDtsJobPasswordResponseBody {
	return s.Body
}

func (s *ModifyDtsJobPasswordResponse) SetHeaders(v map[string]*string) *ModifyDtsJobPasswordResponse {
	s.Headers = v
	return s
}

func (s *ModifyDtsJobPasswordResponse) SetStatusCode(v int32) *ModifyDtsJobPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDtsJobPasswordResponse) SetBody(v *ModifyDtsJobPasswordResponseBody) *ModifyDtsJobPasswordResponse {
	s.Body = v
	return s
}

func (s *ModifyDtsJobPasswordResponse) Validate() error {
	return dara.Validate(s)
}
