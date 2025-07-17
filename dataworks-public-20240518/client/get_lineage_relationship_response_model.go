// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLineageRelationshipResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLineageRelationshipResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLineageRelationshipResponse
	GetStatusCode() *int32
	SetBody(v *GetLineageRelationshipResponseBody) *GetLineageRelationshipResponse
	GetBody() *GetLineageRelationshipResponseBody
}

type GetLineageRelationshipResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLineageRelationshipResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLineageRelationshipResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLineageRelationshipResponse) GoString() string {
	return s.String()
}

func (s *GetLineageRelationshipResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLineageRelationshipResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLineageRelationshipResponse) GetBody() *GetLineageRelationshipResponseBody {
	return s.Body
}

func (s *GetLineageRelationshipResponse) SetHeaders(v map[string]*string) *GetLineageRelationshipResponse {
	s.Headers = v
	return s
}

func (s *GetLineageRelationshipResponse) SetStatusCode(v int32) *GetLineageRelationshipResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLineageRelationshipResponse) SetBody(v *GetLineageRelationshipResponseBody) *GetLineageRelationshipResponse {
	s.Body = v
	return s
}

func (s *GetLineageRelationshipResponse) Validate() error {
	return dara.Validate(s)
}
