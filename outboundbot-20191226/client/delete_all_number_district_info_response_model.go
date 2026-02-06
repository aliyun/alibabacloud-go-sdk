// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAllNumberDistrictInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAllNumberDistrictInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAllNumberDistrictInfoResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAllNumberDistrictInfoResponseBody) *DeleteAllNumberDistrictInfoResponse
	GetBody() *DeleteAllNumberDistrictInfoResponseBody
}

type DeleteAllNumberDistrictInfoResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAllNumberDistrictInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAllNumberDistrictInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAllNumberDistrictInfoResponse) GoString() string {
	return s.String()
}

func (s *DeleteAllNumberDistrictInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAllNumberDistrictInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAllNumberDistrictInfoResponse) GetBody() *DeleteAllNumberDistrictInfoResponseBody {
	return s.Body
}

func (s *DeleteAllNumberDistrictInfoResponse) SetHeaders(v map[string]*string) *DeleteAllNumberDistrictInfoResponse {
	s.Headers = v
	return s
}

func (s *DeleteAllNumberDistrictInfoResponse) SetStatusCode(v int32) *DeleteAllNumberDistrictInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAllNumberDistrictInfoResponse) SetBody(v *DeleteAllNumberDistrictInfoResponseBody) *DeleteAllNumberDistrictInfoResponse {
	s.Body = v
	return s
}

func (s *DeleteAllNumberDistrictInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
