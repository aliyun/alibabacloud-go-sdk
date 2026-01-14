// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDomainAcceleratorRelationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDomainAcceleratorRelationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDomainAcceleratorRelationResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDomainAcceleratorRelationResponseBody) *DeleteDomainAcceleratorRelationResponse
	GetBody() *DeleteDomainAcceleratorRelationResponseBody
}

type DeleteDomainAcceleratorRelationResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDomainAcceleratorRelationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDomainAcceleratorRelationResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDomainAcceleratorRelationResponse) GoString() string {
	return s.String()
}

func (s *DeleteDomainAcceleratorRelationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDomainAcceleratorRelationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDomainAcceleratorRelationResponse) GetBody() *DeleteDomainAcceleratorRelationResponseBody {
	return s.Body
}

func (s *DeleteDomainAcceleratorRelationResponse) SetHeaders(v map[string]*string) *DeleteDomainAcceleratorRelationResponse {
	s.Headers = v
	return s
}

func (s *DeleteDomainAcceleratorRelationResponse) SetStatusCode(v int32) *DeleteDomainAcceleratorRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDomainAcceleratorRelationResponse) SetBody(v *DeleteDomainAcceleratorRelationResponseBody) *DeleteDomainAcceleratorRelationResponse {
	s.Body = v
	return s
}

func (s *DeleteDomainAcceleratorRelationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
