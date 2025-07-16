// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCdnMigrateRegisterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CdnMigrateRegisterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CdnMigrateRegisterResponse
	GetStatusCode() *int32
	SetBody(v *CdnMigrateRegisterResponseBody) *CdnMigrateRegisterResponse
	GetBody() *CdnMigrateRegisterResponseBody
}

type CdnMigrateRegisterResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CdnMigrateRegisterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CdnMigrateRegisterResponse) String() string {
	return dara.Prettify(s)
}

func (s CdnMigrateRegisterResponse) GoString() string {
	return s.String()
}

func (s *CdnMigrateRegisterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CdnMigrateRegisterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CdnMigrateRegisterResponse) GetBody() *CdnMigrateRegisterResponseBody {
	return s.Body
}

func (s *CdnMigrateRegisterResponse) SetHeaders(v map[string]*string) *CdnMigrateRegisterResponse {
	s.Headers = v
	return s
}

func (s *CdnMigrateRegisterResponse) SetStatusCode(v int32) *CdnMigrateRegisterResponse {
	s.StatusCode = &v
	return s
}

func (s *CdnMigrateRegisterResponse) SetBody(v *CdnMigrateRegisterResponseBody) *CdnMigrateRegisterResponse {
	s.Body = v
	return s
}

func (s *CdnMigrateRegisterResponse) Validate() error {
	return dara.Validate(s)
}
