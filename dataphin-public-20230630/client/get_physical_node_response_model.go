// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPhysicalNodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPhysicalNodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPhysicalNodeResponse
	GetStatusCode() *int32
	SetBody(v *GetPhysicalNodeResponseBody) *GetPhysicalNodeResponse
	GetBody() *GetPhysicalNodeResponseBody
}

type GetPhysicalNodeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPhysicalNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPhysicalNodeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPhysicalNodeResponse) GoString() string {
	return s.String()
}

func (s *GetPhysicalNodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPhysicalNodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPhysicalNodeResponse) GetBody() *GetPhysicalNodeResponseBody {
	return s.Body
}

func (s *GetPhysicalNodeResponse) SetHeaders(v map[string]*string) *GetPhysicalNodeResponse {
	s.Headers = v
	return s
}

func (s *GetPhysicalNodeResponse) SetStatusCode(v int32) *GetPhysicalNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPhysicalNodeResponse) SetBody(v *GetPhysicalNodeResponseBody) *GetPhysicalNodeResponse {
	s.Body = v
	return s
}

func (s *GetPhysicalNodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
