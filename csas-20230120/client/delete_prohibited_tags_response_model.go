// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProhibitedTagsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteProhibitedTagsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteProhibitedTagsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteProhibitedTagsResponseBody) *DeleteProhibitedTagsResponse
	GetBody() *DeleteProhibitedTagsResponseBody
}

type DeleteProhibitedTagsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteProhibitedTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteProhibitedTagsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteProhibitedTagsResponse) GoString() string {
	return s.String()
}

func (s *DeleteProhibitedTagsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteProhibitedTagsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteProhibitedTagsResponse) GetBody() *DeleteProhibitedTagsResponseBody {
	return s.Body
}

func (s *DeleteProhibitedTagsResponse) SetHeaders(v map[string]*string) *DeleteProhibitedTagsResponse {
	s.Headers = v
	return s
}

func (s *DeleteProhibitedTagsResponse) SetStatusCode(v int32) *DeleteProhibitedTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteProhibitedTagsResponse) SetBody(v *DeleteProhibitedTagsResponseBody) *DeleteProhibitedTagsResponse {
	s.Body = v
	return s
}

func (s *DeleteProhibitedTagsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
