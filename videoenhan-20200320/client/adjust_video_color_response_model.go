// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAdjustVideoColorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AdjustVideoColorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AdjustVideoColorResponse
	GetStatusCode() *int32
	SetBody(v *AdjustVideoColorResponseBody) *AdjustVideoColorResponse
	GetBody() *AdjustVideoColorResponseBody
}

type AdjustVideoColorResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AdjustVideoColorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AdjustVideoColorResponse) String() string {
	return dara.Prettify(s)
}

func (s AdjustVideoColorResponse) GoString() string {
	return s.String()
}

func (s *AdjustVideoColorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AdjustVideoColorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AdjustVideoColorResponse) GetBody() *AdjustVideoColorResponseBody {
	return s.Body
}

func (s *AdjustVideoColorResponse) SetHeaders(v map[string]*string) *AdjustVideoColorResponse {
	s.Headers = v
	return s
}

func (s *AdjustVideoColorResponse) SetStatusCode(v int32) *AdjustVideoColorResponse {
	s.StatusCode = &v
	return s
}

func (s *AdjustVideoColorResponse) SetBody(v *AdjustVideoColorResponseBody) *AdjustVideoColorResponse {
	s.Body = v
	return s
}

func (s *AdjustVideoColorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
