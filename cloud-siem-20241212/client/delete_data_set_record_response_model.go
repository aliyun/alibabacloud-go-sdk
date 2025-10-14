// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataSetRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDataSetRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDataSetRecordResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDataSetRecordResponseBody) *DeleteDataSetRecordResponse
	GetBody() *DeleteDataSetRecordResponseBody
}

type DeleteDataSetRecordResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDataSetRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDataSetRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataSetRecordResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataSetRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDataSetRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDataSetRecordResponse) GetBody() *DeleteDataSetRecordResponseBody {
	return s.Body
}

func (s *DeleteDataSetRecordResponse) SetHeaders(v map[string]*string) *DeleteDataSetRecordResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataSetRecordResponse) SetStatusCode(v int32) *DeleteDataSetRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataSetRecordResponse) SetBody(v *DeleteDataSetRecordResponseBody) *DeleteDataSetRecordResponse {
	s.Body = v
	return s
}

func (s *DeleteDataSetRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
