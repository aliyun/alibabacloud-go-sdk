// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSearchTaskDialoguesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSearchTaskDialoguesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSearchTaskDialoguesResponse
	GetStatusCode() *int32
	SetBody(v *ListSearchTaskDialoguesResponseBody) *ListSearchTaskDialoguesResponse
	GetBody() *ListSearchTaskDialoguesResponseBody
}

type ListSearchTaskDialoguesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSearchTaskDialoguesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSearchTaskDialoguesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSearchTaskDialoguesResponse) GoString() string {
	return s.String()
}

func (s *ListSearchTaskDialoguesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSearchTaskDialoguesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSearchTaskDialoguesResponse) GetBody() *ListSearchTaskDialoguesResponseBody {
	return s.Body
}

func (s *ListSearchTaskDialoguesResponse) SetHeaders(v map[string]*string) *ListSearchTaskDialoguesResponse {
	s.Headers = v
	return s
}

func (s *ListSearchTaskDialoguesResponse) SetStatusCode(v int32) *ListSearchTaskDialoguesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSearchTaskDialoguesResponse) SetBody(v *ListSearchTaskDialoguesResponseBody) *ListSearchTaskDialoguesResponse {
	s.Body = v
	return s
}

func (s *ListSearchTaskDialoguesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
