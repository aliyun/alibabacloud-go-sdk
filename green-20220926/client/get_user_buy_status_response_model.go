// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserBuyStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserBuyStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserBuyStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetUserBuyStatusResponseBody) *GetUserBuyStatusResponse
	GetBody() *GetUserBuyStatusResponseBody
}

type GetUserBuyStatusResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserBuyStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserBuyStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserBuyStatusResponse) GoString() string {
	return s.String()
}

func (s *GetUserBuyStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserBuyStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserBuyStatusResponse) GetBody() *GetUserBuyStatusResponseBody {
	return s.Body
}

func (s *GetUserBuyStatusResponse) SetHeaders(v map[string]*string) *GetUserBuyStatusResponse {
	s.Headers = v
	return s
}

func (s *GetUserBuyStatusResponse) SetStatusCode(v int32) *GetUserBuyStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserBuyStatusResponse) SetBody(v *GetUserBuyStatusResponseBody) *GetUserBuyStatusResponse {
	s.Body = v
	return s
}

func (s *GetUserBuyStatusResponse) Validate() error {
	return dara.Validate(s)
}
