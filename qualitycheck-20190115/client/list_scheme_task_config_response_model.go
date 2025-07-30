// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSchemeTaskConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSchemeTaskConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSchemeTaskConfigResponse
	GetStatusCode() *int32
	SetBody(v *ListSchemeTaskConfigResponseBody) *ListSchemeTaskConfigResponse
	GetBody() *ListSchemeTaskConfigResponseBody
}

type ListSchemeTaskConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSchemeTaskConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSchemeTaskConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSchemeTaskConfigResponse) GoString() string {
	return s.String()
}

func (s *ListSchemeTaskConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSchemeTaskConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSchemeTaskConfigResponse) GetBody() *ListSchemeTaskConfigResponseBody {
	return s.Body
}

func (s *ListSchemeTaskConfigResponse) SetHeaders(v map[string]*string) *ListSchemeTaskConfigResponse {
	s.Headers = v
	return s
}

func (s *ListSchemeTaskConfigResponse) SetStatusCode(v int32) *ListSchemeTaskConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSchemeTaskConfigResponse) SetBody(v *ListSchemeTaskConfigResponseBody) *ListSchemeTaskConfigResponse {
	s.Body = v
	return s
}

func (s *ListSchemeTaskConfigResponse) Validate() error {
	return dara.Validate(s)
}
