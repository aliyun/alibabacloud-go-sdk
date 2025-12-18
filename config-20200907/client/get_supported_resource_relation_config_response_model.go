// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSupportedResourceRelationConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSupportedResourceRelationConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSupportedResourceRelationConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetSupportedResourceRelationConfigResponseBody) *GetSupportedResourceRelationConfigResponse
	GetBody() *GetSupportedResourceRelationConfigResponseBody
}

type GetSupportedResourceRelationConfigResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSupportedResourceRelationConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSupportedResourceRelationConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSupportedResourceRelationConfigResponse) GoString() string {
	return s.String()
}

func (s *GetSupportedResourceRelationConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSupportedResourceRelationConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSupportedResourceRelationConfigResponse) GetBody() *GetSupportedResourceRelationConfigResponseBody {
	return s.Body
}

func (s *GetSupportedResourceRelationConfigResponse) SetHeaders(v map[string]*string) *GetSupportedResourceRelationConfigResponse {
	s.Headers = v
	return s
}

func (s *GetSupportedResourceRelationConfigResponse) SetStatusCode(v int32) *GetSupportedResourceRelationConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSupportedResourceRelationConfigResponse) SetBody(v *GetSupportedResourceRelationConfigResponseBody) *GetSupportedResourceRelationConfigResponse {
	s.Body = v
	return s
}

func (s *GetSupportedResourceRelationConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
