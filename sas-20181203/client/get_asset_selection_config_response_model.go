// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAssetSelectionConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAssetSelectionConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAssetSelectionConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetAssetSelectionConfigResponseBody) *GetAssetSelectionConfigResponse
	GetBody() *GetAssetSelectionConfigResponseBody
}

type GetAssetSelectionConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAssetSelectionConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAssetSelectionConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAssetSelectionConfigResponse) GoString() string {
	return s.String()
}

func (s *GetAssetSelectionConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAssetSelectionConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAssetSelectionConfigResponse) GetBody() *GetAssetSelectionConfigResponseBody {
	return s.Body
}

func (s *GetAssetSelectionConfigResponse) SetHeaders(v map[string]*string) *GetAssetSelectionConfigResponse {
	s.Headers = v
	return s
}

func (s *GetAssetSelectionConfigResponse) SetStatusCode(v int32) *GetAssetSelectionConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAssetSelectionConfigResponse) SetBody(v *GetAssetSelectionConfigResponseBody) *GetAssetSelectionConfigResponse {
	s.Body = v
	return s
}

func (s *GetAssetSelectionConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
