// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTurnServerListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTurnServerListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTurnServerListResponse
	GetStatusCode() *int32
	SetBody(v *GetTurnServerListResponseBody) *GetTurnServerListResponse
	GetBody() *GetTurnServerListResponseBody
}

type GetTurnServerListResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTurnServerListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTurnServerListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTurnServerListResponse) GoString() string {
	return s.String()
}

func (s *GetTurnServerListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTurnServerListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTurnServerListResponse) GetBody() *GetTurnServerListResponseBody {
	return s.Body
}

func (s *GetTurnServerListResponse) SetHeaders(v map[string]*string) *GetTurnServerListResponse {
	s.Headers = v
	return s
}

func (s *GetTurnServerListResponse) SetStatusCode(v int32) *GetTurnServerListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTurnServerListResponse) SetBody(v *GetTurnServerListResponseBody) *GetTurnServerListResponse {
	s.Body = v
	return s
}

func (s *GetTurnServerListResponse) Validate() error {
	return dara.Validate(s)
}
