// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAssetAttributesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAssetAttributesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAssetAttributesResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAssetAttributesResponseBody) *UpdateAssetAttributesResponse
	GetBody() *UpdateAssetAttributesResponseBody
}

type UpdateAssetAttributesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAssetAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAssetAttributesResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAssetAttributesResponse) GoString() string {
	return s.String()
}

func (s *UpdateAssetAttributesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAssetAttributesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAssetAttributesResponse) GetBody() *UpdateAssetAttributesResponseBody {
	return s.Body
}

func (s *UpdateAssetAttributesResponse) SetHeaders(v map[string]*string) *UpdateAssetAttributesResponse {
	s.Headers = v
	return s
}

func (s *UpdateAssetAttributesResponse) SetStatusCode(v int32) *UpdateAssetAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAssetAttributesResponse) SetBody(v *UpdateAssetAttributesResponseBody) *UpdateAssetAttributesResponse {
	s.Body = v
	return s
}

func (s *UpdateAssetAttributesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
