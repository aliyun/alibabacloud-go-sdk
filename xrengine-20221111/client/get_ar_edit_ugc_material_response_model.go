// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetArEditUgcMaterialResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetArEditUgcMaterialResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetArEditUgcMaterialResponse
	GetStatusCode() *int32
	SetBody(v *GetArEditUgcMaterialResponseBody) *GetArEditUgcMaterialResponse
	GetBody() *GetArEditUgcMaterialResponseBody
}

type GetArEditUgcMaterialResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetArEditUgcMaterialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetArEditUgcMaterialResponse) String() string {
	return dara.Prettify(s)
}

func (s GetArEditUgcMaterialResponse) GoString() string {
	return s.String()
}

func (s *GetArEditUgcMaterialResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetArEditUgcMaterialResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetArEditUgcMaterialResponse) GetBody() *GetArEditUgcMaterialResponseBody {
	return s.Body
}

func (s *GetArEditUgcMaterialResponse) SetHeaders(v map[string]*string) *GetArEditUgcMaterialResponse {
	s.Headers = v
	return s
}

func (s *GetArEditUgcMaterialResponse) SetStatusCode(v int32) *GetArEditUgcMaterialResponse {
	s.StatusCode = &v
	return s
}

func (s *GetArEditUgcMaterialResponse) SetBody(v *GetArEditUgcMaterialResponseBody) *GetArEditUgcMaterialResponse {
	s.Body = v
	return s
}

func (s *GetArEditUgcMaterialResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
