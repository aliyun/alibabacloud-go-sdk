// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSupplementDagrunInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSupplementDagrunInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSupplementDagrunInstanceResponse
	GetStatusCode() *int32
	SetBody(v *GetSupplementDagrunInstanceResponseBody) *GetSupplementDagrunInstanceResponse
	GetBody() *GetSupplementDagrunInstanceResponseBody
}

type GetSupplementDagrunInstanceResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSupplementDagrunInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSupplementDagrunInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSupplementDagrunInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetSupplementDagrunInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSupplementDagrunInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSupplementDagrunInstanceResponse) GetBody() *GetSupplementDagrunInstanceResponseBody {
	return s.Body
}

func (s *GetSupplementDagrunInstanceResponse) SetHeaders(v map[string]*string) *GetSupplementDagrunInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetSupplementDagrunInstanceResponse) SetStatusCode(v int32) *GetSupplementDagrunInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponse) SetBody(v *GetSupplementDagrunInstanceResponseBody) *GetSupplementDagrunInstanceResponse {
	s.Body = v
	return s
}

func (s *GetSupplementDagrunInstanceResponse) Validate() error {
	return dara.Validate(s)
}
