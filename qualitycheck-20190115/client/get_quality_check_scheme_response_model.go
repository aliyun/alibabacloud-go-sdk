// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityCheckSchemeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetQualityCheckSchemeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetQualityCheckSchemeResponse
	GetStatusCode() *int32
	SetBody(v *GetQualityCheckSchemeResponseBody) *GetQualityCheckSchemeResponse
	GetBody() *GetQualityCheckSchemeResponseBody
}

type GetQualityCheckSchemeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQualityCheckSchemeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQualityCheckSchemeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetQualityCheckSchemeResponse) GoString() string {
	return s.String()
}

func (s *GetQualityCheckSchemeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetQualityCheckSchemeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetQualityCheckSchemeResponse) GetBody() *GetQualityCheckSchemeResponseBody {
	return s.Body
}

func (s *GetQualityCheckSchemeResponse) SetHeaders(v map[string]*string) *GetQualityCheckSchemeResponse {
	s.Headers = v
	return s
}

func (s *GetQualityCheckSchemeResponse) SetStatusCode(v int32) *GetQualityCheckSchemeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQualityCheckSchemeResponse) SetBody(v *GetQualityCheckSchemeResponseBody) *GetQualityCheckSchemeResponse {
	s.Body = v
	return s
}

func (s *GetQualityCheckSchemeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
