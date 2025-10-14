// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudGtmInstanceConfigLbStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCloudGtmInstanceConfigLbStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCloudGtmInstanceConfigLbStrategyResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCloudGtmInstanceConfigLbStrategyResponseBody) *UpdateCloudGtmInstanceConfigLbStrategyResponse
	GetBody() *UpdateCloudGtmInstanceConfigLbStrategyResponseBody
}

type UpdateCloudGtmInstanceConfigLbStrategyResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCloudGtmInstanceConfigLbStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCloudGtmInstanceConfigLbStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudGtmInstanceConfigLbStrategyResponse) GoString() string {
	return s.String()
}

func (s *UpdateCloudGtmInstanceConfigLbStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCloudGtmInstanceConfigLbStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCloudGtmInstanceConfigLbStrategyResponse) GetBody() *UpdateCloudGtmInstanceConfigLbStrategyResponseBody {
	return s.Body
}

func (s *UpdateCloudGtmInstanceConfigLbStrategyResponse) SetHeaders(v map[string]*string) *UpdateCloudGtmInstanceConfigLbStrategyResponse {
	s.Headers = v
	return s
}

func (s *UpdateCloudGtmInstanceConfigLbStrategyResponse) SetStatusCode(v int32) *UpdateCloudGtmInstanceConfigLbStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCloudGtmInstanceConfigLbStrategyResponse) SetBody(v *UpdateCloudGtmInstanceConfigLbStrategyResponseBody) *UpdateCloudGtmInstanceConfigLbStrategyResponse {
	s.Body = v
	return s
}

func (s *UpdateCloudGtmInstanceConfigLbStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
