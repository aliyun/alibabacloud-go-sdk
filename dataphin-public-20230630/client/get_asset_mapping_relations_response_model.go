// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAssetMappingRelationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAssetMappingRelationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAssetMappingRelationsResponse
	GetStatusCode() *int32
	SetBody(v *GetAssetMappingRelationsResponseBody) *GetAssetMappingRelationsResponse
	GetBody() *GetAssetMappingRelationsResponseBody
}

type GetAssetMappingRelationsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAssetMappingRelationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAssetMappingRelationsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAssetMappingRelationsResponse) GoString() string {
	return s.String()
}

func (s *GetAssetMappingRelationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAssetMappingRelationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAssetMappingRelationsResponse) GetBody() *GetAssetMappingRelationsResponseBody {
	return s.Body
}

func (s *GetAssetMappingRelationsResponse) SetHeaders(v map[string]*string) *GetAssetMappingRelationsResponse {
	s.Headers = v
	return s
}

func (s *GetAssetMappingRelationsResponse) SetStatusCode(v int32) *GetAssetMappingRelationsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAssetMappingRelationsResponse) SetBody(v *GetAssetMappingRelationsResponseBody) *GetAssetMappingRelationsResponse {
	s.Body = v
	return s
}

func (s *GetAssetMappingRelationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
