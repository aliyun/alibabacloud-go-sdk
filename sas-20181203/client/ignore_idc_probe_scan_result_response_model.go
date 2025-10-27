// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIgnoreIdcProbeScanResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *IgnoreIdcProbeScanResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *IgnoreIdcProbeScanResultResponse
	GetStatusCode() *int32
	SetBody(v *IgnoreIdcProbeScanResultResponseBody) *IgnoreIdcProbeScanResultResponse
	GetBody() *IgnoreIdcProbeScanResultResponseBody
}

type IgnoreIdcProbeScanResultResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IgnoreIdcProbeScanResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IgnoreIdcProbeScanResultResponse) String() string {
	return dara.Prettify(s)
}

func (s IgnoreIdcProbeScanResultResponse) GoString() string {
	return s.String()
}

func (s *IgnoreIdcProbeScanResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *IgnoreIdcProbeScanResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *IgnoreIdcProbeScanResultResponse) GetBody() *IgnoreIdcProbeScanResultResponseBody {
	return s.Body
}

func (s *IgnoreIdcProbeScanResultResponse) SetHeaders(v map[string]*string) *IgnoreIdcProbeScanResultResponse {
	s.Headers = v
	return s
}

func (s *IgnoreIdcProbeScanResultResponse) SetStatusCode(v int32) *IgnoreIdcProbeScanResultResponse {
	s.StatusCode = &v
	return s
}

func (s *IgnoreIdcProbeScanResultResponse) SetBody(v *IgnoreIdcProbeScanResultResponseBody) *IgnoreIdcProbeScanResultResponse {
	s.Body = v
	return s
}

func (s *IgnoreIdcProbeScanResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
