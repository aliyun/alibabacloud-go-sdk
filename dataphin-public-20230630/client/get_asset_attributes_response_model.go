// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAssetAttributesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAssetAttributesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAssetAttributesResponse
	GetStatusCode() *int32
	SetBody(v *GetAssetAttributesResponseBody) *GetAssetAttributesResponse
	GetBody() *GetAssetAttributesResponseBody
}

type GetAssetAttributesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAssetAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAssetAttributesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAssetAttributesResponse) GoString() string {
	return s.String()
}

func (s *GetAssetAttributesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAssetAttributesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAssetAttributesResponse) GetBody() *GetAssetAttributesResponseBody {
	return s.Body
}

func (s *GetAssetAttributesResponse) SetHeaders(v map[string]*string) *GetAssetAttributesResponse {
	s.Headers = v
	return s
}

func (s *GetAssetAttributesResponse) SetStatusCode(v int32) *GetAssetAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAssetAttributesResponse) SetBody(v *GetAssetAttributesResponseBody) *GetAssetAttributesResponse {
	s.Body = v
	return s
}

func (s *GetAssetAttributesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
