// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCommonBuyUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCommonBuyUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCommonBuyUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetCommonBuyUrlResponseBody) *GetCommonBuyUrlResponse
	GetBody() *GetCommonBuyUrlResponseBody
}

type GetCommonBuyUrlResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCommonBuyUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCommonBuyUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCommonBuyUrlResponse) GoString() string {
	return s.String()
}

func (s *GetCommonBuyUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCommonBuyUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCommonBuyUrlResponse) GetBody() *GetCommonBuyUrlResponseBody {
	return s.Body
}

func (s *GetCommonBuyUrlResponse) SetHeaders(v map[string]*string) *GetCommonBuyUrlResponse {
	s.Headers = v
	return s
}

func (s *GetCommonBuyUrlResponse) SetStatusCode(v int32) *GetCommonBuyUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCommonBuyUrlResponse) SetBody(v *GetCommonBuyUrlResponseBody) *GetCommonBuyUrlResponse {
	s.Body = v
	return s
}

func (s *GetCommonBuyUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
