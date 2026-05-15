// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityProjectLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetQualityProjectLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetQualityProjectLogResponse
	GetStatusCode() *int32
	SetBody(v *GetQualityProjectLogResponseBody) *GetQualityProjectLogResponse
	GetBody() *GetQualityProjectLogResponseBody
}

type GetQualityProjectLogResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQualityProjectLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQualityProjectLogResponse) String() string {
	return dara.Prettify(s)
}

func (s GetQualityProjectLogResponse) GoString() string {
	return s.String()
}

func (s *GetQualityProjectLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetQualityProjectLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetQualityProjectLogResponse) GetBody() *GetQualityProjectLogResponseBody {
	return s.Body
}

func (s *GetQualityProjectLogResponse) SetHeaders(v map[string]*string) *GetQualityProjectLogResponse {
	s.Headers = v
	return s
}

func (s *GetQualityProjectLogResponse) SetStatusCode(v int32) *GetQualityProjectLogResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQualityProjectLogResponse) SetBody(v *GetQualityProjectLogResponseBody) *GetQualityProjectLogResponse {
	s.Body = v
	return s
}

func (s *GetQualityProjectLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
