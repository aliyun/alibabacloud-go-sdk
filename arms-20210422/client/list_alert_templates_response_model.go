// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAlertTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAlertTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *ListAlertTemplatesResponseBody) *ListAlertTemplatesResponse
	GetBody() *ListAlertTemplatesResponseBody
}

type ListAlertTemplatesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAlertTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAlertTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAlertTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListAlertTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAlertTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAlertTemplatesResponse) GetBody() *ListAlertTemplatesResponseBody {
	return s.Body
}

func (s *ListAlertTemplatesResponse) SetHeaders(v map[string]*string) *ListAlertTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListAlertTemplatesResponse) SetStatusCode(v int32) *ListAlertTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAlertTemplatesResponse) SetBody(v *ListAlertTemplatesResponseBody) *ListAlertTemplatesResponse {
	s.Body = v
	return s
}

func (s *ListAlertTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
