// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachVscFromMountPointsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachVscFromMountPointsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachVscFromMountPointsResponse
	GetStatusCode() *int32
	SetBody(v *DetachVscFromMountPointsResponseBody) *DetachVscFromMountPointsResponse
	GetBody() *DetachVscFromMountPointsResponseBody
}

type DetachVscFromMountPointsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachVscFromMountPointsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachVscFromMountPointsResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachVscFromMountPointsResponse) GoString() string {
	return s.String()
}

func (s *DetachVscFromMountPointsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachVscFromMountPointsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachVscFromMountPointsResponse) GetBody() *DetachVscFromMountPointsResponseBody {
	return s.Body
}

func (s *DetachVscFromMountPointsResponse) SetHeaders(v map[string]*string) *DetachVscFromMountPointsResponse {
	s.Headers = v
	return s
}

func (s *DetachVscFromMountPointsResponse) SetStatusCode(v int32) *DetachVscFromMountPointsResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachVscFromMountPointsResponse) SetBody(v *DetachVscFromMountPointsResponseBody) *DetachVscFromMountPointsResponse {
	s.Body = v
	return s
}

func (s *DetachVscFromMountPointsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
