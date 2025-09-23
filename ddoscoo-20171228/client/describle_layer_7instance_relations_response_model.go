// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribleLayer7InstanceRelationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribleLayer7InstanceRelationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribleLayer7InstanceRelationsResponse
	GetStatusCode() *int32
	SetBody(v *DescribleLayer7InstanceRelationsResponseBody) *DescribleLayer7InstanceRelationsResponse
	GetBody() *DescribleLayer7InstanceRelationsResponseBody
}

type DescribleLayer7InstanceRelationsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribleLayer7InstanceRelationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribleLayer7InstanceRelationsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribleLayer7InstanceRelationsResponse) GoString() string {
	return s.String()
}

func (s *DescribleLayer7InstanceRelationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribleLayer7InstanceRelationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribleLayer7InstanceRelationsResponse) GetBody() *DescribleLayer7InstanceRelationsResponseBody {
	return s.Body
}

func (s *DescribleLayer7InstanceRelationsResponse) SetHeaders(v map[string]*string) *DescribleLayer7InstanceRelationsResponse {
	s.Headers = v
	return s
}

func (s *DescribleLayer7InstanceRelationsResponse) SetStatusCode(v int32) *DescribleLayer7InstanceRelationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribleLayer7InstanceRelationsResponse) SetBody(v *DescribleLayer7InstanceRelationsResponseBody) *DescribleLayer7InstanceRelationsResponse {
	s.Body = v
	return s
}

func (s *DescribleLayer7InstanceRelationsResponse) Validate() error {
	return dara.Validate(s)
}
