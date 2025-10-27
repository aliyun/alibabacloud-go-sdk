// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartIdcProbeScanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartIdcProbeScanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartIdcProbeScanResponse
	GetStatusCode() *int32
	SetBody(v *StartIdcProbeScanResponseBody) *StartIdcProbeScanResponse
	GetBody() *StartIdcProbeScanResponseBody
}

type StartIdcProbeScanResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartIdcProbeScanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartIdcProbeScanResponse) String() string {
	return dara.Prettify(s)
}

func (s StartIdcProbeScanResponse) GoString() string {
	return s.String()
}

func (s *StartIdcProbeScanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartIdcProbeScanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartIdcProbeScanResponse) GetBody() *StartIdcProbeScanResponseBody {
	return s.Body
}

func (s *StartIdcProbeScanResponse) SetHeaders(v map[string]*string) *StartIdcProbeScanResponse {
	s.Headers = v
	return s
}

func (s *StartIdcProbeScanResponse) SetStatusCode(v int32) *StartIdcProbeScanResponse {
	s.StatusCode = &v
	return s
}

func (s *StartIdcProbeScanResponse) SetBody(v *StartIdcProbeScanResponseBody) *StartIdcProbeScanResponse {
	s.Body = v
	return s
}

func (s *StartIdcProbeScanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
