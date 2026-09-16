// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iId3MetaVerifyPROResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *Id3MetaVerifyPROResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *Id3MetaVerifyPROResponse
	GetStatusCode() *int32
	SetBody(v *Id3MetaVerifyPROResponseBody) *Id3MetaVerifyPROResponse
	GetBody() *Id3MetaVerifyPROResponseBody
}

type Id3MetaVerifyPROResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Id3MetaVerifyPROResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s Id3MetaVerifyPROResponse) String() string {
	return dara.Prettify(s)
}

func (s Id3MetaVerifyPROResponse) GoString() string {
	return s.String()
}

func (s *Id3MetaVerifyPROResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *Id3MetaVerifyPROResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *Id3MetaVerifyPROResponse) GetBody() *Id3MetaVerifyPROResponseBody {
	return s.Body
}

func (s *Id3MetaVerifyPROResponse) SetHeaders(v map[string]*string) *Id3MetaVerifyPROResponse {
	s.Headers = v
	return s
}

func (s *Id3MetaVerifyPROResponse) SetStatusCode(v int32) *Id3MetaVerifyPROResponse {
	s.StatusCode = &v
	return s
}

func (s *Id3MetaVerifyPROResponse) SetBody(v *Id3MetaVerifyPROResponseBody) *Id3MetaVerifyPROResponse {
	s.Body = v
	return s
}

func (s *Id3MetaVerifyPROResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
