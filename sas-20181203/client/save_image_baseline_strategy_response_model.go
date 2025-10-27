// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveImageBaselineStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveImageBaselineStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveImageBaselineStrategyResponse
	GetStatusCode() *int32
	SetBody(v *SaveImageBaselineStrategyResponseBody) *SaveImageBaselineStrategyResponse
	GetBody() *SaveImageBaselineStrategyResponseBody
}

type SaveImageBaselineStrategyResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveImageBaselineStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveImageBaselineStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveImageBaselineStrategyResponse) GoString() string {
	return s.String()
}

func (s *SaveImageBaselineStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveImageBaselineStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveImageBaselineStrategyResponse) GetBody() *SaveImageBaselineStrategyResponseBody {
	return s.Body
}

func (s *SaveImageBaselineStrategyResponse) SetHeaders(v map[string]*string) *SaveImageBaselineStrategyResponse {
	s.Headers = v
	return s
}

func (s *SaveImageBaselineStrategyResponse) SetStatusCode(v int32) *SaveImageBaselineStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveImageBaselineStrategyResponse) SetBody(v *SaveImageBaselineStrategyResponseBody) *SaveImageBaselineStrategyResponse {
	s.Body = v
	return s
}

func (s *SaveImageBaselineStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
