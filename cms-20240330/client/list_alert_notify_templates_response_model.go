// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertNotifyTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAlertNotifyTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAlertNotifyTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *ListAlertNotifyTemplatesResponseBody) *ListAlertNotifyTemplatesResponse
	GetBody() *ListAlertNotifyTemplatesResponseBody
}

type ListAlertNotifyTemplatesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAlertNotifyTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAlertNotifyTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAlertNotifyTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListAlertNotifyTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAlertNotifyTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAlertNotifyTemplatesResponse) GetBody() *ListAlertNotifyTemplatesResponseBody {
	return s.Body
}

func (s *ListAlertNotifyTemplatesResponse) SetHeaders(v map[string]*string) *ListAlertNotifyTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListAlertNotifyTemplatesResponse) SetStatusCode(v int32) *ListAlertNotifyTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAlertNotifyTemplatesResponse) SetBody(v *ListAlertNotifyTemplatesResponseBody) *ListAlertNotifyTemplatesResponse {
	s.Body = v
	return s
}

func (s *ListAlertNotifyTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
