// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryVirtualNumberRelationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryVirtualNumberRelationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryVirtualNumberRelationResponse
	GetStatusCode() *int32
	SetBody(v *QueryVirtualNumberRelationResponseBody) *QueryVirtualNumberRelationResponse
	GetBody() *QueryVirtualNumberRelationResponseBody
}

type QueryVirtualNumberRelationResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryVirtualNumberRelationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryVirtualNumberRelationResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryVirtualNumberRelationResponse) GoString() string {
	return s.String()
}

func (s *QueryVirtualNumberRelationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryVirtualNumberRelationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryVirtualNumberRelationResponse) GetBody() *QueryVirtualNumberRelationResponseBody {
	return s.Body
}

func (s *QueryVirtualNumberRelationResponse) SetHeaders(v map[string]*string) *QueryVirtualNumberRelationResponse {
	s.Headers = v
	return s
}

func (s *QueryVirtualNumberRelationResponse) SetStatusCode(v int32) *QueryVirtualNumberRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryVirtualNumberRelationResponse) SetBody(v *QueryVirtualNumberRelationResponseBody) *QueryVirtualNumberRelationResponse {
	s.Body = v
	return s
}

func (s *QueryVirtualNumberRelationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
