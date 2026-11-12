// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMyAgenticNotificationChannelAuthorizeUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMyAgenticNotificationChannelAuthorizeUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMyAgenticNotificationChannelAuthorizeUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetMyAgenticNotificationChannelAuthorizeUrlResponseBody) *GetMyAgenticNotificationChannelAuthorizeUrlResponse
	GetBody() *GetMyAgenticNotificationChannelAuthorizeUrlResponseBody
}

type GetMyAgenticNotificationChannelAuthorizeUrlResponse struct {
	Headers    map[string]*string                                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMyAgenticNotificationChannelAuthorizeUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMyAgenticNotificationChannelAuthorizeUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMyAgenticNotificationChannelAuthorizeUrlResponse) GoString() string {
	return s.String()
}

func (s *GetMyAgenticNotificationChannelAuthorizeUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMyAgenticNotificationChannelAuthorizeUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMyAgenticNotificationChannelAuthorizeUrlResponse) GetBody() *GetMyAgenticNotificationChannelAuthorizeUrlResponseBody {
	return s.Body
}

func (s *GetMyAgenticNotificationChannelAuthorizeUrlResponse) SetHeaders(v map[string]*string) *GetMyAgenticNotificationChannelAuthorizeUrlResponse {
	s.Headers = v
	return s
}

func (s *GetMyAgenticNotificationChannelAuthorizeUrlResponse) SetStatusCode(v int32) *GetMyAgenticNotificationChannelAuthorizeUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMyAgenticNotificationChannelAuthorizeUrlResponse) SetBody(v *GetMyAgenticNotificationChannelAuthorizeUrlResponseBody) *GetMyAgenticNotificationChannelAuthorizeUrlResponse {
	s.Body = v
	return s
}

func (s *GetMyAgenticNotificationChannelAuthorizeUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
