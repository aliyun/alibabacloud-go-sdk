// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAttackEventDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAttackEventDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAttackEventDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetAttackEventDetailResponseBody) *GetAttackEventDetailResponse
	GetBody() *GetAttackEventDetailResponseBody
}

type GetAttackEventDetailResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAttackEventDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAttackEventDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAttackEventDetailResponse) GoString() string {
	return s.String()
}

func (s *GetAttackEventDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAttackEventDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAttackEventDetailResponse) GetBody() *GetAttackEventDetailResponseBody {
	return s.Body
}

func (s *GetAttackEventDetailResponse) SetHeaders(v map[string]*string) *GetAttackEventDetailResponse {
	s.Headers = v
	return s
}

func (s *GetAttackEventDetailResponse) SetStatusCode(v int32) *GetAttackEventDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAttackEventDetailResponse) SetBody(v *GetAttackEventDetailResponseBody) *GetAttackEventDetailResponse {
	s.Body = v
	return s
}

func (s *GetAttackEventDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
