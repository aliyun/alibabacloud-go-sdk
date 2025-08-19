// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetControlPolicyEnablementStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetControlPolicyEnablementStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetControlPolicyEnablementStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetControlPolicyEnablementStatusResponseBody) *GetControlPolicyEnablementStatusResponse
	GetBody() *GetControlPolicyEnablementStatusResponseBody
}

type GetControlPolicyEnablementStatusResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetControlPolicyEnablementStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetControlPolicyEnablementStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetControlPolicyEnablementStatusResponse) GoString() string {
	return s.String()
}

func (s *GetControlPolicyEnablementStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetControlPolicyEnablementStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetControlPolicyEnablementStatusResponse) GetBody() *GetControlPolicyEnablementStatusResponseBody {
	return s.Body
}

func (s *GetControlPolicyEnablementStatusResponse) SetHeaders(v map[string]*string) *GetControlPolicyEnablementStatusResponse {
	s.Headers = v
	return s
}

func (s *GetControlPolicyEnablementStatusResponse) SetStatusCode(v int32) *GetControlPolicyEnablementStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetControlPolicyEnablementStatusResponse) SetBody(v *GetControlPolicyEnablementStatusResponseBody) *GetControlPolicyEnablementStatusResponse {
	s.Body = v
	return s
}

func (s *GetControlPolicyEnablementStatusResponse) Validate() error {
	return dara.Validate(s)
}
