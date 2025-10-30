// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPhysicalNodeByOutputNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPhysicalNodeByOutputNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPhysicalNodeByOutputNameResponse
	GetStatusCode() *int32
	SetBody(v *GetPhysicalNodeByOutputNameResponseBody) *GetPhysicalNodeByOutputNameResponse
	GetBody() *GetPhysicalNodeByOutputNameResponseBody
}

type GetPhysicalNodeByOutputNameResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPhysicalNodeByOutputNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPhysicalNodeByOutputNameResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPhysicalNodeByOutputNameResponse) GoString() string {
	return s.String()
}

func (s *GetPhysicalNodeByOutputNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPhysicalNodeByOutputNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPhysicalNodeByOutputNameResponse) GetBody() *GetPhysicalNodeByOutputNameResponseBody {
	return s.Body
}

func (s *GetPhysicalNodeByOutputNameResponse) SetHeaders(v map[string]*string) *GetPhysicalNodeByOutputNameResponse {
	s.Headers = v
	return s
}

func (s *GetPhysicalNodeByOutputNameResponse) SetStatusCode(v int32) *GetPhysicalNodeByOutputNameResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPhysicalNodeByOutputNameResponse) SetBody(v *GetPhysicalNodeByOutputNameResponseBody) *GetPhysicalNodeByOutputNameResponse {
	s.Body = v
	return s
}

func (s *GetPhysicalNodeByOutputNameResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
