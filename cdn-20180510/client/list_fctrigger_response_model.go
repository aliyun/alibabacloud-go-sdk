// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFCTriggerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFCTriggerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFCTriggerResponse
	GetStatusCode() *int32
	SetBody(v *ListFCTriggerResponseBody) *ListFCTriggerResponse
	GetBody() *ListFCTriggerResponseBody
}

type ListFCTriggerResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFCTriggerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFCTriggerResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFCTriggerResponse) GoString() string {
	return s.String()
}

func (s *ListFCTriggerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFCTriggerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFCTriggerResponse) GetBody() *ListFCTriggerResponseBody {
	return s.Body
}

func (s *ListFCTriggerResponse) SetHeaders(v map[string]*string) *ListFCTriggerResponse {
	s.Headers = v
	return s
}

func (s *ListFCTriggerResponse) SetStatusCode(v int32) *ListFCTriggerResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFCTriggerResponse) SetBody(v *ListFCTriggerResponseBody) *ListFCTriggerResponse {
	s.Body = v
	return s
}

func (s *ListFCTriggerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
