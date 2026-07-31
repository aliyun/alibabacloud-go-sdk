// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSemanticJobDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSemanticJobDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSemanticJobDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetSemanticJobDetailResponseBody) *GetSemanticJobDetailResponse
	GetBody() *GetSemanticJobDetailResponseBody
}

type GetSemanticJobDetailResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSemanticJobDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSemanticJobDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSemanticJobDetailResponse) GoString() string {
	return s.String()
}

func (s *GetSemanticJobDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSemanticJobDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSemanticJobDetailResponse) GetBody() *GetSemanticJobDetailResponseBody {
	return s.Body
}

func (s *GetSemanticJobDetailResponse) SetHeaders(v map[string]*string) *GetSemanticJobDetailResponse {
	s.Headers = v
	return s
}

func (s *GetSemanticJobDetailResponse) SetStatusCode(v int32) *GetSemanticJobDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSemanticJobDetailResponse) SetBody(v *GetSemanticJobDetailResponseBody) *GetSemanticJobDetailResponse {
	s.Body = v
	return s
}

func (s *GetSemanticJobDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
