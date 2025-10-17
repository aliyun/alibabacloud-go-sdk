// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSparkAppInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSparkAppInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSparkAppInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetSparkAppInfoResponseBody) *GetSparkAppInfoResponse
	GetBody() *GetSparkAppInfoResponseBody
}

type GetSparkAppInfoResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSparkAppInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSparkAppInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSparkAppInfoResponse) GoString() string {
	return s.String()
}

func (s *GetSparkAppInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSparkAppInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSparkAppInfoResponse) GetBody() *GetSparkAppInfoResponseBody {
	return s.Body
}

func (s *GetSparkAppInfoResponse) SetHeaders(v map[string]*string) *GetSparkAppInfoResponse {
	s.Headers = v
	return s
}

func (s *GetSparkAppInfoResponse) SetStatusCode(v int32) *GetSparkAppInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSparkAppInfoResponse) SetBody(v *GetSparkAppInfoResponseBody) *GetSparkAppInfoResponse {
	s.Body = v
	return s
}

func (s *GetSparkAppInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
