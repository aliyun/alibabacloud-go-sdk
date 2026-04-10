// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUsageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUsageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUsageResponse
	GetStatusCode() *int32
	SetBody(v *GetUsageResponseBody) *GetUsageResponse
	GetBody() *GetUsageResponseBody
}

type GetUsageResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUsageResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUsageResponse) GoString() string {
	return s.String()
}

func (s *GetUsageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUsageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUsageResponse) GetBody() *GetUsageResponseBody {
	return s.Body
}

func (s *GetUsageResponse) SetHeaders(v map[string]*string) *GetUsageResponse {
	s.Headers = v
	return s
}

func (s *GetUsageResponse) SetStatusCode(v int32) *GetUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUsageResponse) SetBody(v *GetUsageResponseBody) *GetUsageResponse {
	s.Body = v
	return s
}

func (s *GetUsageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
