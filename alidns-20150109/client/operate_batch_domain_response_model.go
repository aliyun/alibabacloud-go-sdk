// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateBatchDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperateBatchDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperateBatchDomainResponse
	GetStatusCode() *int32
	SetBody(v *OperateBatchDomainResponseBody) *OperateBatchDomainResponse
	GetBody() *OperateBatchDomainResponseBody
}

type OperateBatchDomainResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateBatchDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateBatchDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s OperateBatchDomainResponse) GoString() string {
	return s.String()
}

func (s *OperateBatchDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperateBatchDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperateBatchDomainResponse) GetBody() *OperateBatchDomainResponseBody {
	return s.Body
}

func (s *OperateBatchDomainResponse) SetHeaders(v map[string]*string) *OperateBatchDomainResponse {
	s.Headers = v
	return s
}

func (s *OperateBatchDomainResponse) SetStatusCode(v int32) *OperateBatchDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateBatchDomainResponse) SetBody(v *OperateBatchDomainResponseBody) *OperateBatchDomainResponse {
	s.Body = v
	return s
}

func (s *OperateBatchDomainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
