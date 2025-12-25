// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBotAppKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBotAppKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBotAppKeyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBotAppKeyResponseBody) *DescribeBotAppKeyResponse
	GetBody() *DescribeBotAppKeyResponseBody
}

type DescribeBotAppKeyResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBotAppKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBotAppKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBotAppKeyResponse) GoString() string {
	return s.String()
}

func (s *DescribeBotAppKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBotAppKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBotAppKeyResponse) GetBody() *DescribeBotAppKeyResponseBody {
	return s.Body
}

func (s *DescribeBotAppKeyResponse) SetHeaders(v map[string]*string) *DescribeBotAppKeyResponse {
	s.Headers = v
	return s
}

func (s *DescribeBotAppKeyResponse) SetStatusCode(v int32) *DescribeBotAppKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBotAppKeyResponse) SetBody(v *DescribeBotAppKeyResponseBody) *DescribeBotAppKeyResponse {
	s.Body = v
	return s
}

func (s *DescribeBotAppKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
