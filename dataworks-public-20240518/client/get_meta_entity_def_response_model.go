// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaEntityDefResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMetaEntityDefResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMetaEntityDefResponse
	GetStatusCode() *int32
	SetBody(v *GetMetaEntityDefResponseBody) *GetMetaEntityDefResponse
	GetBody() *GetMetaEntityDefResponseBody
}

type GetMetaEntityDefResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMetaEntityDefResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMetaEntityDefResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMetaEntityDefResponse) GoString() string {
	return s.String()
}

func (s *GetMetaEntityDefResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMetaEntityDefResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMetaEntityDefResponse) GetBody() *GetMetaEntityDefResponseBody {
	return s.Body
}

func (s *GetMetaEntityDefResponse) SetHeaders(v map[string]*string) *GetMetaEntityDefResponse {
	s.Headers = v
	return s
}

func (s *GetMetaEntityDefResponse) SetStatusCode(v int32) *GetMetaEntityDefResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMetaEntityDefResponse) SetBody(v *GetMetaEntityDefResponseBody) *GetMetaEntityDefResponse {
	s.Body = v
	return s
}

func (s *GetMetaEntityDefResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
