// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflineLaboratoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OfflineLaboratoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OfflineLaboratoryResponse
	GetStatusCode() *int32
	SetBody(v *OfflineLaboratoryResponseBody) *OfflineLaboratoryResponse
	GetBody() *OfflineLaboratoryResponseBody
}

type OfflineLaboratoryResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OfflineLaboratoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OfflineLaboratoryResponse) String() string {
	return dara.Prettify(s)
}

func (s OfflineLaboratoryResponse) GoString() string {
	return s.String()
}

func (s *OfflineLaboratoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OfflineLaboratoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OfflineLaboratoryResponse) GetBody() *OfflineLaboratoryResponseBody {
	return s.Body
}

func (s *OfflineLaboratoryResponse) SetHeaders(v map[string]*string) *OfflineLaboratoryResponse {
	s.Headers = v
	return s
}

func (s *OfflineLaboratoryResponse) SetStatusCode(v int32) *OfflineLaboratoryResponse {
	s.StatusCode = &v
	return s
}

func (s *OfflineLaboratoryResponse) SetBody(v *OfflineLaboratoryResponseBody) *OfflineLaboratoryResponse {
	s.Body = v
	return s
}

func (s *OfflineLaboratoryResponse) Validate() error {
	return dara.Validate(s)
}
