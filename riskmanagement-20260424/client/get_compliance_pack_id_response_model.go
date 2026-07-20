// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCompliancePackIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCompliancePackIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCompliancePackIdResponse
	GetStatusCode() *int32
	SetBody(v *GetCompliancePackIdResponseBody) *GetCompliancePackIdResponse
	GetBody() *GetCompliancePackIdResponseBody
}

type GetCompliancePackIdResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCompliancePackIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCompliancePackIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCompliancePackIdResponse) GoString() string {
	return s.String()
}

func (s *GetCompliancePackIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCompliancePackIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCompliancePackIdResponse) GetBody() *GetCompliancePackIdResponseBody {
	return s.Body
}

func (s *GetCompliancePackIdResponse) SetHeaders(v map[string]*string) *GetCompliancePackIdResponse {
	s.Headers = v
	return s
}

func (s *GetCompliancePackIdResponse) SetStatusCode(v int32) *GetCompliancePackIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCompliancePackIdResponse) SetBody(v *GetCompliancePackIdResponseBody) *GetCompliancePackIdResponse {
	s.Body = v
	return s
}

func (s *GetCompliancePackIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
