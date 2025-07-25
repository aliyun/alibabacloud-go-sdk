// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDomainRecordRemarkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDomainRecordRemarkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDomainRecordRemarkResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDomainRecordRemarkResponseBody) *UpdateDomainRecordRemarkResponse
	GetBody() *UpdateDomainRecordRemarkResponseBody
}

type UpdateDomainRecordRemarkResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDomainRecordRemarkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDomainRecordRemarkResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDomainRecordRemarkResponse) GoString() string {
	return s.String()
}

func (s *UpdateDomainRecordRemarkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDomainRecordRemarkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDomainRecordRemarkResponse) GetBody() *UpdateDomainRecordRemarkResponseBody {
	return s.Body
}

func (s *UpdateDomainRecordRemarkResponse) SetHeaders(v map[string]*string) *UpdateDomainRecordRemarkResponse {
	s.Headers = v
	return s
}

func (s *UpdateDomainRecordRemarkResponse) SetStatusCode(v int32) *UpdateDomainRecordRemarkResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDomainRecordRemarkResponse) SetBody(v *UpdateDomainRecordRemarkResponseBody) *UpdateDomainRecordRemarkResponse {
	s.Body = v
	return s
}

func (s *UpdateDomainRecordRemarkResponse) Validate() error {
	return dara.Validate(s)
}
