// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySubscriptionPermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySubscriptionPermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySubscriptionPermissionResponse
	GetStatusCode() *int32
	SetBody(v *ModifySubscriptionPermissionResponseBody) *ModifySubscriptionPermissionResponse
	GetBody() *ModifySubscriptionPermissionResponseBody
}

type ModifySubscriptionPermissionResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySubscriptionPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySubscriptionPermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySubscriptionPermissionResponse) GoString() string {
	return s.String()
}

func (s *ModifySubscriptionPermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySubscriptionPermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySubscriptionPermissionResponse) GetBody() *ModifySubscriptionPermissionResponseBody {
	return s.Body
}

func (s *ModifySubscriptionPermissionResponse) SetHeaders(v map[string]*string) *ModifySubscriptionPermissionResponse {
	s.Headers = v
	return s
}

func (s *ModifySubscriptionPermissionResponse) SetStatusCode(v int32) *ModifySubscriptionPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySubscriptionPermissionResponse) SetBody(v *ModifySubscriptionPermissionResponseBody) *ModifySubscriptionPermissionResponse {
	s.Body = v
	return s
}

func (s *ModifySubscriptionPermissionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
