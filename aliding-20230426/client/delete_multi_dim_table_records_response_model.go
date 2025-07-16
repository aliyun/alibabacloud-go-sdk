// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMultiDimTableRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMultiDimTableRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMultiDimTableRecordsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMultiDimTableRecordsResponseBody) *DeleteMultiDimTableRecordsResponse
	GetBody() *DeleteMultiDimTableRecordsResponseBody
}

type DeleteMultiDimTableRecordsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMultiDimTableRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMultiDimTableRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMultiDimTableRecordsResponse) GoString() string {
	return s.String()
}

func (s *DeleteMultiDimTableRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMultiDimTableRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMultiDimTableRecordsResponse) GetBody() *DeleteMultiDimTableRecordsResponseBody {
	return s.Body
}

func (s *DeleteMultiDimTableRecordsResponse) SetHeaders(v map[string]*string) *DeleteMultiDimTableRecordsResponse {
	s.Headers = v
	return s
}

func (s *DeleteMultiDimTableRecordsResponse) SetStatusCode(v int32) *DeleteMultiDimTableRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMultiDimTableRecordsResponse) SetBody(v *DeleteMultiDimTableRecordsResponseBody) *DeleteMultiDimTableRecordsResponse {
	s.Body = v
	return s
}

func (s *DeleteMultiDimTableRecordsResponse) Validate() error {
	return dara.Validate(s)
}
