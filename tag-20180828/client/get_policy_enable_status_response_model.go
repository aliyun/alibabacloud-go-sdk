// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPolicyEnableStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPolicyEnableStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPolicyEnableStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetPolicyEnableStatusResponseBody) *GetPolicyEnableStatusResponse
	GetBody() *GetPolicyEnableStatusResponseBody
}

type GetPolicyEnableStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPolicyEnableStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPolicyEnableStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyEnableStatusResponse) GoString() string {
	return s.String()
}

func (s *GetPolicyEnableStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPolicyEnableStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPolicyEnableStatusResponse) GetBody() *GetPolicyEnableStatusResponseBody {
	return s.Body
}

func (s *GetPolicyEnableStatusResponse) SetHeaders(v map[string]*string) *GetPolicyEnableStatusResponse {
	s.Headers = v
	return s
}

func (s *GetPolicyEnableStatusResponse) SetStatusCode(v int32) *GetPolicyEnableStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPolicyEnableStatusResponse) SetBody(v *GetPolicyEnableStatusResponseBody) *GetPolicyEnableStatusResponse {
	s.Body = v
	return s
}

func (s *GetPolicyEnableStatusResponse) Validate() error {
	return dara.Validate(s)
}
