// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iId2MetaVerifyWithOCRResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *Id2MetaVerifyWithOCRResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *Id2MetaVerifyWithOCRResponse
	GetStatusCode() *int32
	SetBody(v *Id2MetaVerifyWithOCRResponseBody) *Id2MetaVerifyWithOCRResponse
	GetBody() *Id2MetaVerifyWithOCRResponseBody
}

type Id2MetaVerifyWithOCRResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Id2MetaVerifyWithOCRResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s Id2MetaVerifyWithOCRResponse) String() string {
	return dara.Prettify(s)
}

func (s Id2MetaVerifyWithOCRResponse) GoString() string {
	return s.String()
}

func (s *Id2MetaVerifyWithOCRResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *Id2MetaVerifyWithOCRResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *Id2MetaVerifyWithOCRResponse) GetBody() *Id2MetaVerifyWithOCRResponseBody {
	return s.Body
}

func (s *Id2MetaVerifyWithOCRResponse) SetHeaders(v map[string]*string) *Id2MetaVerifyWithOCRResponse {
	s.Headers = v
	return s
}

func (s *Id2MetaVerifyWithOCRResponse) SetStatusCode(v int32) *Id2MetaVerifyWithOCRResponse {
	s.StatusCode = &v
	return s
}

func (s *Id2MetaVerifyWithOCRResponse) SetBody(v *Id2MetaVerifyWithOCRResponseBody) *Id2MetaVerifyWithOCRResponse {
	s.Body = v
	return s
}

func (s *Id2MetaVerifyWithOCRResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
