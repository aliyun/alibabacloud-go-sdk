// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceInspectionRunResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceInspectionRunResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceInspectionRunResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceInspectionRunResponseBody) *GetInstanceInspectionRunResponse
	GetBody() *GetInstanceInspectionRunResponseBody
}

type GetInstanceInspectionRunResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceInspectionRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceInspectionRunResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionRunResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionRunResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceInspectionRunResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceInspectionRunResponse) GetBody() *GetInstanceInspectionRunResponseBody {
	return s.Body
}

func (s *GetInstanceInspectionRunResponse) SetHeaders(v map[string]*string) *GetInstanceInspectionRunResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceInspectionRunResponse) SetStatusCode(v int32) *GetInstanceInspectionRunResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceInspectionRunResponse) SetBody(v *GetInstanceInspectionRunResponseBody) *GetInstanceInspectionRunResponse {
	s.Body = v
	return s
}

func (s *GetInstanceInspectionRunResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
