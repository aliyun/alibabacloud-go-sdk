// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppGroupQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAppGroupQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAppGroupQuotaResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAppGroupQuotaResponseBody) *ModifyAppGroupQuotaResponse
	GetBody() *ModifyAppGroupQuotaResponseBody
}

type ModifyAppGroupQuotaResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAppGroupQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAppGroupQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppGroupQuotaResponse) GoString() string {
	return s.String()
}

func (s *ModifyAppGroupQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAppGroupQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAppGroupQuotaResponse) GetBody() *ModifyAppGroupQuotaResponseBody {
	return s.Body
}

func (s *ModifyAppGroupQuotaResponse) SetHeaders(v map[string]*string) *ModifyAppGroupQuotaResponse {
	s.Headers = v
	return s
}

func (s *ModifyAppGroupQuotaResponse) SetStatusCode(v int32) *ModifyAppGroupQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAppGroupQuotaResponse) SetBody(v *ModifyAppGroupQuotaResponseBody) *ModifyAppGroupQuotaResponse {
	s.Body = v
	return s
}

func (s *ModifyAppGroupQuotaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
