// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDomainBackupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddDomainBackupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddDomainBackupResponse
	GetStatusCode() *int32
	SetBody(v *AddDomainBackupResponseBody) *AddDomainBackupResponse
	GetBody() *AddDomainBackupResponseBody
}

type AddDomainBackupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddDomainBackupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddDomainBackupResponse) String() string {
	return dara.Prettify(s)
}

func (s AddDomainBackupResponse) GoString() string {
	return s.String()
}

func (s *AddDomainBackupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddDomainBackupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddDomainBackupResponse) GetBody() *AddDomainBackupResponseBody {
	return s.Body
}

func (s *AddDomainBackupResponse) SetHeaders(v map[string]*string) *AddDomainBackupResponse {
	s.Headers = v
	return s
}

func (s *AddDomainBackupResponse) SetStatusCode(v int32) *AddDomainBackupResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDomainBackupResponse) SetBody(v *AddDomainBackupResponseBody) *AddDomainBackupResponse {
	s.Body = v
	return s
}

func (s *AddDomainBackupResponse) Validate() error {
	return dara.Validate(s)
}
