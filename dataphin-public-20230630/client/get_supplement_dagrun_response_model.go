// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSupplementDagrunResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSupplementDagrunResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSupplementDagrunResponse
	GetStatusCode() *int32
	SetBody(v *GetSupplementDagrunResponseBody) *GetSupplementDagrunResponse
	GetBody() *GetSupplementDagrunResponseBody
}

type GetSupplementDagrunResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSupplementDagrunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSupplementDagrunResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSupplementDagrunResponse) GoString() string {
	return s.String()
}

func (s *GetSupplementDagrunResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSupplementDagrunResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSupplementDagrunResponse) GetBody() *GetSupplementDagrunResponseBody {
	return s.Body
}

func (s *GetSupplementDagrunResponse) SetHeaders(v map[string]*string) *GetSupplementDagrunResponse {
	s.Headers = v
	return s
}

func (s *GetSupplementDagrunResponse) SetStatusCode(v int32) *GetSupplementDagrunResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSupplementDagrunResponse) SetBody(v *GetSupplementDagrunResponseBody) *GetSupplementDagrunResponse {
	s.Body = v
	return s
}

func (s *GetSupplementDagrunResponse) Validate() error {
	return dara.Validate(s)
}
