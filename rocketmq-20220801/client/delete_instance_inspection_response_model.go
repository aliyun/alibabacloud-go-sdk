// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceInspectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteInstanceInspectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteInstanceInspectionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteInstanceInspectionResponseBody) *DeleteInstanceInspectionResponse
	GetBody() *DeleteInstanceInspectionResponseBody
}

type DeleteInstanceInspectionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteInstanceInspectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteInstanceInspectionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceInspectionResponse) GoString() string {
	return s.String()
}

func (s *DeleteInstanceInspectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteInstanceInspectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteInstanceInspectionResponse) GetBody() *DeleteInstanceInspectionResponseBody {
	return s.Body
}

func (s *DeleteInstanceInspectionResponse) SetHeaders(v map[string]*string) *DeleteInstanceInspectionResponse {
	s.Headers = v
	return s
}

func (s *DeleteInstanceInspectionResponse) SetStatusCode(v int32) *DeleteInstanceInspectionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInstanceInspectionResponse) SetBody(v *DeleteInstanceInspectionResponseBody) *DeleteInstanceInspectionResponse {
	s.Body = v
	return s
}

func (s *DeleteInstanceInspectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
