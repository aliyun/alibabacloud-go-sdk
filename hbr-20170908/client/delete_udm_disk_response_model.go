// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUdmDiskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteUdmDiskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteUdmDiskResponse
	GetStatusCode() *int32
	SetBody(v *DeleteUdmDiskResponseBody) *DeleteUdmDiskResponse
	GetBody() *DeleteUdmDiskResponseBody
}

type DeleteUdmDiskResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUdmDiskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUdmDiskResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteUdmDiskResponse) GoString() string {
	return s.String()
}

func (s *DeleteUdmDiskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteUdmDiskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteUdmDiskResponse) GetBody() *DeleteUdmDiskResponseBody {
	return s.Body
}

func (s *DeleteUdmDiskResponse) SetHeaders(v map[string]*string) *DeleteUdmDiskResponse {
	s.Headers = v
	return s
}

func (s *DeleteUdmDiskResponse) SetStatusCode(v int32) *DeleteUdmDiskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUdmDiskResponse) SetBody(v *DeleteUdmDiskResponseBody) *DeleteUdmDiskResponse {
	s.Body = v
	return s
}

func (s *DeleteUdmDiskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
