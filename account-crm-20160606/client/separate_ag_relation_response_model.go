// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSeparateAgRelationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SeparateAgRelationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SeparateAgRelationResponse
	GetStatusCode() *int32
	SetBody(v *SeparateAgRelationResponseBody) *SeparateAgRelationResponse
	GetBody() *SeparateAgRelationResponseBody
}

type SeparateAgRelationResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SeparateAgRelationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SeparateAgRelationResponse) String() string {
	return dara.Prettify(s)
}

func (s SeparateAgRelationResponse) GoString() string {
	return s.String()
}

func (s *SeparateAgRelationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SeparateAgRelationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SeparateAgRelationResponse) GetBody() *SeparateAgRelationResponseBody {
	return s.Body
}

func (s *SeparateAgRelationResponse) SetHeaders(v map[string]*string) *SeparateAgRelationResponse {
	s.Headers = v
	return s
}

func (s *SeparateAgRelationResponse) SetStatusCode(v int32) *SeparateAgRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *SeparateAgRelationResponse) SetBody(v *SeparateAgRelationResponseBody) *SeparateAgRelationResponse {
	s.Body = v
	return s
}

func (s *SeparateAgRelationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
