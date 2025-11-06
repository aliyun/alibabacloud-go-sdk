// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryOperationAuditInfoDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryOperationAuditInfoDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryOperationAuditInfoDetailResponse
	GetStatusCode() *int32
	SetBody(v *QueryOperationAuditInfoDetailResponseBody) *QueryOperationAuditInfoDetailResponse
	GetBody() *QueryOperationAuditInfoDetailResponseBody
}

type QueryOperationAuditInfoDetailResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryOperationAuditInfoDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryOperationAuditInfoDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryOperationAuditInfoDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryOperationAuditInfoDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryOperationAuditInfoDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryOperationAuditInfoDetailResponse) GetBody() *QueryOperationAuditInfoDetailResponseBody {
	return s.Body
}

func (s *QueryOperationAuditInfoDetailResponse) SetHeaders(v map[string]*string) *QueryOperationAuditInfoDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryOperationAuditInfoDetailResponse) SetStatusCode(v int32) *QueryOperationAuditInfoDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryOperationAuditInfoDetailResponse) SetBody(v *QueryOperationAuditInfoDetailResponseBody) *QueryOperationAuditInfoDetailResponse {
	s.Body = v
	return s
}

func (s *QueryOperationAuditInfoDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
