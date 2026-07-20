// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNotificationClickRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetNotificationClickRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetNotificationClickRecordResponse
	GetStatusCode() *int32
	SetBody(v *GetNotificationClickRecordResponseBody) *GetNotificationClickRecordResponse
	GetBody() *GetNotificationClickRecordResponseBody
}

type GetNotificationClickRecordResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNotificationClickRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNotificationClickRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s GetNotificationClickRecordResponse) GoString() string {
	return s.String()
}

func (s *GetNotificationClickRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetNotificationClickRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetNotificationClickRecordResponse) GetBody() *GetNotificationClickRecordResponseBody {
	return s.Body
}

func (s *GetNotificationClickRecordResponse) SetHeaders(v map[string]*string) *GetNotificationClickRecordResponse {
	s.Headers = v
	return s
}

func (s *GetNotificationClickRecordResponse) SetStatusCode(v int32) *GetNotificationClickRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNotificationClickRecordResponse) SetBody(v *GetNotificationClickRecordResponseBody) *GetNotificationClickRecordResponse {
	s.Body = v
	return s
}

func (s *GetNotificationClickRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
