// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLineageRelationshipResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLineageRelationshipResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLineageRelationshipResponse
	GetStatusCode() *int32
	SetBody(v *CreateLineageRelationshipResponseBody) *CreateLineageRelationshipResponse
	GetBody() *CreateLineageRelationshipResponseBody
}

type CreateLineageRelationshipResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLineageRelationshipResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLineageRelationshipResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLineageRelationshipResponse) GoString() string {
	return s.String()
}

func (s *CreateLineageRelationshipResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLineageRelationshipResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLineageRelationshipResponse) GetBody() *CreateLineageRelationshipResponseBody {
	return s.Body
}

func (s *CreateLineageRelationshipResponse) SetHeaders(v map[string]*string) *CreateLineageRelationshipResponse {
	s.Headers = v
	return s
}

func (s *CreateLineageRelationshipResponse) SetStatusCode(v int32) *CreateLineageRelationshipResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLineageRelationshipResponse) SetBody(v *CreateLineageRelationshipResponseBody) *CreateLineageRelationshipResponse {
	s.Body = v
	return s
}

func (s *CreateLineageRelationshipResponse) Validate() error {
	return dara.Validate(s)
}
