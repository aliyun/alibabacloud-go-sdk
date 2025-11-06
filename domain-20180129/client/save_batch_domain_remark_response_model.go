// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveBatchDomainRemarkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveBatchDomainRemarkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveBatchDomainRemarkResponse
	GetStatusCode() *int32
	SetBody(v *SaveBatchDomainRemarkResponseBody) *SaveBatchDomainRemarkResponse
	GetBody() *SaveBatchDomainRemarkResponseBody
}

type SaveBatchDomainRemarkResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveBatchDomainRemarkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveBatchDomainRemarkResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveBatchDomainRemarkResponse) GoString() string {
	return s.String()
}

func (s *SaveBatchDomainRemarkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveBatchDomainRemarkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveBatchDomainRemarkResponse) GetBody() *SaveBatchDomainRemarkResponseBody {
	return s.Body
}

func (s *SaveBatchDomainRemarkResponse) SetHeaders(v map[string]*string) *SaveBatchDomainRemarkResponse {
	s.Headers = v
	return s
}

func (s *SaveBatchDomainRemarkResponse) SetStatusCode(v int32) *SaveBatchDomainRemarkResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveBatchDomainRemarkResponse) SetBody(v *SaveBatchDomainRemarkResponseBody) *SaveBatchDomainRemarkResponse {
	s.Body = v
	return s
}

func (s *SaveBatchDomainRemarkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
