// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSelectDataSetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SelectDataSetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SelectDataSetResponse
	GetStatusCode() *int32
	SetBody(v *SelectDataSetResponseBody) *SelectDataSetResponse
	GetBody() *SelectDataSetResponseBody
}

type SelectDataSetResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SelectDataSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SelectDataSetResponse) String() string {
	return dara.Prettify(s)
}

func (s SelectDataSetResponse) GoString() string {
	return s.String()
}

func (s *SelectDataSetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SelectDataSetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SelectDataSetResponse) GetBody() *SelectDataSetResponseBody {
	return s.Body
}

func (s *SelectDataSetResponse) SetHeaders(v map[string]*string) *SelectDataSetResponse {
	s.Headers = v
	return s
}

func (s *SelectDataSetResponse) SetStatusCode(v int32) *SelectDataSetResponse {
	s.StatusCode = &v
	return s
}

func (s *SelectDataSetResponse) SetBody(v *SelectDataSetResponseBody) *SelectDataSetResponse {
	s.Body = v
	return s
}

func (s *SelectDataSetResponse) Validate() error {
	return dara.Validate(s)
}
