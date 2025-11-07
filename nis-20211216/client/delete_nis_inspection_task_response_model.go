// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNisInspectionTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteNisInspectionTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteNisInspectionTaskResponse
	GetStatusCode() *int32
	SetBody(v *DeleteNisInspectionTaskResponseBody) *DeleteNisInspectionTaskResponse
	GetBody() *DeleteNisInspectionTaskResponseBody
}

type DeleteNisInspectionTaskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNisInspectionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNisInspectionTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteNisInspectionTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteNisInspectionTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteNisInspectionTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteNisInspectionTaskResponse) GetBody() *DeleteNisInspectionTaskResponseBody {
	return s.Body
}

func (s *DeleteNisInspectionTaskResponse) SetHeaders(v map[string]*string) *DeleteNisInspectionTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteNisInspectionTaskResponse) SetStatusCode(v int32) *DeleteNisInspectionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNisInspectionTaskResponse) SetBody(v *DeleteNisInspectionTaskResponseBody) *DeleteNisInspectionTaskResponse {
	s.Body = v
	return s
}

func (s *DeleteNisInspectionTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
