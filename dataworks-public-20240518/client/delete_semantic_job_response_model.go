// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSemanticJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSemanticJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSemanticJobResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSemanticJobResponseBody) *DeleteSemanticJobResponse
	GetBody() *DeleteSemanticJobResponseBody
}

type DeleteSemanticJobResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSemanticJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSemanticJobResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSemanticJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteSemanticJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSemanticJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSemanticJobResponse) GetBody() *DeleteSemanticJobResponseBody {
	return s.Body
}

func (s *DeleteSemanticJobResponse) SetHeaders(v map[string]*string) *DeleteSemanticJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteSemanticJobResponse) SetStatusCode(v int32) *DeleteSemanticJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSemanticJobResponse) SetBody(v *DeleteSemanticJobResponseBody) *DeleteSemanticJobResponse {
	s.Body = v
	return s
}

func (s *DeleteSemanticJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
