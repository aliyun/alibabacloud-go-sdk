// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChangeOrderInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetChangeOrderInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetChangeOrderInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetChangeOrderInfoResponseBody) *GetChangeOrderInfoResponse
	GetBody() *GetChangeOrderInfoResponseBody
}

type GetChangeOrderInfoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetChangeOrderInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetChangeOrderInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetChangeOrderInfoResponse) GoString() string {
	return s.String()
}

func (s *GetChangeOrderInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetChangeOrderInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetChangeOrderInfoResponse) GetBody() *GetChangeOrderInfoResponseBody {
	return s.Body
}

func (s *GetChangeOrderInfoResponse) SetHeaders(v map[string]*string) *GetChangeOrderInfoResponse {
	s.Headers = v
	return s
}

func (s *GetChangeOrderInfoResponse) SetStatusCode(v int32) *GetChangeOrderInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetChangeOrderInfoResponse) SetBody(v *GetChangeOrderInfoResponseBody) *GetChangeOrderInfoResponse {
	s.Body = v
	return s
}

func (s *GetChangeOrderInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
