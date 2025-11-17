// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLabelTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLabelTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLabelTableResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLabelTableResponseBody) *DeleteLabelTableResponse
	GetBody() *DeleteLabelTableResponseBody
}

type DeleteLabelTableResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLabelTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLabelTableResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLabelTableResponse) GoString() string {
	return s.String()
}

func (s *DeleteLabelTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLabelTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLabelTableResponse) GetBody() *DeleteLabelTableResponseBody {
	return s.Body
}

func (s *DeleteLabelTableResponse) SetHeaders(v map[string]*string) *DeleteLabelTableResponse {
	s.Headers = v
	return s
}

func (s *DeleteLabelTableResponse) SetStatusCode(v int32) *DeleteLabelTableResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLabelTableResponse) SetBody(v *DeleteLabelTableResponseBody) *DeleteLabelTableResponse {
	s.Body = v
	return s
}

func (s *DeleteLabelTableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
