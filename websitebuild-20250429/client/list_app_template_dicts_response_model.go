// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppTemplateDictsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAppTemplateDictsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAppTemplateDictsResponse
	GetStatusCode() *int32
	SetBody(v *ListAppTemplateDictsResponseBody) *ListAppTemplateDictsResponse
	GetBody() *ListAppTemplateDictsResponseBody
}

type ListAppTemplateDictsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAppTemplateDictsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAppTemplateDictsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAppTemplateDictsResponse) GoString() string {
	return s.String()
}

func (s *ListAppTemplateDictsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAppTemplateDictsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAppTemplateDictsResponse) GetBody() *ListAppTemplateDictsResponseBody {
	return s.Body
}

func (s *ListAppTemplateDictsResponse) SetHeaders(v map[string]*string) *ListAppTemplateDictsResponse {
	s.Headers = v
	return s
}

func (s *ListAppTemplateDictsResponse) SetStatusCode(v int32) *ListAppTemplateDictsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppTemplateDictsResponse) SetBody(v *ListAppTemplateDictsResponseBody) *ListAppTemplateDictsResponse {
	s.Body = v
	return s
}

func (s *ListAppTemplateDictsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
