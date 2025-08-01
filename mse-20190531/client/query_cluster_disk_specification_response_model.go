// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryClusterDiskSpecificationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryClusterDiskSpecificationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryClusterDiskSpecificationResponse
	GetStatusCode() *int32
	SetBody(v *QueryClusterDiskSpecificationResponseBody) *QueryClusterDiskSpecificationResponse
	GetBody() *QueryClusterDiskSpecificationResponseBody
}

type QueryClusterDiskSpecificationResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryClusterDiskSpecificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryClusterDiskSpecificationResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryClusterDiskSpecificationResponse) GoString() string {
	return s.String()
}

func (s *QueryClusterDiskSpecificationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryClusterDiskSpecificationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryClusterDiskSpecificationResponse) GetBody() *QueryClusterDiskSpecificationResponseBody {
	return s.Body
}

func (s *QueryClusterDiskSpecificationResponse) SetHeaders(v map[string]*string) *QueryClusterDiskSpecificationResponse {
	s.Headers = v
	return s
}

func (s *QueryClusterDiskSpecificationResponse) SetStatusCode(v int32) *QueryClusterDiskSpecificationResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryClusterDiskSpecificationResponse) SetBody(v *QueryClusterDiskSpecificationResponseBody) *QueryClusterDiskSpecificationResponse {
	s.Body = v
	return s
}

func (s *QueryClusterDiskSpecificationResponse) Validate() error {
	return dara.Validate(s)
}
