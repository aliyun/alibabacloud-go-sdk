// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEnterpriseInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEnterpriseInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEnterpriseInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetEnterpriseInfoResponseBody) *GetEnterpriseInfoResponse
	GetBody() *GetEnterpriseInfoResponseBody
}

type GetEnterpriseInfoResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEnterpriseInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEnterpriseInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEnterpriseInfoResponse) GoString() string {
	return s.String()
}

func (s *GetEnterpriseInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEnterpriseInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEnterpriseInfoResponse) GetBody() *GetEnterpriseInfoResponseBody {
	return s.Body
}

func (s *GetEnterpriseInfoResponse) SetHeaders(v map[string]*string) *GetEnterpriseInfoResponse {
	s.Headers = v
	return s
}

func (s *GetEnterpriseInfoResponse) SetStatusCode(v int32) *GetEnterpriseInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEnterpriseInfoResponse) SetBody(v *GetEnterpriseInfoResponseBody) *GetEnterpriseInfoResponse {
	s.Body = v
	return s
}

func (s *GetEnterpriseInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
