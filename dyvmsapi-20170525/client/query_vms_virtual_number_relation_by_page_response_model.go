// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryVmsVirtualNumberRelationByPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryVmsVirtualNumberRelationByPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryVmsVirtualNumberRelationByPageResponse
	GetStatusCode() *int32
	SetBody(v *QueryVmsVirtualNumberRelationByPageResponseBody) *QueryVmsVirtualNumberRelationByPageResponse
	GetBody() *QueryVmsVirtualNumberRelationByPageResponseBody
}

type QueryVmsVirtualNumberRelationByPageResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryVmsVirtualNumberRelationByPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryVmsVirtualNumberRelationByPageResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryVmsVirtualNumberRelationByPageResponse) GoString() string {
	return s.String()
}

func (s *QueryVmsVirtualNumberRelationByPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryVmsVirtualNumberRelationByPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryVmsVirtualNumberRelationByPageResponse) GetBody() *QueryVmsVirtualNumberRelationByPageResponseBody {
	return s.Body
}

func (s *QueryVmsVirtualNumberRelationByPageResponse) SetHeaders(v map[string]*string) *QueryVmsVirtualNumberRelationByPageResponse {
	s.Headers = v
	return s
}

func (s *QueryVmsVirtualNumberRelationByPageResponse) SetStatusCode(v int32) *QueryVmsVirtualNumberRelationByPageResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryVmsVirtualNumberRelationByPageResponse) SetBody(v *QueryVmsVirtualNumberRelationByPageResponseBody) *QueryVmsVirtualNumberRelationByPageResponse {
	s.Body = v
	return s
}

func (s *QueryVmsVirtualNumberRelationByPageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
