// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceUsageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetResourceUsageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetResourceUsageResponse
	GetStatusCode() *int32
	SetBody(v *GetResourceUsageResponseBody) *GetResourceUsageResponse
	GetBody() *GetResourceUsageResponseBody
}

type GetResourceUsageResponse struct {
	Headers    map[string]*string            `json:"headers" xml:"headers"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceUsageResponse) String() string {
	return dara.Prettify(s)
}

func (s GetResourceUsageResponse) GoString() string {
	return s.String()
}

func (s *GetResourceUsageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetResourceUsageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetResourceUsageResponse) GetBody() *GetResourceUsageResponseBody {
	return s.Body
}

func (s *GetResourceUsageResponse) SetHeaders(v map[string]*string) *GetResourceUsageResponse {
	s.Headers = v
	return s
}

func (s *GetResourceUsageResponse) SetStatusCode(v int32) *GetResourceUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceUsageResponse) SetBody(v *GetResourceUsageResponseBody) *GetResourceUsageResponse {
	s.Body = v
	return s
}

func (s *GetResourceUsageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
