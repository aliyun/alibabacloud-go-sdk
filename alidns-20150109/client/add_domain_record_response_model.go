// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDomainRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddDomainRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddDomainRecordResponse
	GetStatusCode() *int32
	SetBody(v *AddDomainRecordResponseBody) *AddDomainRecordResponse
	GetBody() *AddDomainRecordResponseBody
}

type AddDomainRecordResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddDomainRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddDomainRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s AddDomainRecordResponse) GoString() string {
	return s.String()
}

func (s *AddDomainRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddDomainRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddDomainRecordResponse) GetBody() *AddDomainRecordResponseBody {
	return s.Body
}

func (s *AddDomainRecordResponse) SetHeaders(v map[string]*string) *AddDomainRecordResponse {
	s.Headers = v
	return s
}

func (s *AddDomainRecordResponse) SetStatusCode(v int32) *AddDomainRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDomainRecordResponse) SetBody(v *AddDomainRecordResponseBody) *AddDomainRecordResponse {
	s.Body = v
	return s
}

func (s *AddDomainRecordResponse) Validate() error {
	return dara.Validate(s)
}
