// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomBlockRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCustomBlockRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCustomBlockRecordResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCustomBlockRecordResponseBody) *DeleteCustomBlockRecordResponse
	GetBody() *DeleteCustomBlockRecordResponseBody
}

type DeleteCustomBlockRecordResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustomBlockRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustomBlockRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomBlockRecordResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomBlockRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCustomBlockRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCustomBlockRecordResponse) GetBody() *DeleteCustomBlockRecordResponseBody {
	return s.Body
}

func (s *DeleteCustomBlockRecordResponse) SetHeaders(v map[string]*string) *DeleteCustomBlockRecordResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomBlockRecordResponse) SetStatusCode(v int32) *DeleteCustomBlockRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomBlockRecordResponse) SetBody(v *DeleteCustomBlockRecordResponseBody) *DeleteCustomBlockRecordResponse {
	s.Body = v
	return s
}

func (s *DeleteCustomBlockRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
