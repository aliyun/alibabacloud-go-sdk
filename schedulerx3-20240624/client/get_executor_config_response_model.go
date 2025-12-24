// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExecutorConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetExecutorConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetExecutorConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetExecutorConfigResponseBody) *GetExecutorConfigResponse
	GetBody() *GetExecutorConfigResponseBody
}

type GetExecutorConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetExecutorConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetExecutorConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetExecutorConfigResponse) GoString() string {
	return s.String()
}

func (s *GetExecutorConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetExecutorConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetExecutorConfigResponse) GetBody() *GetExecutorConfigResponseBody {
	return s.Body
}

func (s *GetExecutorConfigResponse) SetHeaders(v map[string]*string) *GetExecutorConfigResponse {
	s.Headers = v
	return s
}

func (s *GetExecutorConfigResponse) SetStatusCode(v int32) *GetExecutorConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExecutorConfigResponse) SetBody(v *GetExecutorConfigResponseBody) *GetExecutorConfigResponse {
	s.Body = v
	return s
}

func (s *GetExecutorConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
