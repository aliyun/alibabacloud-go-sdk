// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataMaskingColumnCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataMaskingColumnCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataMaskingColumnCountResponse
	GetStatusCode() *int32
	SetBody(v *GetDataMaskingColumnCountResponseBody) *GetDataMaskingColumnCountResponse
	GetBody() *GetDataMaskingColumnCountResponseBody
}

type GetDataMaskingColumnCountResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataMaskingColumnCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataMaskingColumnCountResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataMaskingColumnCountResponse) GoString() string {
	return s.String()
}

func (s *GetDataMaskingColumnCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataMaskingColumnCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataMaskingColumnCountResponse) GetBody() *GetDataMaskingColumnCountResponseBody {
	return s.Body
}

func (s *GetDataMaskingColumnCountResponse) SetHeaders(v map[string]*string) *GetDataMaskingColumnCountResponse {
	s.Headers = v
	return s
}

func (s *GetDataMaskingColumnCountResponse) SetStatusCode(v int32) *GetDataMaskingColumnCountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataMaskingColumnCountResponse) SetBody(v *GetDataMaskingColumnCountResponseBody) *GetDataMaskingColumnCountResponse {
	s.Body = v
	return s
}

func (s *GetDataMaskingColumnCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
