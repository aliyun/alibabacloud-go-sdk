// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDesignateWorkersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DesignateWorkersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DesignateWorkersResponse
	GetStatusCode() *int32
	SetBody(v *DesignateWorkersResponseBody) *DesignateWorkersResponse
	GetBody() *DesignateWorkersResponseBody
}

type DesignateWorkersResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DesignateWorkersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DesignateWorkersResponse) String() string {
	return dara.Prettify(s)
}

func (s DesignateWorkersResponse) GoString() string {
	return s.String()
}

func (s *DesignateWorkersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DesignateWorkersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DesignateWorkersResponse) GetBody() *DesignateWorkersResponseBody {
	return s.Body
}

func (s *DesignateWorkersResponse) SetHeaders(v map[string]*string) *DesignateWorkersResponse {
	s.Headers = v
	return s
}

func (s *DesignateWorkersResponse) SetStatusCode(v int32) *DesignateWorkersResponse {
	s.StatusCode = &v
	return s
}

func (s *DesignateWorkersResponse) SetBody(v *DesignateWorkersResponseBody) *DesignateWorkersResponse {
	s.Body = v
	return s
}

func (s *DesignateWorkersResponse) Validate() error {
	return dara.Validate(s)
}
