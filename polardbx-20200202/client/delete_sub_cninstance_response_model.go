// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSubCNInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSubCNInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSubCNInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSubCNInstanceResponseBody) *DeleteSubCNInstanceResponse
	GetBody() *DeleteSubCNInstanceResponseBody
}

type DeleteSubCNInstanceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSubCNInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSubCNInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSubCNInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteSubCNInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSubCNInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSubCNInstanceResponse) GetBody() *DeleteSubCNInstanceResponseBody {
	return s.Body
}

func (s *DeleteSubCNInstanceResponse) SetHeaders(v map[string]*string) *DeleteSubCNInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteSubCNInstanceResponse) SetStatusCode(v int32) *DeleteSubCNInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSubCNInstanceResponse) SetBody(v *DeleteSubCNInstanceResponseBody) *DeleteSubCNInstanceResponse {
	s.Body = v
	return s
}

func (s *DeleteSubCNInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
