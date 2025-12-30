// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCheckConnectivityJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCheckConnectivityJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCheckConnectivityJobsResponse
	GetStatusCode() *int32
	SetBody(v *GetCheckConnectivityJobsResponseBody) *GetCheckConnectivityJobsResponse
	GetBody() *GetCheckConnectivityJobsResponseBody
}

type GetCheckConnectivityJobsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCheckConnectivityJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCheckConnectivityJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCheckConnectivityJobsResponse) GoString() string {
	return s.String()
}

func (s *GetCheckConnectivityJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCheckConnectivityJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCheckConnectivityJobsResponse) GetBody() *GetCheckConnectivityJobsResponseBody {
	return s.Body
}

func (s *GetCheckConnectivityJobsResponse) SetHeaders(v map[string]*string) *GetCheckConnectivityJobsResponse {
	s.Headers = v
	return s
}

func (s *GetCheckConnectivityJobsResponse) SetStatusCode(v int32) *GetCheckConnectivityJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCheckConnectivityJobsResponse) SetBody(v *GetCheckConnectivityJobsResponseBody) *GetCheckConnectivityJobsResponse {
	s.Body = v
	return s
}

func (s *GetCheckConnectivityJobsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
