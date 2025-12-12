// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWarningStrategyConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWarningStrategyConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWarningStrategyConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetWarningStrategyConfigResponseBody) *GetWarningStrategyConfigResponse
	GetBody() *GetWarningStrategyConfigResponseBody
}

type GetWarningStrategyConfigResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWarningStrategyConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWarningStrategyConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWarningStrategyConfigResponse) GoString() string {
	return s.String()
}

func (s *GetWarningStrategyConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWarningStrategyConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWarningStrategyConfigResponse) GetBody() *GetWarningStrategyConfigResponseBody {
	return s.Body
}

func (s *GetWarningStrategyConfigResponse) SetHeaders(v map[string]*string) *GetWarningStrategyConfigResponse {
	s.Headers = v
	return s
}

func (s *GetWarningStrategyConfigResponse) SetStatusCode(v int32) *GetWarningStrategyConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWarningStrategyConfigResponse) SetBody(v *GetWarningStrategyConfigResponseBody) *GetWarningStrategyConfigResponse {
	s.Body = v
	return s
}

func (s *GetWarningStrategyConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
