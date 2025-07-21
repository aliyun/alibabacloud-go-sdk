// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppOtaVersionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAppOtaVersionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAppOtaVersionsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAppOtaVersionsResponseBody) *DeleteAppOtaVersionsResponse
	GetBody() *DeleteAppOtaVersionsResponseBody
}

type DeleteAppOtaVersionsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAppOtaVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAppOtaVersionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppOtaVersionsResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppOtaVersionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAppOtaVersionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAppOtaVersionsResponse) GetBody() *DeleteAppOtaVersionsResponseBody {
	return s.Body
}

func (s *DeleteAppOtaVersionsResponse) SetHeaders(v map[string]*string) *DeleteAppOtaVersionsResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppOtaVersionsResponse) SetStatusCode(v int32) *DeleteAppOtaVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAppOtaVersionsResponse) SetBody(v *DeleteAppOtaVersionsResponseBody) *DeleteAppOtaVersionsResponse {
	s.Body = v
	return s
}

func (s *DeleteAppOtaVersionsResponse) Validate() error {
	return dara.Validate(s)
}
