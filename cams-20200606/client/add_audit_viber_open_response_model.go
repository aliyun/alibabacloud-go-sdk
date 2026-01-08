// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAuditViberOpenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddAuditViberOpenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddAuditViberOpenResponse
	GetStatusCode() *int32
	SetBody(v *AddAuditViberOpenResponseBody) *AddAuditViberOpenResponse
	GetBody() *AddAuditViberOpenResponseBody
}

type AddAuditViberOpenResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddAuditViberOpenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddAuditViberOpenResponse) String() string {
	return dara.Prettify(s)
}

func (s AddAuditViberOpenResponse) GoString() string {
	return s.String()
}

func (s *AddAuditViberOpenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddAuditViberOpenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddAuditViberOpenResponse) GetBody() *AddAuditViberOpenResponseBody {
	return s.Body
}

func (s *AddAuditViberOpenResponse) SetHeaders(v map[string]*string) *AddAuditViberOpenResponse {
	s.Headers = v
	return s
}

func (s *AddAuditViberOpenResponse) SetStatusCode(v int32) *AddAuditViberOpenResponse {
	s.StatusCode = &v
	return s
}

func (s *AddAuditViberOpenResponse) SetBody(v *AddAuditViberOpenResponseBody) *AddAuditViberOpenResponse {
	s.Body = v
	return s
}

func (s *AddAuditViberOpenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
