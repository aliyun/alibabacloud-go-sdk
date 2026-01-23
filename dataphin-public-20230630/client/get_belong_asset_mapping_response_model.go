// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBelongAssetMappingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBelongAssetMappingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBelongAssetMappingResponse
	GetStatusCode() *int32
	SetBody(v *GetBelongAssetMappingResponseBody) *GetBelongAssetMappingResponse
	GetBody() *GetBelongAssetMappingResponseBody
}

type GetBelongAssetMappingResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBelongAssetMappingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBelongAssetMappingResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBelongAssetMappingResponse) GoString() string {
	return s.String()
}

func (s *GetBelongAssetMappingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBelongAssetMappingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBelongAssetMappingResponse) GetBody() *GetBelongAssetMappingResponseBody {
	return s.Body
}

func (s *GetBelongAssetMappingResponse) SetHeaders(v map[string]*string) *GetBelongAssetMappingResponse {
	s.Headers = v
	return s
}

func (s *GetBelongAssetMappingResponse) SetStatusCode(v int32) *GetBelongAssetMappingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBelongAssetMappingResponse) SetBody(v *GetBelongAssetMappingResponseBody) *GetBelongAssetMappingResponse {
	s.Body = v
	return s
}

func (s *GetBelongAssetMappingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
