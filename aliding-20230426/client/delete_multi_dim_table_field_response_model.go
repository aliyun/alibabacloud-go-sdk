// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMultiDimTableFieldResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMultiDimTableFieldResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMultiDimTableFieldResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMultiDimTableFieldResponseBody) *DeleteMultiDimTableFieldResponse
	GetBody() *DeleteMultiDimTableFieldResponseBody
}

type DeleteMultiDimTableFieldResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMultiDimTableFieldResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMultiDimTableFieldResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMultiDimTableFieldResponse) GoString() string {
	return s.String()
}

func (s *DeleteMultiDimTableFieldResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMultiDimTableFieldResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMultiDimTableFieldResponse) GetBody() *DeleteMultiDimTableFieldResponseBody {
	return s.Body
}

func (s *DeleteMultiDimTableFieldResponse) SetHeaders(v map[string]*string) *DeleteMultiDimTableFieldResponse {
	s.Headers = v
	return s
}

func (s *DeleteMultiDimTableFieldResponse) SetStatusCode(v int32) *DeleteMultiDimTableFieldResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMultiDimTableFieldResponse) SetBody(v *DeleteMultiDimTableFieldResponseBody) *DeleteMultiDimTableFieldResponse {
	s.Body = v
	return s
}

func (s *DeleteMultiDimTableFieldResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
