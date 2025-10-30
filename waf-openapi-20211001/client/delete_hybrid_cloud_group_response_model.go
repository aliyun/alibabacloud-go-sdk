// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHybridCloudGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteHybridCloudGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteHybridCloudGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteHybridCloudGroupResponseBody) *DeleteHybridCloudGroupResponse
	GetBody() *DeleteHybridCloudGroupResponseBody
}

type DeleteHybridCloudGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHybridCloudGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHybridCloudGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteHybridCloudGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteHybridCloudGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteHybridCloudGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteHybridCloudGroupResponse) GetBody() *DeleteHybridCloudGroupResponseBody {
	return s.Body
}

func (s *DeleteHybridCloudGroupResponse) SetHeaders(v map[string]*string) *DeleteHybridCloudGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteHybridCloudGroupResponse) SetStatusCode(v int32) *DeleteHybridCloudGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHybridCloudGroupResponse) SetBody(v *DeleteHybridCloudGroupResponseBody) *DeleteHybridCloudGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteHybridCloudGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
