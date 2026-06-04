// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaEntityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMetaEntityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMetaEntityResponse
	GetStatusCode() *int32
	SetBody(v *GetMetaEntityResponseBody) *GetMetaEntityResponse
	GetBody() *GetMetaEntityResponseBody
}

type GetMetaEntityResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMetaEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMetaEntityResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMetaEntityResponse) GoString() string {
	return s.String()
}

func (s *GetMetaEntityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMetaEntityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMetaEntityResponse) GetBody() *GetMetaEntityResponseBody {
	return s.Body
}

func (s *GetMetaEntityResponse) SetHeaders(v map[string]*string) *GetMetaEntityResponse {
	s.Headers = v
	return s
}

func (s *GetMetaEntityResponse) SetStatusCode(v int32) *GetMetaEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMetaEntityResponse) SetBody(v *GetMetaEntityResponseBody) *GetMetaEntityResponse {
	s.Body = v
	return s
}

func (s *GetMetaEntityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
