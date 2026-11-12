// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEffectiveImageByIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEffectiveImageByIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEffectiveImageByIdResponse
	GetStatusCode() *int32
	SetBody(v *GetEffectiveImageByIdResponseBody) *GetEffectiveImageByIdResponse
	GetBody() *GetEffectiveImageByIdResponseBody
}

type GetEffectiveImageByIdResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEffectiveImageByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEffectiveImageByIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEffectiveImageByIdResponse) GoString() string {
	return s.String()
}

func (s *GetEffectiveImageByIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEffectiveImageByIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEffectiveImageByIdResponse) GetBody() *GetEffectiveImageByIdResponseBody {
	return s.Body
}

func (s *GetEffectiveImageByIdResponse) SetHeaders(v map[string]*string) *GetEffectiveImageByIdResponse {
	s.Headers = v
	return s
}

func (s *GetEffectiveImageByIdResponse) SetStatusCode(v int32) *GetEffectiveImageByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEffectiveImageByIdResponse) SetBody(v *GetEffectiveImageByIdResponseBody) *GetEffectiveImageByIdResponse {
	s.Body = v
	return s
}

func (s *GetEffectiveImageByIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
