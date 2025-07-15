// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStateConfigurationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteStateConfigurationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteStateConfigurationsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteStateConfigurationsResponseBody) *DeleteStateConfigurationsResponse
	GetBody() *DeleteStateConfigurationsResponseBody
}

type DeleteStateConfigurationsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteStateConfigurationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteStateConfigurationsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteStateConfigurationsResponse) GoString() string {
	return s.String()
}

func (s *DeleteStateConfigurationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteStateConfigurationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteStateConfigurationsResponse) GetBody() *DeleteStateConfigurationsResponseBody {
	return s.Body
}

func (s *DeleteStateConfigurationsResponse) SetHeaders(v map[string]*string) *DeleteStateConfigurationsResponse {
	s.Headers = v
	return s
}

func (s *DeleteStateConfigurationsResponse) SetStatusCode(v int32) *DeleteStateConfigurationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteStateConfigurationsResponse) SetBody(v *DeleteStateConfigurationsResponseBody) *DeleteStateConfigurationsResponse {
	s.Body = v
	return s
}

func (s *DeleteStateConfigurationsResponse) Validate() error {
	return dara.Validate(s)
}
