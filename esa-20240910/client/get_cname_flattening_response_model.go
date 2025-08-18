// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCnameFlatteningResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCnameFlatteningResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCnameFlatteningResponse
	GetStatusCode() *int32
	SetBody(v *GetCnameFlatteningResponseBody) *GetCnameFlatteningResponse
	GetBody() *GetCnameFlatteningResponseBody
}

type GetCnameFlatteningResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCnameFlatteningResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCnameFlatteningResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCnameFlatteningResponse) GoString() string {
	return s.String()
}

func (s *GetCnameFlatteningResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCnameFlatteningResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCnameFlatteningResponse) GetBody() *GetCnameFlatteningResponseBody {
	return s.Body
}

func (s *GetCnameFlatteningResponse) SetHeaders(v map[string]*string) *GetCnameFlatteningResponse {
	s.Headers = v
	return s
}

func (s *GetCnameFlatteningResponse) SetStatusCode(v int32) *GetCnameFlatteningResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCnameFlatteningResponse) SetBody(v *GetCnameFlatteningResponseBody) *GetCnameFlatteningResponse {
	s.Body = v
	return s
}

func (s *GetCnameFlatteningResponse) Validate() error {
	return dara.Validate(s)
}
