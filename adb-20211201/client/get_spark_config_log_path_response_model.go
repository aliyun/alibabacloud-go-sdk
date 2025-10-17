// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSparkConfigLogPathResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSparkConfigLogPathResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSparkConfigLogPathResponse
	GetStatusCode() *int32
	SetBody(v *GetSparkConfigLogPathResponseBody) *GetSparkConfigLogPathResponse
	GetBody() *GetSparkConfigLogPathResponseBody
}

type GetSparkConfigLogPathResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSparkConfigLogPathResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSparkConfigLogPathResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSparkConfigLogPathResponse) GoString() string {
	return s.String()
}

func (s *GetSparkConfigLogPathResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSparkConfigLogPathResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSparkConfigLogPathResponse) GetBody() *GetSparkConfigLogPathResponseBody {
	return s.Body
}

func (s *GetSparkConfigLogPathResponse) SetHeaders(v map[string]*string) *GetSparkConfigLogPathResponse {
	s.Headers = v
	return s
}

func (s *GetSparkConfigLogPathResponse) SetStatusCode(v int32) *GetSparkConfigLogPathResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSparkConfigLogPathResponse) SetBody(v *GetSparkConfigLogPathResponseBody) *GetSparkConfigLogPathResponse {
	s.Body = v
	return s
}

func (s *GetSparkConfigLogPathResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
