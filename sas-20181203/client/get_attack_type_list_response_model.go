// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAttackTypeListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAttackTypeListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAttackTypeListResponse
	GetStatusCode() *int32
	SetBody(v *GetAttackTypeListResponseBody) *GetAttackTypeListResponse
	GetBody() *GetAttackTypeListResponseBody
}

type GetAttackTypeListResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAttackTypeListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAttackTypeListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAttackTypeListResponse) GoString() string {
	return s.String()
}

func (s *GetAttackTypeListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAttackTypeListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAttackTypeListResponse) GetBody() *GetAttackTypeListResponseBody {
	return s.Body
}

func (s *GetAttackTypeListResponse) SetHeaders(v map[string]*string) *GetAttackTypeListResponse {
	s.Headers = v
	return s
}

func (s *GetAttackTypeListResponse) SetStatusCode(v int32) *GetAttackTypeListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAttackTypeListResponse) SetBody(v *GetAttackTypeListResponseBody) *GetAttackTypeListResponse {
	s.Body = v
	return s
}

func (s *GetAttackTypeListResponse) Validate() error {
	return dara.Validate(s)
}
