// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectNotifyReceiversResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListProjectNotifyReceiversResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListProjectNotifyReceiversResponse
	GetStatusCode() *int32
	SetBody(v *ListProjectNotifyReceiversResponseBody) *ListProjectNotifyReceiversResponse
	GetBody() *ListProjectNotifyReceiversResponseBody
}

type ListProjectNotifyReceiversResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProjectNotifyReceiversResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProjectNotifyReceiversResponse) String() string {
	return dara.Prettify(s)
}

func (s ListProjectNotifyReceiversResponse) GoString() string {
	return s.String()
}

func (s *ListProjectNotifyReceiversResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListProjectNotifyReceiversResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListProjectNotifyReceiversResponse) GetBody() *ListProjectNotifyReceiversResponseBody {
	return s.Body
}

func (s *ListProjectNotifyReceiversResponse) SetHeaders(v map[string]*string) *ListProjectNotifyReceiversResponse {
	s.Headers = v
	return s
}

func (s *ListProjectNotifyReceiversResponse) SetStatusCode(v int32) *ListProjectNotifyReceiversResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProjectNotifyReceiversResponse) SetBody(v *ListProjectNotifyReceiversResponseBody) *ListProjectNotifyReceiversResponse {
	s.Body = v
	return s
}

func (s *ListProjectNotifyReceiversResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
