// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAssetSelectionConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAssetSelectionConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAssetSelectionConfigResponse
	GetStatusCode() *int32
	SetBody(v *CreateAssetSelectionConfigResponseBody) *CreateAssetSelectionConfigResponse
	GetBody() *CreateAssetSelectionConfigResponseBody
}

type CreateAssetSelectionConfigResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAssetSelectionConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAssetSelectionConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAssetSelectionConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateAssetSelectionConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAssetSelectionConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAssetSelectionConfigResponse) GetBody() *CreateAssetSelectionConfigResponseBody {
	return s.Body
}

func (s *CreateAssetSelectionConfigResponse) SetHeaders(v map[string]*string) *CreateAssetSelectionConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateAssetSelectionConfigResponse) SetStatusCode(v int32) *CreateAssetSelectionConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAssetSelectionConfigResponse) SetBody(v *CreateAssetSelectionConfigResponseBody) *CreateAssetSelectionConfigResponse {
	s.Body = v
	return s
}

func (s *CreateAssetSelectionConfigResponse) Validate() error {
	return dara.Validate(s)
}
