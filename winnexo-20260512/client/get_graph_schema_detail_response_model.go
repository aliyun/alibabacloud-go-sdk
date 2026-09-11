// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGraphSchemaDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetGraphSchemaDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetGraphSchemaDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetGraphSchemaDetailResponseBody) *GetGraphSchemaDetailResponse
	GetBody() *GetGraphSchemaDetailResponseBody
}

type GetGraphSchemaDetailResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGraphSchemaDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGraphSchemaDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetGraphSchemaDetailResponse) GoString() string {
	return s.String()
}

func (s *GetGraphSchemaDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetGraphSchemaDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetGraphSchemaDetailResponse) GetBody() *GetGraphSchemaDetailResponseBody {
	return s.Body
}

func (s *GetGraphSchemaDetailResponse) SetHeaders(v map[string]*string) *GetGraphSchemaDetailResponse {
	s.Headers = v
	return s
}

func (s *GetGraphSchemaDetailResponse) SetStatusCode(v int32) *GetGraphSchemaDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGraphSchemaDetailResponse) SetBody(v *GetGraphSchemaDetailResponseBody) *GetGraphSchemaDetailResponse {
	s.Body = v
	return s
}

func (s *GetGraphSchemaDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
