// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetQualityResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetQualityResultResponse
	GetStatusCode() *int32
	SetBody(v *GetQualityResultResponseBody) *GetQualityResultResponse
	GetBody() *GetQualityResultResponseBody
}

type GetQualityResultResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQualityResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQualityResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetQualityResultResponse) GoString() string {
	return s.String()
}

func (s *GetQualityResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetQualityResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetQualityResultResponse) GetBody() *GetQualityResultResponseBody {
	return s.Body
}

func (s *GetQualityResultResponse) SetHeaders(v map[string]*string) *GetQualityResultResponse {
	s.Headers = v
	return s
}

func (s *GetQualityResultResponse) SetStatusCode(v int32) *GetQualityResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQualityResultResponse) SetBody(v *GetQualityResultResponseBody) *GetQualityResultResponse {
	s.Body = v
	return s
}

func (s *GetQualityResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
