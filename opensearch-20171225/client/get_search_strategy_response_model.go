// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSearchStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSearchStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSearchStrategyResponse
	GetStatusCode() *int32
	SetBody(v *GetSearchStrategyResponseBody) *GetSearchStrategyResponse
	GetBody() *GetSearchStrategyResponseBody
}

type GetSearchStrategyResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSearchStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSearchStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSearchStrategyResponse) GoString() string {
	return s.String()
}

func (s *GetSearchStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSearchStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSearchStrategyResponse) GetBody() *GetSearchStrategyResponseBody {
	return s.Body
}

func (s *GetSearchStrategyResponse) SetHeaders(v map[string]*string) *GetSearchStrategyResponse {
	s.Headers = v
	return s
}

func (s *GetSearchStrategyResponse) SetStatusCode(v int32) *GetSearchStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSearchStrategyResponse) SetBody(v *GetSearchStrategyResponseBody) *GetSearchStrategyResponse {
	s.Body = v
	return s
}

func (s *GetSearchStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
