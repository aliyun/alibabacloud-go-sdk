// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAssetTypeAttributeCodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAssetTypeAttributeCodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAssetTypeAttributeCodesResponse
	GetStatusCode() *int32
	SetBody(v *GetAssetTypeAttributeCodesResponseBody) *GetAssetTypeAttributeCodesResponse
	GetBody() *GetAssetTypeAttributeCodesResponseBody
}

type GetAssetTypeAttributeCodesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAssetTypeAttributeCodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAssetTypeAttributeCodesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAssetTypeAttributeCodesResponse) GoString() string {
	return s.String()
}

func (s *GetAssetTypeAttributeCodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAssetTypeAttributeCodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAssetTypeAttributeCodesResponse) GetBody() *GetAssetTypeAttributeCodesResponseBody {
	return s.Body
}

func (s *GetAssetTypeAttributeCodesResponse) SetHeaders(v map[string]*string) *GetAssetTypeAttributeCodesResponse {
	s.Headers = v
	return s
}

func (s *GetAssetTypeAttributeCodesResponse) SetStatusCode(v int32) *GetAssetTypeAttributeCodesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAssetTypeAttributeCodesResponse) SetBody(v *GetAssetTypeAttributeCodesResponseBody) *GetAssetTypeAttributeCodesResponse {
	s.Body = v
	return s
}

func (s *GetAssetTypeAttributeCodesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
