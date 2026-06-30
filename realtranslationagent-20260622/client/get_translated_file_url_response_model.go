// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTranslatedFileUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTranslatedFileUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTranslatedFileUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetTranslatedFileUrlResponseBody) *GetTranslatedFileUrlResponse
	GetBody() *GetTranslatedFileUrlResponseBody
}

type GetTranslatedFileUrlResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTranslatedFileUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTranslatedFileUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTranslatedFileUrlResponse) GoString() string {
	return s.String()
}

func (s *GetTranslatedFileUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTranslatedFileUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTranslatedFileUrlResponse) GetBody() *GetTranslatedFileUrlResponseBody {
	return s.Body
}

func (s *GetTranslatedFileUrlResponse) SetHeaders(v map[string]*string) *GetTranslatedFileUrlResponse {
	s.Headers = v
	return s
}

func (s *GetTranslatedFileUrlResponse) SetStatusCode(v int32) *GetTranslatedFileUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTranslatedFileUrlResponse) SetBody(v *GetTranslatedFileUrlResponseBody) *GetTranslatedFileUrlResponse {
	s.Body = v
	return s
}

func (s *GetTranslatedFileUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
