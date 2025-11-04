// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDynamicImageJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDynamicImageJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDynamicImageJobResponse
	GetStatusCode() *int32
	SetBody(v *GetDynamicImageJobResponseBody) *GetDynamicImageJobResponse
	GetBody() *GetDynamicImageJobResponseBody
}

type GetDynamicImageJobResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDynamicImageJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDynamicImageJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDynamicImageJobResponse) GoString() string {
	return s.String()
}

func (s *GetDynamicImageJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDynamicImageJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDynamicImageJobResponse) GetBody() *GetDynamicImageJobResponseBody {
	return s.Body
}

func (s *GetDynamicImageJobResponse) SetHeaders(v map[string]*string) *GetDynamicImageJobResponse {
	s.Headers = v
	return s
}

func (s *GetDynamicImageJobResponse) SetStatusCode(v int32) *GetDynamicImageJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDynamicImageJobResponse) SetBody(v *GetDynamicImageJobResponseBody) *GetDynamicImageJobResponse {
	s.Body = v
	return s
}

func (s *GetDynamicImageJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
