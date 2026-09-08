// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataMaskingAccountCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataMaskingAccountCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataMaskingAccountCountResponse
	GetStatusCode() *int32
	SetBody(v *GetDataMaskingAccountCountResponseBody) *GetDataMaskingAccountCountResponse
	GetBody() *GetDataMaskingAccountCountResponseBody
}

type GetDataMaskingAccountCountResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataMaskingAccountCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataMaskingAccountCountResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataMaskingAccountCountResponse) GoString() string {
	return s.String()
}

func (s *GetDataMaskingAccountCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataMaskingAccountCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataMaskingAccountCountResponse) GetBody() *GetDataMaskingAccountCountResponseBody {
	return s.Body
}

func (s *GetDataMaskingAccountCountResponse) SetHeaders(v map[string]*string) *GetDataMaskingAccountCountResponse {
	s.Headers = v
	return s
}

func (s *GetDataMaskingAccountCountResponse) SetStatusCode(v int32) *GetDataMaskingAccountCountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataMaskingAccountCountResponse) SetBody(v *GetDataMaskingAccountCountResponseBody) *GetDataMaskingAccountCountResponse {
	s.Body = v
	return s
}

func (s *GetDataMaskingAccountCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
