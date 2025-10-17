// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApsLifecycleStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListApsLifecycleStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListApsLifecycleStrategyResponse
	GetStatusCode() *int32
	SetBody(v *ListApsLifecycleStrategyResponseBody) *ListApsLifecycleStrategyResponse
	GetBody() *ListApsLifecycleStrategyResponseBody
}

type ListApsLifecycleStrategyResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApsLifecycleStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApsLifecycleStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s ListApsLifecycleStrategyResponse) GoString() string {
	return s.String()
}

func (s *ListApsLifecycleStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListApsLifecycleStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListApsLifecycleStrategyResponse) GetBody() *ListApsLifecycleStrategyResponseBody {
	return s.Body
}

func (s *ListApsLifecycleStrategyResponse) SetHeaders(v map[string]*string) *ListApsLifecycleStrategyResponse {
	s.Headers = v
	return s
}

func (s *ListApsLifecycleStrategyResponse) SetStatusCode(v int32) *ListApsLifecycleStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApsLifecycleStrategyResponse) SetBody(v *ListApsLifecycleStrategyResponseBody) *ListApsLifecycleStrategyResponse {
	s.Body = v
	return s
}

func (s *ListApsLifecycleStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
