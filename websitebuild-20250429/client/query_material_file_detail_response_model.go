// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMaterialFileDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMaterialFileDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMaterialFileDetailResponse
	GetStatusCode() *int32
	SetBody(v *QueryMaterialFileDetailResponseBody) *QueryMaterialFileDetailResponse
	GetBody() *QueryMaterialFileDetailResponseBody
}

type QueryMaterialFileDetailResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMaterialFileDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMaterialFileDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMaterialFileDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryMaterialFileDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMaterialFileDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMaterialFileDetailResponse) GetBody() *QueryMaterialFileDetailResponseBody {
	return s.Body
}

func (s *QueryMaterialFileDetailResponse) SetHeaders(v map[string]*string) *QueryMaterialFileDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryMaterialFileDetailResponse) SetStatusCode(v int32) *QueryMaterialFileDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMaterialFileDetailResponse) SetBody(v *QueryMaterialFileDetailResponseBody) *QueryMaterialFileDetailResponse {
	s.Body = v
	return s
}

func (s *QueryMaterialFileDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
