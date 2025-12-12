// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataSetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDataSetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDataSetResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDataSetResponseBody) *DeleteDataSetResponse
	GetBody() *DeleteDataSetResponseBody
}

type DeleteDataSetResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDataSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDataSetResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataSetResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataSetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDataSetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDataSetResponse) GetBody() *DeleteDataSetResponseBody {
	return s.Body
}

func (s *DeleteDataSetResponse) SetHeaders(v map[string]*string) *DeleteDataSetResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataSetResponse) SetStatusCode(v int32) *DeleteDataSetResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataSetResponse) SetBody(v *DeleteDataSetResponseBody) *DeleteDataSetResponse {
	s.Body = v
	return s
}

func (s *DeleteDataSetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
