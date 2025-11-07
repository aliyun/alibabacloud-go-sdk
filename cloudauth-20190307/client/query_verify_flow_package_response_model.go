// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryVerifyFlowPackageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryVerifyFlowPackageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryVerifyFlowPackageResponse
	GetStatusCode() *int32
	SetBody(v *QueryVerifyFlowPackageResponseBody) *QueryVerifyFlowPackageResponse
	GetBody() *QueryVerifyFlowPackageResponseBody
}

type QueryVerifyFlowPackageResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryVerifyFlowPackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryVerifyFlowPackageResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryVerifyFlowPackageResponse) GoString() string {
	return s.String()
}

func (s *QueryVerifyFlowPackageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryVerifyFlowPackageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryVerifyFlowPackageResponse) GetBody() *QueryVerifyFlowPackageResponseBody {
	return s.Body
}

func (s *QueryVerifyFlowPackageResponse) SetHeaders(v map[string]*string) *QueryVerifyFlowPackageResponse {
	s.Headers = v
	return s
}

func (s *QueryVerifyFlowPackageResponse) SetStatusCode(v int32) *QueryVerifyFlowPackageResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryVerifyFlowPackageResponse) SetBody(v *QueryVerifyFlowPackageResponseBody) *QueryVerifyFlowPackageResponse {
	s.Body = v
	return s
}

func (s *QueryVerifyFlowPackageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
