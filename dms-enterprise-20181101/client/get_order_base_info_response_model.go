// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOrderBaseInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOrderBaseInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOrderBaseInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetOrderBaseInfoResponseBody) *GetOrderBaseInfoResponse
	GetBody() *GetOrderBaseInfoResponseBody
}

type GetOrderBaseInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOrderBaseInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOrderBaseInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOrderBaseInfoResponse) GoString() string {
	return s.String()
}

func (s *GetOrderBaseInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOrderBaseInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOrderBaseInfoResponse) GetBody() *GetOrderBaseInfoResponseBody {
	return s.Body
}

func (s *GetOrderBaseInfoResponse) SetHeaders(v map[string]*string) *GetOrderBaseInfoResponse {
	s.Headers = v
	return s
}

func (s *GetOrderBaseInfoResponse) SetStatusCode(v int32) *GetOrderBaseInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOrderBaseInfoResponse) SetBody(v *GetOrderBaseInfoResponseBody) *GetOrderBaseInfoResponse {
	s.Body = v
	return s
}

func (s *GetOrderBaseInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
