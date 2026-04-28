// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceInspectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateInstanceInspectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateInstanceInspectionResponse
	GetStatusCode() *int32
	SetBody(v *CreateInstanceInspectionResponseBody) *CreateInstanceInspectionResponse
	GetBody() *CreateInstanceInspectionResponseBody
}

type CreateInstanceInspectionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInstanceInspectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateInstanceInspectionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceInspectionResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceInspectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateInstanceInspectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateInstanceInspectionResponse) GetBody() *CreateInstanceInspectionResponseBody {
	return s.Body
}

func (s *CreateInstanceInspectionResponse) SetHeaders(v map[string]*string) *CreateInstanceInspectionResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceInspectionResponse) SetStatusCode(v int32) *CreateInstanceInspectionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstanceInspectionResponse) SetBody(v *CreateInstanceInspectionResponseBody) *CreateInstanceInspectionResponse {
	s.Body = v
	return s
}

func (s *CreateInstanceInspectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
