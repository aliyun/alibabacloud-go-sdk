// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryOperationAuditInfoListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryOperationAuditInfoListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryOperationAuditInfoListResponse
	GetStatusCode() *int32
	SetBody(v *QueryOperationAuditInfoListResponseBody) *QueryOperationAuditInfoListResponse
	GetBody() *QueryOperationAuditInfoListResponseBody
}

type QueryOperationAuditInfoListResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryOperationAuditInfoListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryOperationAuditInfoListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryOperationAuditInfoListResponse) GoString() string {
	return s.String()
}

func (s *QueryOperationAuditInfoListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryOperationAuditInfoListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryOperationAuditInfoListResponse) GetBody() *QueryOperationAuditInfoListResponseBody {
	return s.Body
}

func (s *QueryOperationAuditInfoListResponse) SetHeaders(v map[string]*string) *QueryOperationAuditInfoListResponse {
	s.Headers = v
	return s
}

func (s *QueryOperationAuditInfoListResponse) SetStatusCode(v int32) *QueryOperationAuditInfoListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryOperationAuditInfoListResponse) SetBody(v *QueryOperationAuditInfoListResponseBody) *QueryOperationAuditInfoListResponse {
	s.Body = v
	return s
}

func (s *QueryOperationAuditInfoListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
