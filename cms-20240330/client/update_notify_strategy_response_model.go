// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNotifyStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateNotifyStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateNotifyStrategyResponse
	GetStatusCode() *int32
	SetBody(v *UpdateNotifyStrategyResponseBody) *UpdateNotifyStrategyResponse
	GetBody() *UpdateNotifyStrategyResponseBody
}

type UpdateNotifyStrategyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateNotifyStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateNotifyStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateNotifyStrategyResponse) GoString() string {
	return s.String()
}

func (s *UpdateNotifyStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateNotifyStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateNotifyStrategyResponse) GetBody() *UpdateNotifyStrategyResponseBody {
	return s.Body
}

func (s *UpdateNotifyStrategyResponse) SetHeaders(v map[string]*string) *UpdateNotifyStrategyResponse {
	s.Headers = v
	return s
}

func (s *UpdateNotifyStrategyResponse) SetStatusCode(v int32) *UpdateNotifyStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateNotifyStrategyResponse) SetBody(v *UpdateNotifyStrategyResponseBody) *UpdateNotifyStrategyResponse {
	s.Body = v
	return s
}

func (s *UpdateNotifyStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
