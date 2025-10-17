// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLivyComputeTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLivyComputeTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLivyComputeTokenResponse
	GetStatusCode() *int32
	SetBody(v *GetLivyComputeTokenResponseBody) *GetLivyComputeTokenResponse
	GetBody() *GetLivyComputeTokenResponseBody
}

type GetLivyComputeTokenResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLivyComputeTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLivyComputeTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLivyComputeTokenResponse) GoString() string {
	return s.String()
}

func (s *GetLivyComputeTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLivyComputeTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLivyComputeTokenResponse) GetBody() *GetLivyComputeTokenResponseBody {
	return s.Body
}

func (s *GetLivyComputeTokenResponse) SetHeaders(v map[string]*string) *GetLivyComputeTokenResponse {
	s.Headers = v
	return s
}

func (s *GetLivyComputeTokenResponse) SetStatusCode(v int32) *GetLivyComputeTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLivyComputeTokenResponse) SetBody(v *GetLivyComputeTokenResponseBody) *GetLivyComputeTokenResponse {
	s.Body = v
	return s
}

func (s *GetLivyComputeTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
