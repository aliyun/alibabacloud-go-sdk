// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNotificationPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteNotificationPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteNotificationPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteNotificationPolicyResponseBody) *DeleteNotificationPolicyResponse
	GetBody() *DeleteNotificationPolicyResponseBody
}

type DeleteNotificationPolicyResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNotificationPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNotificationPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteNotificationPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteNotificationPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteNotificationPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteNotificationPolicyResponse) GetBody() *DeleteNotificationPolicyResponseBody {
	return s.Body
}

func (s *DeleteNotificationPolicyResponse) SetHeaders(v map[string]*string) *DeleteNotificationPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteNotificationPolicyResponse) SetStatusCode(v int32) *DeleteNotificationPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNotificationPolicyResponse) SetBody(v *DeleteNotificationPolicyResponseBody) *DeleteNotificationPolicyResponse {
	s.Body = v
	return s
}

func (s *DeleteNotificationPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
