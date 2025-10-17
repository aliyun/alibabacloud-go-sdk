// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApsDatasoureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteApsDatasoureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteApsDatasoureResponse
	GetStatusCode() *int32
	SetBody(v *DeleteApsDatasoureResponseBody) *DeleteApsDatasoureResponse
	GetBody() *DeleteApsDatasoureResponseBody
}

type DeleteApsDatasoureResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteApsDatasoureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteApsDatasoureResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteApsDatasoureResponse) GoString() string {
	return s.String()
}

func (s *DeleteApsDatasoureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteApsDatasoureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteApsDatasoureResponse) GetBody() *DeleteApsDatasoureResponseBody {
	return s.Body
}

func (s *DeleteApsDatasoureResponse) SetHeaders(v map[string]*string) *DeleteApsDatasoureResponse {
	s.Headers = v
	return s
}

func (s *DeleteApsDatasoureResponse) SetStatusCode(v int32) *DeleteApsDatasoureResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteApsDatasoureResponse) SetBody(v *DeleteApsDatasoureResponseBody) *DeleteApsDatasoureResponse {
	s.Body = v
	return s
}

func (s *DeleteApsDatasoureResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
