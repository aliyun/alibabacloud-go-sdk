// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCheckConnectivityJobByJobIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCheckConnectivityJobByJobIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCheckConnectivityJobByJobIdResponse
	GetStatusCode() *int32
	SetBody(v *GetCheckConnectivityJobByJobIdResponseBody) *GetCheckConnectivityJobByJobIdResponse
	GetBody() *GetCheckConnectivityJobByJobIdResponseBody
}

type GetCheckConnectivityJobByJobIdResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCheckConnectivityJobByJobIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCheckConnectivityJobByJobIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCheckConnectivityJobByJobIdResponse) GoString() string {
	return s.String()
}

func (s *GetCheckConnectivityJobByJobIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCheckConnectivityJobByJobIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCheckConnectivityJobByJobIdResponse) GetBody() *GetCheckConnectivityJobByJobIdResponseBody {
	return s.Body
}

func (s *GetCheckConnectivityJobByJobIdResponse) SetHeaders(v map[string]*string) *GetCheckConnectivityJobByJobIdResponse {
	s.Headers = v
	return s
}

func (s *GetCheckConnectivityJobByJobIdResponse) SetStatusCode(v int32) *GetCheckConnectivityJobByJobIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCheckConnectivityJobByJobIdResponse) SetBody(v *GetCheckConnectivityJobByJobIdResponseBody) *GetCheckConnectivityJobByJobIdResponse {
	s.Body = v
	return s
}

func (s *GetCheckConnectivityJobByJobIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
