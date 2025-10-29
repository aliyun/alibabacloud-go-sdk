// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLineageRelationshipResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLineageRelationshipResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLineageRelationshipResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLineageRelationshipResponseBody) *DeleteLineageRelationshipResponse
	GetBody() *DeleteLineageRelationshipResponseBody
}

type DeleteLineageRelationshipResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLineageRelationshipResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLineageRelationshipResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLineageRelationshipResponse) GoString() string {
	return s.String()
}

func (s *DeleteLineageRelationshipResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLineageRelationshipResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLineageRelationshipResponse) GetBody() *DeleteLineageRelationshipResponseBody {
	return s.Body
}

func (s *DeleteLineageRelationshipResponse) SetHeaders(v map[string]*string) *DeleteLineageRelationshipResponse {
	s.Headers = v
	return s
}

func (s *DeleteLineageRelationshipResponse) SetStatusCode(v int32) *DeleteLineageRelationshipResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLineageRelationshipResponse) SetBody(v *DeleteLineageRelationshipResponseBody) *DeleteLineageRelationshipResponse {
	s.Body = v
	return s
}

func (s *DeleteLineageRelationshipResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
