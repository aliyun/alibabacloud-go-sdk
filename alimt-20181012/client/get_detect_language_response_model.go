// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDetectLanguageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDetectLanguageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDetectLanguageResponse
	GetStatusCode() *int32
	SetBody(v *GetDetectLanguageResponseBody) *GetDetectLanguageResponse
	GetBody() *GetDetectLanguageResponseBody
}

type GetDetectLanguageResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDetectLanguageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDetectLanguageResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDetectLanguageResponse) GoString() string {
	return s.String()
}

func (s *GetDetectLanguageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDetectLanguageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDetectLanguageResponse) GetBody() *GetDetectLanguageResponseBody {
	return s.Body
}

func (s *GetDetectLanguageResponse) SetHeaders(v map[string]*string) *GetDetectLanguageResponse {
	s.Headers = v
	return s
}

func (s *GetDetectLanguageResponse) SetStatusCode(v int32) *GetDetectLanguageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDetectLanguageResponse) SetBody(v *GetDetectLanguageResponseBody) *GetDetectLanguageResponse {
	s.Body = v
	return s
}

func (s *GetDetectLanguageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
