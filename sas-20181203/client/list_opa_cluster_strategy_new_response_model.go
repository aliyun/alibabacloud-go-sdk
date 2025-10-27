// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOpaClusterStrategyNewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOpaClusterStrategyNewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOpaClusterStrategyNewResponse
	GetStatusCode() *int32
	SetBody(v *ListOpaClusterStrategyNewResponseBody) *ListOpaClusterStrategyNewResponse
	GetBody() *ListOpaClusterStrategyNewResponseBody
}

type ListOpaClusterStrategyNewResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOpaClusterStrategyNewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOpaClusterStrategyNewResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOpaClusterStrategyNewResponse) GoString() string {
	return s.String()
}

func (s *ListOpaClusterStrategyNewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOpaClusterStrategyNewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOpaClusterStrategyNewResponse) GetBody() *ListOpaClusterStrategyNewResponseBody {
	return s.Body
}

func (s *ListOpaClusterStrategyNewResponse) SetHeaders(v map[string]*string) *ListOpaClusterStrategyNewResponse {
	s.Headers = v
	return s
}

func (s *ListOpaClusterStrategyNewResponse) SetStatusCode(v int32) *ListOpaClusterStrategyNewResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOpaClusterStrategyNewResponse) SetBody(v *ListOpaClusterStrategyNewResponseBody) *ListOpaClusterStrategyNewResponse {
	s.Body = v
	return s
}

func (s *ListOpaClusterStrategyNewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
