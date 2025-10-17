// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshLivyComputeTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RefreshLivyComputeTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RefreshLivyComputeTokenResponse
	GetStatusCode() *int32
	SetBody(v *RefreshLivyComputeTokenResponseBody) *RefreshLivyComputeTokenResponse
	GetBody() *RefreshLivyComputeTokenResponseBody
}

type RefreshLivyComputeTokenResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefreshLivyComputeTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefreshLivyComputeTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s RefreshLivyComputeTokenResponse) GoString() string {
	return s.String()
}

func (s *RefreshLivyComputeTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RefreshLivyComputeTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RefreshLivyComputeTokenResponse) GetBody() *RefreshLivyComputeTokenResponseBody {
	return s.Body
}

func (s *RefreshLivyComputeTokenResponse) SetHeaders(v map[string]*string) *RefreshLivyComputeTokenResponse {
	s.Headers = v
	return s
}

func (s *RefreshLivyComputeTokenResponse) SetStatusCode(v int32) *RefreshLivyComputeTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshLivyComputeTokenResponse) SetBody(v *RefreshLivyComputeTokenResponseBody) *RefreshLivyComputeTokenResponse {
	s.Body = v
	return s
}

func (s *RefreshLivyComputeTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
