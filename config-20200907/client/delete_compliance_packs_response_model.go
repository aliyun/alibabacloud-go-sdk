// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCompliancePacksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCompliancePacksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCompliancePacksResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCompliancePacksResponseBody) *DeleteCompliancePacksResponse
	GetBody() *DeleteCompliancePacksResponseBody
}

type DeleteCompliancePacksResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCompliancePacksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCompliancePacksResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCompliancePacksResponse) GoString() string {
	return s.String()
}

func (s *DeleteCompliancePacksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCompliancePacksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCompliancePacksResponse) GetBody() *DeleteCompliancePacksResponseBody {
	return s.Body
}

func (s *DeleteCompliancePacksResponse) SetHeaders(v map[string]*string) *DeleteCompliancePacksResponse {
	s.Headers = v
	return s
}

func (s *DeleteCompliancePacksResponse) SetStatusCode(v int32) *DeleteCompliancePacksResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCompliancePacksResponse) SetBody(v *DeleteCompliancePacksResponseBody) *DeleteCompliancePacksResponse {
	s.Body = v
	return s
}

func (s *DeleteCompliancePacksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
