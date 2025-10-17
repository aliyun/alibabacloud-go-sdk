// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSparkAppStateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSparkAppStateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSparkAppStateResponse
	GetStatusCode() *int32
	SetBody(v *GetSparkAppStateResponseBody) *GetSparkAppStateResponse
	GetBody() *GetSparkAppStateResponseBody
}

type GetSparkAppStateResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSparkAppStateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSparkAppStateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSparkAppStateResponse) GoString() string {
	return s.String()
}

func (s *GetSparkAppStateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSparkAppStateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSparkAppStateResponse) GetBody() *GetSparkAppStateResponseBody {
	return s.Body
}

func (s *GetSparkAppStateResponse) SetHeaders(v map[string]*string) *GetSparkAppStateResponse {
	s.Headers = v
	return s
}

func (s *GetSparkAppStateResponse) SetStatusCode(v int32) *GetSparkAppStateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSparkAppStateResponse) SetBody(v *GetSparkAppStateResponseBody) *GetSparkAppStateResponse {
	s.Body = v
	return s
}

func (s *GetSparkAppStateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
