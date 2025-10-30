// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPhysicalNodeContentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPhysicalNodeContentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPhysicalNodeContentResponse
	GetStatusCode() *int32
	SetBody(v *GetPhysicalNodeContentResponseBody) *GetPhysicalNodeContentResponse
	GetBody() *GetPhysicalNodeContentResponseBody
}

type GetPhysicalNodeContentResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPhysicalNodeContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPhysicalNodeContentResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPhysicalNodeContentResponse) GoString() string {
	return s.String()
}

func (s *GetPhysicalNodeContentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPhysicalNodeContentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPhysicalNodeContentResponse) GetBody() *GetPhysicalNodeContentResponseBody {
	return s.Body
}

func (s *GetPhysicalNodeContentResponse) SetHeaders(v map[string]*string) *GetPhysicalNodeContentResponse {
	s.Headers = v
	return s
}

func (s *GetPhysicalNodeContentResponse) SetStatusCode(v int32) *GetPhysicalNodeContentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPhysicalNodeContentResponse) SetBody(v *GetPhysicalNodeContentResponseBody) *GetPhysicalNodeContentResponse {
	s.Body = v
	return s
}

func (s *GetPhysicalNodeContentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
