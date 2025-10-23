// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocExtractionResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDocExtractionResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDocExtractionResultResponse
	GetStatusCode() *int32
	SetBody(v *GetDocExtractionResultResponseBody) *GetDocExtractionResultResponse
	GetBody() *GetDocExtractionResultResponseBody
}

type GetDocExtractionResultResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDocExtractionResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDocExtractionResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDocExtractionResultResponse) GoString() string {
	return s.String()
}

func (s *GetDocExtractionResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDocExtractionResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDocExtractionResultResponse) GetBody() *GetDocExtractionResultResponseBody {
	return s.Body
}

func (s *GetDocExtractionResultResponse) SetHeaders(v map[string]*string) *GetDocExtractionResultResponse {
	s.Headers = v
	return s
}

func (s *GetDocExtractionResultResponse) SetStatusCode(v int32) *GetDocExtractionResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDocExtractionResultResponse) SetBody(v *GetDocExtractionResultResponseBody) *GetDocExtractionResultResponse {
	s.Body = v
	return s
}

func (s *GetDocExtractionResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
