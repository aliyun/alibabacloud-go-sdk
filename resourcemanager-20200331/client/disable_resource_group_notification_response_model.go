// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableResourceGroupNotificationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableResourceGroupNotificationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableResourceGroupNotificationResponse
	GetStatusCode() *int32
	SetBody(v *DisableResourceGroupNotificationResponseBody) *DisableResourceGroupNotificationResponse
	GetBody() *DisableResourceGroupNotificationResponseBody
}

type DisableResourceGroupNotificationResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableResourceGroupNotificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableResourceGroupNotificationResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableResourceGroupNotificationResponse) GoString() string {
	return s.String()
}

func (s *DisableResourceGroupNotificationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableResourceGroupNotificationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableResourceGroupNotificationResponse) GetBody() *DisableResourceGroupNotificationResponseBody {
	return s.Body
}

func (s *DisableResourceGroupNotificationResponse) SetHeaders(v map[string]*string) *DisableResourceGroupNotificationResponse {
	s.Headers = v
	return s
}

func (s *DisableResourceGroupNotificationResponse) SetStatusCode(v int32) *DisableResourceGroupNotificationResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableResourceGroupNotificationResponse) SetBody(v *DisableResourceGroupNotificationResponseBody) *DisableResourceGroupNotificationResponse {
	s.Body = v
	return s
}

func (s *DisableResourceGroupNotificationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
