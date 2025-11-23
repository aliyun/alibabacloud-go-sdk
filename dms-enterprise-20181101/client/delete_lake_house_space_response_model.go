// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLakeHouseSpaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLakeHouseSpaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLakeHouseSpaceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLakeHouseSpaceResponseBody) *DeleteLakeHouseSpaceResponse
	GetBody() *DeleteLakeHouseSpaceResponseBody
}

type DeleteLakeHouseSpaceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLakeHouseSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLakeHouseSpaceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLakeHouseSpaceResponse) GoString() string {
	return s.String()
}

func (s *DeleteLakeHouseSpaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLakeHouseSpaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLakeHouseSpaceResponse) GetBody() *DeleteLakeHouseSpaceResponseBody {
	return s.Body
}

func (s *DeleteLakeHouseSpaceResponse) SetHeaders(v map[string]*string) *DeleteLakeHouseSpaceResponse {
	s.Headers = v
	return s
}

func (s *DeleteLakeHouseSpaceResponse) SetStatusCode(v int32) *DeleteLakeHouseSpaceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLakeHouseSpaceResponse) SetBody(v *DeleteLakeHouseSpaceResponseBody) *DeleteLakeHouseSpaceResponse {
	s.Body = v
	return s
}

func (s *DeleteLakeHouseSpaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
