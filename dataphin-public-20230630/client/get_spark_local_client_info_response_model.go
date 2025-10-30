// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSparkLocalClientInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSparkLocalClientInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSparkLocalClientInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetSparkLocalClientInfoResponseBody) *GetSparkLocalClientInfoResponse
	GetBody() *GetSparkLocalClientInfoResponseBody
}

type GetSparkLocalClientInfoResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSparkLocalClientInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSparkLocalClientInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSparkLocalClientInfoResponse) GoString() string {
	return s.String()
}

func (s *GetSparkLocalClientInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSparkLocalClientInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSparkLocalClientInfoResponse) GetBody() *GetSparkLocalClientInfoResponseBody {
	return s.Body
}

func (s *GetSparkLocalClientInfoResponse) SetHeaders(v map[string]*string) *GetSparkLocalClientInfoResponse {
	s.Headers = v
	return s
}

func (s *GetSparkLocalClientInfoResponse) SetStatusCode(v int32) *GetSparkLocalClientInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSparkLocalClientInfoResponse) SetBody(v *GetSparkLocalClientInfoResponseBody) *GetSparkLocalClientInfoResponse {
	s.Body = v
	return s
}

func (s *GetSparkLocalClientInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
