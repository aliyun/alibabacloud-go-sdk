// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iId3MetaVerifyWithOCRResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *Id3MetaVerifyWithOCRResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *Id3MetaVerifyWithOCRResponse
	GetStatusCode() *int32
	SetBody(v *Id3MetaVerifyWithOCRResponseBody) *Id3MetaVerifyWithOCRResponse
	GetBody() *Id3MetaVerifyWithOCRResponseBody
}

type Id3MetaVerifyWithOCRResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Id3MetaVerifyWithOCRResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s Id3MetaVerifyWithOCRResponse) String() string {
	return dara.Prettify(s)
}

func (s Id3MetaVerifyWithOCRResponse) GoString() string {
	return s.String()
}

func (s *Id3MetaVerifyWithOCRResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *Id3MetaVerifyWithOCRResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *Id3MetaVerifyWithOCRResponse) GetBody() *Id3MetaVerifyWithOCRResponseBody {
	return s.Body
}

func (s *Id3MetaVerifyWithOCRResponse) SetHeaders(v map[string]*string) *Id3MetaVerifyWithOCRResponse {
	s.Headers = v
	return s
}

func (s *Id3MetaVerifyWithOCRResponse) SetStatusCode(v int32) *Id3MetaVerifyWithOCRResponse {
	s.StatusCode = &v
	return s
}

func (s *Id3MetaVerifyWithOCRResponse) SetBody(v *Id3MetaVerifyWithOCRResponseBody) *Id3MetaVerifyWithOCRResponse {
	s.Body = v
	return s
}

func (s *Id3MetaVerifyWithOCRResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
