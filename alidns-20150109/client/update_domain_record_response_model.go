// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDomainRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDomainRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDomainRecordResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDomainRecordResponseBody) *UpdateDomainRecordResponse
	GetBody() *UpdateDomainRecordResponseBody
}

type UpdateDomainRecordResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDomainRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDomainRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDomainRecordResponse) GoString() string {
	return s.String()
}

func (s *UpdateDomainRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDomainRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDomainRecordResponse) GetBody() *UpdateDomainRecordResponseBody {
	return s.Body
}

func (s *UpdateDomainRecordResponse) SetHeaders(v map[string]*string) *UpdateDomainRecordResponse {
	s.Headers = v
	return s
}

func (s *UpdateDomainRecordResponse) SetStatusCode(v int32) *UpdateDomainRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDomainRecordResponse) SetBody(v *UpdateDomainRecordResponseBody) *UpdateDomainRecordResponse {
	s.Body = v
	return s
}

func (s *UpdateDomainRecordResponse) Validate() error {
	return dara.Validate(s)
}
