// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlertNotifyTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAlertNotifyTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAlertNotifyTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAlertNotifyTemplatesResponseBody) *DeleteAlertNotifyTemplatesResponse
	GetBody() *DeleteAlertNotifyTemplatesResponseBody
}

type DeleteAlertNotifyTemplatesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAlertNotifyTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAlertNotifyTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlertNotifyTemplatesResponse) GoString() string {
	return s.String()
}

func (s *DeleteAlertNotifyTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAlertNotifyTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAlertNotifyTemplatesResponse) GetBody() *DeleteAlertNotifyTemplatesResponseBody {
	return s.Body
}

func (s *DeleteAlertNotifyTemplatesResponse) SetHeaders(v map[string]*string) *DeleteAlertNotifyTemplatesResponse {
	s.Headers = v
	return s
}

func (s *DeleteAlertNotifyTemplatesResponse) SetStatusCode(v int32) *DeleteAlertNotifyTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAlertNotifyTemplatesResponse) SetBody(v *DeleteAlertNotifyTemplatesResponseBody) *DeleteAlertNotifyTemplatesResponse {
	s.Body = v
	return s
}

func (s *DeleteAlertNotifyTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
