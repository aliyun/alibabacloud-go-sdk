// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceInspectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceInspectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceInspectionResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceInspectionResponseBody) *GetInstanceInspectionResponse
	GetBody() *GetInstanceInspectionResponseBody
}

type GetInstanceInspectionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceInspectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceInspectionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceInspectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceInspectionResponse) GetBody() *GetInstanceInspectionResponseBody {
	return s.Body
}

func (s *GetInstanceInspectionResponse) SetHeaders(v map[string]*string) *GetInstanceInspectionResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceInspectionResponse) SetStatusCode(v int32) *GetInstanceInspectionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceInspectionResponse) SetBody(v *GetInstanceInspectionResponseBody) *GetInstanceInspectionResponse {
	s.Body = v
	return s
}

func (s *GetInstanceInspectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
