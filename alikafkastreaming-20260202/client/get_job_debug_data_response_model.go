// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobDebugDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetJobDebugDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetJobDebugDataResponse
	GetStatusCode() *int32
	SetBody(v *GetJobDebugDataResponseBody) *GetJobDebugDataResponse
	GetBody() *GetJobDebugDataResponseBody
}

type GetJobDebugDataResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetJobDebugDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetJobDebugDataResponse) String() string {
	return dara.Prettify(s)
}

func (s GetJobDebugDataResponse) GoString() string {
	return s.String()
}

func (s *GetJobDebugDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetJobDebugDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetJobDebugDataResponse) GetBody() *GetJobDebugDataResponseBody {
	return s.Body
}

func (s *GetJobDebugDataResponse) SetHeaders(v map[string]*string) *GetJobDebugDataResponse {
	s.Headers = v
	return s
}

func (s *GetJobDebugDataResponse) SetStatusCode(v int32) *GetJobDebugDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobDebugDataResponse) SetBody(v *GetJobDebugDataResponseBody) *GetJobDebugDataResponse {
	s.Body = v
	return s
}

func (s *GetJobDebugDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
