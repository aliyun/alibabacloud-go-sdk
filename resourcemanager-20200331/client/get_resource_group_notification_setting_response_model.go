// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceGroupNotificationSettingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetResourceGroupNotificationSettingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetResourceGroupNotificationSettingResponse
	GetStatusCode() *int32
	SetBody(v *GetResourceGroupNotificationSettingResponseBody) *GetResourceGroupNotificationSettingResponse
	GetBody() *GetResourceGroupNotificationSettingResponseBody
}

type GetResourceGroupNotificationSettingResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceGroupNotificationSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceGroupNotificationSettingResponse) String() string {
	return dara.Prettify(s)
}

func (s GetResourceGroupNotificationSettingResponse) GoString() string {
	return s.String()
}

func (s *GetResourceGroupNotificationSettingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetResourceGroupNotificationSettingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetResourceGroupNotificationSettingResponse) GetBody() *GetResourceGroupNotificationSettingResponseBody {
	return s.Body
}

func (s *GetResourceGroupNotificationSettingResponse) SetHeaders(v map[string]*string) *GetResourceGroupNotificationSettingResponse {
	s.Headers = v
	return s
}

func (s *GetResourceGroupNotificationSettingResponse) SetStatusCode(v int32) *GetResourceGroupNotificationSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceGroupNotificationSettingResponse) SetBody(v *GetResourceGroupNotificationSettingResponseBody) *GetResourceGroupNotificationSettingResponse {
	s.Body = v
	return s
}

func (s *GetResourceGroupNotificationSettingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
