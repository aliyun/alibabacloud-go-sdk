// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportEnterpriseAccelerateTargetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImportEnterpriseAccelerateTargetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImportEnterpriseAccelerateTargetsResponse
	GetStatusCode() *int32
	SetBody(v *ImportEnterpriseAccelerateTargetsResponseBody) *ImportEnterpriseAccelerateTargetsResponse
	GetBody() *ImportEnterpriseAccelerateTargetsResponseBody
}

type ImportEnterpriseAccelerateTargetsResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportEnterpriseAccelerateTargetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportEnterpriseAccelerateTargetsResponse) String() string {
	return dara.Prettify(s)
}

func (s ImportEnterpriseAccelerateTargetsResponse) GoString() string {
	return s.String()
}

func (s *ImportEnterpriseAccelerateTargetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImportEnterpriseAccelerateTargetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImportEnterpriseAccelerateTargetsResponse) GetBody() *ImportEnterpriseAccelerateTargetsResponseBody {
	return s.Body
}

func (s *ImportEnterpriseAccelerateTargetsResponse) SetHeaders(v map[string]*string) *ImportEnterpriseAccelerateTargetsResponse {
	s.Headers = v
	return s
}

func (s *ImportEnterpriseAccelerateTargetsResponse) SetStatusCode(v int32) *ImportEnterpriseAccelerateTargetsResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportEnterpriseAccelerateTargetsResponse) SetBody(v *ImportEnterpriseAccelerateTargetsResponseBody) *ImportEnterpriseAccelerateTargetsResponse {
	s.Body = v
	return s
}

func (s *ImportEnterpriseAccelerateTargetsResponse) Validate() error {
	return dara.Validate(s)
}
