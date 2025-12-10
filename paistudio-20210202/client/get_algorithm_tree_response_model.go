// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlgorithmTreeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAlgorithmTreeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAlgorithmTreeResponse
	GetStatusCode() *int32
	SetBody(v *GetAlgorithmTreeResponseBody) *GetAlgorithmTreeResponse
	GetBody() *GetAlgorithmTreeResponseBody
}

type GetAlgorithmTreeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAlgorithmTreeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAlgorithmTreeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAlgorithmTreeResponse) GoString() string {
	return s.String()
}

func (s *GetAlgorithmTreeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAlgorithmTreeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAlgorithmTreeResponse) GetBody() *GetAlgorithmTreeResponseBody {
	return s.Body
}

func (s *GetAlgorithmTreeResponse) SetHeaders(v map[string]*string) *GetAlgorithmTreeResponse {
	s.Headers = v
	return s
}

func (s *GetAlgorithmTreeResponse) SetStatusCode(v int32) *GetAlgorithmTreeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAlgorithmTreeResponse) SetBody(v *GetAlgorithmTreeResponseBody) *GetAlgorithmTreeResponse {
	s.Body = v
	return s
}

func (s *GetAlgorithmTreeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
