// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDirectSupplyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDirectSupplyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDirectSupplyResponse
	GetStatusCode() *int32
	SetBody(v *GetDirectSupplyResponseBody) *GetDirectSupplyResponse
	GetBody() *GetDirectSupplyResponseBody
}

type GetDirectSupplyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDirectSupplyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDirectSupplyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDirectSupplyResponse) GoString() string {
	return s.String()
}

func (s *GetDirectSupplyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDirectSupplyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDirectSupplyResponse) GetBody() *GetDirectSupplyResponseBody {
	return s.Body
}

func (s *GetDirectSupplyResponse) SetHeaders(v map[string]*string) *GetDirectSupplyResponse {
	s.Headers = v
	return s
}

func (s *GetDirectSupplyResponse) SetStatusCode(v int32) *GetDirectSupplyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDirectSupplyResponse) SetBody(v *GetDirectSupplyResponseBody) *GetDirectSupplyResponse {
	s.Body = v
	return s
}

func (s *GetDirectSupplyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
