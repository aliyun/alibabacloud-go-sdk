// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLivyComputeTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLivyComputeTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLivyComputeTokenResponse
	GetStatusCode() *int32
	SetBody(v *ListLivyComputeTokenResponseBody) *ListLivyComputeTokenResponse
	GetBody() *ListLivyComputeTokenResponseBody
}

type ListLivyComputeTokenResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLivyComputeTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLivyComputeTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLivyComputeTokenResponse) GoString() string {
	return s.String()
}

func (s *ListLivyComputeTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLivyComputeTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLivyComputeTokenResponse) GetBody() *ListLivyComputeTokenResponseBody {
	return s.Body
}

func (s *ListLivyComputeTokenResponse) SetHeaders(v map[string]*string) *ListLivyComputeTokenResponse {
	s.Headers = v
	return s
}

func (s *ListLivyComputeTokenResponse) SetStatusCode(v int32) *ListLivyComputeTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLivyComputeTokenResponse) SetBody(v *ListLivyComputeTokenResponseBody) *ListLivyComputeTokenResponse {
	s.Body = v
	return s
}

func (s *ListLivyComputeTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
