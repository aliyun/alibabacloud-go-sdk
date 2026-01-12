// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetArEditCommonMaterialResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetArEditCommonMaterialResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetArEditCommonMaterialResponse
	GetStatusCode() *int32
	SetBody(v *GetArEditCommonMaterialResponseBody) *GetArEditCommonMaterialResponse
	GetBody() *GetArEditCommonMaterialResponseBody
}

type GetArEditCommonMaterialResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetArEditCommonMaterialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetArEditCommonMaterialResponse) String() string {
	return dara.Prettify(s)
}

func (s GetArEditCommonMaterialResponse) GoString() string {
	return s.String()
}

func (s *GetArEditCommonMaterialResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetArEditCommonMaterialResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetArEditCommonMaterialResponse) GetBody() *GetArEditCommonMaterialResponseBody {
	return s.Body
}

func (s *GetArEditCommonMaterialResponse) SetHeaders(v map[string]*string) *GetArEditCommonMaterialResponse {
	s.Headers = v
	return s
}

func (s *GetArEditCommonMaterialResponse) SetStatusCode(v int32) *GetArEditCommonMaterialResponse {
	s.StatusCode = &v
	return s
}

func (s *GetArEditCommonMaterialResponse) SetBody(v *GetArEditCommonMaterialResponseBody) *GetArEditCommonMaterialResponse {
	s.Body = v
	return s
}

func (s *GetArEditCommonMaterialResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
