// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribleCertListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribleCertListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribleCertListResponse
	GetStatusCode() *int32
	SetBody(v *DescribleCertListResponseBody) *DescribleCertListResponse
	GetBody() *DescribleCertListResponseBody
}

type DescribleCertListResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribleCertListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribleCertListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribleCertListResponse) GoString() string {
	return s.String()
}

func (s *DescribleCertListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribleCertListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribleCertListResponse) GetBody() *DescribleCertListResponseBody {
	return s.Body
}

func (s *DescribleCertListResponse) SetHeaders(v map[string]*string) *DescribleCertListResponse {
	s.Headers = v
	return s
}

func (s *DescribleCertListResponse) SetStatusCode(v int32) *DescribleCertListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribleCertListResponse) SetBody(v *DescribleCertListResponseBody) *DescribleCertListResponse {
	s.Body = v
	return s
}

func (s *DescribleCertListResponse) Validate() error {
	return dara.Validate(s)
}
