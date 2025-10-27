// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterStrategyCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetClusterStrategyCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetClusterStrategyCountResponse
	GetStatusCode() *int32
	SetBody(v *GetClusterStrategyCountResponseBody) *GetClusterStrategyCountResponse
	GetBody() *GetClusterStrategyCountResponseBody
}

type GetClusterStrategyCountResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetClusterStrategyCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetClusterStrategyCountResponse) String() string {
	return dara.Prettify(s)
}

func (s GetClusterStrategyCountResponse) GoString() string {
	return s.String()
}

func (s *GetClusterStrategyCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetClusterStrategyCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetClusterStrategyCountResponse) GetBody() *GetClusterStrategyCountResponseBody {
	return s.Body
}

func (s *GetClusterStrategyCountResponse) SetHeaders(v map[string]*string) *GetClusterStrategyCountResponse {
	s.Headers = v
	return s
}

func (s *GetClusterStrategyCountResponse) SetStatusCode(v int32) *GetClusterStrategyCountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetClusterStrategyCountResponse) SetBody(v *GetClusterStrategyCountResponseBody) *GetClusterStrategyCountResponse {
	s.Body = v
	return s
}

func (s *GetClusterStrategyCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
