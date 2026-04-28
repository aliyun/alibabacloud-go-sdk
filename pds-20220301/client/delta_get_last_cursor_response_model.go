// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeltaGetLastCursorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeltaGetLastCursorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeltaGetLastCursorResponse
	GetStatusCode() *int32
	SetBody(v *DeltaGetLastCursorResponseBody) *DeltaGetLastCursorResponse
	GetBody() *DeltaGetLastCursorResponseBody
}

type DeltaGetLastCursorResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeltaGetLastCursorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeltaGetLastCursorResponse) String() string {
	return dara.Prettify(s)
}

func (s DeltaGetLastCursorResponse) GoString() string {
	return s.String()
}

func (s *DeltaGetLastCursorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeltaGetLastCursorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeltaGetLastCursorResponse) GetBody() *DeltaGetLastCursorResponseBody {
	return s.Body
}

func (s *DeltaGetLastCursorResponse) SetHeaders(v map[string]*string) *DeltaGetLastCursorResponse {
	s.Headers = v
	return s
}

func (s *DeltaGetLastCursorResponse) SetStatusCode(v int32) *DeltaGetLastCursorResponse {
	s.StatusCode = &v
	return s
}

func (s *DeltaGetLastCursorResponse) SetBody(v *DeltaGetLastCursorResponseBody) *DeltaGetLastCursorResponse {
	s.Body = v
	return s
}

func (s *DeltaGetLastCursorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
