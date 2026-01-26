// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBlockAlarmNotificationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BlockAlarmNotificationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BlockAlarmNotificationResponse
	GetStatusCode() *int32
	SetBody(v *BlockAlarmNotificationResponseBody) *BlockAlarmNotificationResponse
	GetBody() *BlockAlarmNotificationResponseBody
}

type BlockAlarmNotificationResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BlockAlarmNotificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BlockAlarmNotificationResponse) String() string {
	return dara.Prettify(s)
}

func (s BlockAlarmNotificationResponse) GoString() string {
	return s.String()
}

func (s *BlockAlarmNotificationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BlockAlarmNotificationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BlockAlarmNotificationResponse) GetBody() *BlockAlarmNotificationResponseBody {
	return s.Body
}

func (s *BlockAlarmNotificationResponse) SetHeaders(v map[string]*string) *BlockAlarmNotificationResponse {
	s.Headers = v
	return s
}

func (s *BlockAlarmNotificationResponse) SetStatusCode(v int32) *BlockAlarmNotificationResponse {
	s.StatusCode = &v
	return s
}

func (s *BlockAlarmNotificationResponse) SetBody(v *BlockAlarmNotificationResponseBody) *BlockAlarmNotificationResponse {
	s.Body = v
	return s
}

func (s *BlockAlarmNotificationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
