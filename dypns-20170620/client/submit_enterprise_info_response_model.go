// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitEnterpriseInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitEnterpriseInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitEnterpriseInfoResponse
	GetStatusCode() *int32
	SetBody(v *SubmitEnterpriseInfoResponseBody) *SubmitEnterpriseInfoResponse
	GetBody() *SubmitEnterpriseInfoResponseBody
}

type SubmitEnterpriseInfoResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitEnterpriseInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitEnterpriseInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitEnterpriseInfoResponse) GoString() string {
	return s.String()
}

func (s *SubmitEnterpriseInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitEnterpriseInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitEnterpriseInfoResponse) GetBody() *SubmitEnterpriseInfoResponseBody {
	return s.Body
}

func (s *SubmitEnterpriseInfoResponse) SetHeaders(v map[string]*string) *SubmitEnterpriseInfoResponse {
	s.Headers = v
	return s
}

func (s *SubmitEnterpriseInfoResponse) SetStatusCode(v int32) *SubmitEnterpriseInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitEnterpriseInfoResponse) SetBody(v *SubmitEnterpriseInfoResponseBody) *SubmitEnterpriseInfoResponse {
	s.Body = v
	return s
}

func (s *SubmitEnterpriseInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
