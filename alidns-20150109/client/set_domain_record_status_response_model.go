// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDomainRecordStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetDomainRecordStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetDomainRecordStatusResponse
	GetStatusCode() *int32
	SetBody(v *SetDomainRecordStatusResponseBody) *SetDomainRecordStatusResponse
	GetBody() *SetDomainRecordStatusResponseBody
}

type SetDomainRecordStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDomainRecordStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDomainRecordStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s SetDomainRecordStatusResponse) GoString() string {
	return s.String()
}

func (s *SetDomainRecordStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetDomainRecordStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetDomainRecordStatusResponse) GetBody() *SetDomainRecordStatusResponseBody {
	return s.Body
}

func (s *SetDomainRecordStatusResponse) SetHeaders(v map[string]*string) *SetDomainRecordStatusResponse {
	s.Headers = v
	return s
}

func (s *SetDomainRecordStatusResponse) SetStatusCode(v int32) *SetDomainRecordStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDomainRecordStatusResponse) SetBody(v *SetDomainRecordStatusResponseBody) *SetDomainRecordStatusResponse {
	s.Body = v
	return s
}

func (s *SetDomainRecordStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
