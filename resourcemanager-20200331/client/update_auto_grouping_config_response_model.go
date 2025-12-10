// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAutoGroupingConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAutoGroupingConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAutoGroupingConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAutoGroupingConfigResponseBody) *UpdateAutoGroupingConfigResponse
	GetBody() *UpdateAutoGroupingConfigResponseBody
}

type UpdateAutoGroupingConfigResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAutoGroupingConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAutoGroupingConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAutoGroupingConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateAutoGroupingConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAutoGroupingConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAutoGroupingConfigResponse) GetBody() *UpdateAutoGroupingConfigResponseBody {
	return s.Body
}

func (s *UpdateAutoGroupingConfigResponse) SetHeaders(v map[string]*string) *UpdateAutoGroupingConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateAutoGroupingConfigResponse) SetStatusCode(v int32) *UpdateAutoGroupingConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAutoGroupingConfigResponse) SetBody(v *UpdateAutoGroupingConfigResponseBody) *UpdateAutoGroupingConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateAutoGroupingConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
