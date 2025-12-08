// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchAddFacesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchAddFacesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchAddFacesResponse
	GetStatusCode() *int32
	SetBody(v *BatchAddFacesResponseBody) *BatchAddFacesResponse
	GetBody() *BatchAddFacesResponseBody
}

type BatchAddFacesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchAddFacesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchAddFacesResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchAddFacesResponse) GoString() string {
	return s.String()
}

func (s *BatchAddFacesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchAddFacesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchAddFacesResponse) GetBody() *BatchAddFacesResponseBody {
	return s.Body
}

func (s *BatchAddFacesResponse) SetHeaders(v map[string]*string) *BatchAddFacesResponse {
	s.Headers = v
	return s
}

func (s *BatchAddFacesResponse) SetStatusCode(v int32) *BatchAddFacesResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchAddFacesResponse) SetBody(v *BatchAddFacesResponseBody) *BatchAddFacesResponse {
	s.Body = v
	return s
}

func (s *BatchAddFacesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
