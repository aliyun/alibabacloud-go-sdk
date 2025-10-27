// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAssetsPropertyDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAssetsPropertyDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAssetsPropertyDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetAssetsPropertyDetailResponseBody) *GetAssetsPropertyDetailResponse
	GetBody() *GetAssetsPropertyDetailResponseBody
}

type GetAssetsPropertyDetailResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAssetsPropertyDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAssetsPropertyDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAssetsPropertyDetailResponse) GoString() string {
	return s.String()
}

func (s *GetAssetsPropertyDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAssetsPropertyDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAssetsPropertyDetailResponse) GetBody() *GetAssetsPropertyDetailResponseBody {
	return s.Body
}

func (s *GetAssetsPropertyDetailResponse) SetHeaders(v map[string]*string) *GetAssetsPropertyDetailResponse {
	s.Headers = v
	return s
}

func (s *GetAssetsPropertyDetailResponse) SetStatusCode(v int32) *GetAssetsPropertyDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAssetsPropertyDetailResponse) SetBody(v *GetAssetsPropertyDetailResponseBody) *GetAssetsPropertyDetailResponse {
	s.Body = v
	return s
}

func (s *GetAssetsPropertyDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
