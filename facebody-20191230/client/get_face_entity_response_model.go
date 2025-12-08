// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFaceEntityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFaceEntityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFaceEntityResponse
	GetStatusCode() *int32
	SetBody(v *GetFaceEntityResponseBody) *GetFaceEntityResponse
	GetBody() *GetFaceEntityResponseBody
}

type GetFaceEntityResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFaceEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFaceEntityResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFaceEntityResponse) GoString() string {
	return s.String()
}

func (s *GetFaceEntityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFaceEntityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFaceEntityResponse) GetBody() *GetFaceEntityResponseBody {
	return s.Body
}

func (s *GetFaceEntityResponse) SetHeaders(v map[string]*string) *GetFaceEntityResponse {
	s.Headers = v
	return s
}

func (s *GetFaceEntityResponse) SetStatusCode(v int32) *GetFaceEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFaceEntityResponse) SetBody(v *GetFaceEntityResponseBody) *GetFaceEntityResponse {
	s.Body = v
	return s
}

func (s *GetFaceEntityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
