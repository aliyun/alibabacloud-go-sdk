// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataDomainInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataDomainInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataDomainInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetDataDomainInfoResponseBody) *GetDataDomainInfoResponse
	GetBody() *GetDataDomainInfoResponseBody
}

type GetDataDomainInfoResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataDomainInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataDomainInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataDomainInfoResponse) GoString() string {
	return s.String()
}

func (s *GetDataDomainInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataDomainInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataDomainInfoResponse) GetBody() *GetDataDomainInfoResponseBody {
	return s.Body
}

func (s *GetDataDomainInfoResponse) SetHeaders(v map[string]*string) *GetDataDomainInfoResponse {
	s.Headers = v
	return s
}

func (s *GetDataDomainInfoResponse) SetStatusCode(v int32) *GetDataDomainInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataDomainInfoResponse) SetBody(v *GetDataDomainInfoResponseBody) *GetDataDomainInfoResponse {
	s.Body = v
	return s
}

func (s *GetDataDomainInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
