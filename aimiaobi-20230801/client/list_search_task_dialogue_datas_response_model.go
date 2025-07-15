// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSearchTaskDialogueDatasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSearchTaskDialogueDatasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSearchTaskDialogueDatasResponse
	GetStatusCode() *int32
	SetBody(v *ListSearchTaskDialogueDatasResponseBody) *ListSearchTaskDialogueDatasResponse
	GetBody() *ListSearchTaskDialogueDatasResponseBody
}

type ListSearchTaskDialogueDatasResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSearchTaskDialogueDatasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSearchTaskDialogueDatasResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSearchTaskDialogueDatasResponse) GoString() string {
	return s.String()
}

func (s *ListSearchTaskDialogueDatasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSearchTaskDialogueDatasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSearchTaskDialogueDatasResponse) GetBody() *ListSearchTaskDialogueDatasResponseBody {
	return s.Body
}

func (s *ListSearchTaskDialogueDatasResponse) SetHeaders(v map[string]*string) *ListSearchTaskDialogueDatasResponse {
	s.Headers = v
	return s
}

func (s *ListSearchTaskDialogueDatasResponse) SetStatusCode(v int32) *ListSearchTaskDialogueDatasResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSearchTaskDialogueDatasResponse) SetBody(v *ListSearchTaskDialogueDatasResponseBody) *ListSearchTaskDialogueDatasResponse {
	s.Body = v
	return s
}

func (s *ListSearchTaskDialogueDatasResponse) Validate() error {
	return dara.Validate(s)
}
