// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPagedInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPagedInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPagedInstanceResponse
	GetStatusCode() *int32
	SetBody(v *GetPagedInstanceResponseBody) *GetPagedInstanceResponse
	GetBody() *GetPagedInstanceResponseBody
}

type GetPagedInstanceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPagedInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPagedInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPagedInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetPagedInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPagedInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPagedInstanceResponse) GetBody() *GetPagedInstanceResponseBody {
	return s.Body
}

func (s *GetPagedInstanceResponse) SetHeaders(v map[string]*string) *GetPagedInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetPagedInstanceResponse) SetStatusCode(v int32) *GetPagedInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPagedInstanceResponse) SetBody(v *GetPagedInstanceResponseBody) *GetPagedInstanceResponse {
	s.Body = v
	return s
}

func (s *GetPagedInstanceResponse) Validate() error {
	return dara.Validate(s)
}
