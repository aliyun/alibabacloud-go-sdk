// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrUpdateNotificationPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateOrUpdateNotificationPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateOrUpdateNotificationPolicyResponse
	GetStatusCode() *int32
	SetBody(v *CreateOrUpdateNotificationPolicyResponseBody) *CreateOrUpdateNotificationPolicyResponse
	GetBody() *CreateOrUpdateNotificationPolicyResponseBody
}

type CreateOrUpdateNotificationPolicyResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOrUpdateNotificationPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOrUpdateNotificationPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateNotificationPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateNotificationPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateOrUpdateNotificationPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateOrUpdateNotificationPolicyResponse) GetBody() *CreateOrUpdateNotificationPolicyResponseBody {
	return s.Body
}

func (s *CreateOrUpdateNotificationPolicyResponse) SetHeaders(v map[string]*string) *CreateOrUpdateNotificationPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponse) SetStatusCode(v int32) *CreateOrUpdateNotificationPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponse) SetBody(v *CreateOrUpdateNotificationPolicyResponseBody) *CreateOrUpdateNotificationPolicyResponse {
	s.Body = v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponse) Validate() error {
	return dara.Validate(s)
}
