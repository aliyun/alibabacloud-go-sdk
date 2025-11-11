// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMaterialByIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMaterialByIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMaterialByIdResponse
	GetStatusCode() *int32
	SetBody(v *GetMaterialByIdResponseBody) *GetMaterialByIdResponse
	GetBody() *GetMaterialByIdResponseBody
}

type GetMaterialByIdResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMaterialByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMaterialByIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMaterialByIdResponse) GoString() string {
	return s.String()
}

func (s *GetMaterialByIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMaterialByIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMaterialByIdResponse) GetBody() *GetMaterialByIdResponseBody {
	return s.Body
}

func (s *GetMaterialByIdResponse) SetHeaders(v map[string]*string) *GetMaterialByIdResponse {
	s.Headers = v
	return s
}

func (s *GetMaterialByIdResponse) SetStatusCode(v int32) *GetMaterialByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMaterialByIdResponse) SetBody(v *GetMaterialByIdResponseBody) *GetMaterialByIdResponse {
	s.Body = v
	return s
}

func (s *GetMaterialByIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
