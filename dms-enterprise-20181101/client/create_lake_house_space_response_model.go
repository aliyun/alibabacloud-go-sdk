// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLakeHouseSpaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLakeHouseSpaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLakeHouseSpaceResponse
	GetStatusCode() *int32
	SetBody(v *CreateLakeHouseSpaceResponseBody) *CreateLakeHouseSpaceResponse
	GetBody() *CreateLakeHouseSpaceResponseBody
}

type CreateLakeHouseSpaceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLakeHouseSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLakeHouseSpaceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLakeHouseSpaceResponse) GoString() string {
	return s.String()
}

func (s *CreateLakeHouseSpaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLakeHouseSpaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLakeHouseSpaceResponse) GetBody() *CreateLakeHouseSpaceResponseBody {
	return s.Body
}

func (s *CreateLakeHouseSpaceResponse) SetHeaders(v map[string]*string) *CreateLakeHouseSpaceResponse {
	s.Headers = v
	return s
}

func (s *CreateLakeHouseSpaceResponse) SetStatusCode(v int32) *CreateLakeHouseSpaceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLakeHouseSpaceResponse) SetBody(v *CreateLakeHouseSpaceResponseBody) *CreateLakeHouseSpaceResponse {
	s.Body = v
	return s
}

func (s *CreateLakeHouseSpaceResponse) Validate() error {
	return dara.Validate(s)
}
