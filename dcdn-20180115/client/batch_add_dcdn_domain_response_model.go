// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchAddDcdnDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchAddDcdnDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchAddDcdnDomainResponse
	GetStatusCode() *int32
	SetBody(v *BatchAddDcdnDomainResponseBody) *BatchAddDcdnDomainResponse
	GetBody() *BatchAddDcdnDomainResponseBody
}

type BatchAddDcdnDomainResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchAddDcdnDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchAddDcdnDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchAddDcdnDomainResponse) GoString() string {
	return s.String()
}

func (s *BatchAddDcdnDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchAddDcdnDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchAddDcdnDomainResponse) GetBody() *BatchAddDcdnDomainResponseBody {
	return s.Body
}

func (s *BatchAddDcdnDomainResponse) SetHeaders(v map[string]*string) *BatchAddDcdnDomainResponse {
	s.Headers = v
	return s
}

func (s *BatchAddDcdnDomainResponse) SetStatusCode(v int32) *BatchAddDcdnDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchAddDcdnDomainResponse) SetBody(v *BatchAddDcdnDomainResponseBody) *BatchAddDcdnDomainResponse {
	s.Body = v
	return s
}

func (s *BatchAddDcdnDomainResponse) Validate() error {
	return dara.Validate(s)
}
