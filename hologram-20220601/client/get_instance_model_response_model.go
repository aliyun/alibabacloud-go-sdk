// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceModelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceModelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceModelResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceModelResponseBody) *GetInstanceModelResponse
	GetBody() *GetInstanceModelResponseBody
}

type GetInstanceModelResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceModelResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceModelResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceModelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceModelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceModelResponse) GetBody() *GetInstanceModelResponseBody {
	return s.Body
}

func (s *GetInstanceModelResponse) SetHeaders(v map[string]*string) *GetInstanceModelResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceModelResponse) SetStatusCode(v int32) *GetInstanceModelResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceModelResponse) SetBody(v *GetInstanceModelResponseBody) *GetInstanceModelResponse {
	s.Body = v
	return s
}

func (s *GetInstanceModelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
