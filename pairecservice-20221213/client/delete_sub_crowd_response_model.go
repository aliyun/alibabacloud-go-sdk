// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSubCrowdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSubCrowdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSubCrowdResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSubCrowdResponseBody) *DeleteSubCrowdResponse
	GetBody() *DeleteSubCrowdResponseBody
}

type DeleteSubCrowdResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSubCrowdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSubCrowdResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSubCrowdResponse) GoString() string {
	return s.String()
}

func (s *DeleteSubCrowdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSubCrowdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSubCrowdResponse) GetBody() *DeleteSubCrowdResponseBody {
	return s.Body
}

func (s *DeleteSubCrowdResponse) SetHeaders(v map[string]*string) *DeleteSubCrowdResponse {
	s.Headers = v
	return s
}

func (s *DeleteSubCrowdResponse) SetStatusCode(v int32) *DeleteSubCrowdResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSubCrowdResponse) SetBody(v *DeleteSubCrowdResponseBody) *DeleteSubCrowdResponse {
	s.Body = v
	return s
}

func (s *DeleteSubCrowdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
