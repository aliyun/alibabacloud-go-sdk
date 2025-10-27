// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAttackPathEventDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAttackPathEventDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAttackPathEventDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetAttackPathEventDetailResponseBody) *GetAttackPathEventDetailResponse
	GetBody() *GetAttackPathEventDetailResponseBody
}

type GetAttackPathEventDetailResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAttackPathEventDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAttackPathEventDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAttackPathEventDetailResponse) GoString() string {
	return s.String()
}

func (s *GetAttackPathEventDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAttackPathEventDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAttackPathEventDetailResponse) GetBody() *GetAttackPathEventDetailResponseBody {
	return s.Body
}

func (s *GetAttackPathEventDetailResponse) SetHeaders(v map[string]*string) *GetAttackPathEventDetailResponse {
	s.Headers = v
	return s
}

func (s *GetAttackPathEventDetailResponse) SetStatusCode(v int32) *GetAttackPathEventDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAttackPathEventDetailResponse) SetBody(v *GetAttackPathEventDetailResponseBody) *GetAttackPathEventDetailResponse {
	s.Body = v
	return s
}

func (s *GetAttackPathEventDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
