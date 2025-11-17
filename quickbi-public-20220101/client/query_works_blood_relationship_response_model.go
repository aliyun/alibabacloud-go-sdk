// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryWorksBloodRelationshipResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryWorksBloodRelationshipResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryWorksBloodRelationshipResponse
	GetStatusCode() *int32
	SetBody(v *QueryWorksBloodRelationshipResponseBody) *QueryWorksBloodRelationshipResponse
	GetBody() *QueryWorksBloodRelationshipResponseBody
}

type QueryWorksBloodRelationshipResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryWorksBloodRelationshipResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryWorksBloodRelationshipResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryWorksBloodRelationshipResponse) GoString() string {
	return s.String()
}

func (s *QueryWorksBloodRelationshipResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryWorksBloodRelationshipResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryWorksBloodRelationshipResponse) GetBody() *QueryWorksBloodRelationshipResponseBody {
	return s.Body
}

func (s *QueryWorksBloodRelationshipResponse) SetHeaders(v map[string]*string) *QueryWorksBloodRelationshipResponse {
	s.Headers = v
	return s
}

func (s *QueryWorksBloodRelationshipResponse) SetStatusCode(v int32) *QueryWorksBloodRelationshipResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryWorksBloodRelationshipResponse) SetBody(v *QueryWorksBloodRelationshipResponseBody) *QueryWorksBloodRelationshipResponse {
	s.Body = v
	return s
}

func (s *QueryWorksBloodRelationshipResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
