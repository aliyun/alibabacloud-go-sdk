// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAssetSelectionCriteriaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddAssetSelectionCriteriaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddAssetSelectionCriteriaResponse
	GetStatusCode() *int32
	SetBody(v *AddAssetSelectionCriteriaResponseBody) *AddAssetSelectionCriteriaResponse
	GetBody() *AddAssetSelectionCriteriaResponseBody
}

type AddAssetSelectionCriteriaResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddAssetSelectionCriteriaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddAssetSelectionCriteriaResponse) String() string {
	return dara.Prettify(s)
}

func (s AddAssetSelectionCriteriaResponse) GoString() string {
	return s.String()
}

func (s *AddAssetSelectionCriteriaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddAssetSelectionCriteriaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddAssetSelectionCriteriaResponse) GetBody() *AddAssetSelectionCriteriaResponseBody {
	return s.Body
}

func (s *AddAssetSelectionCriteriaResponse) SetHeaders(v map[string]*string) *AddAssetSelectionCriteriaResponse {
	s.Headers = v
	return s
}

func (s *AddAssetSelectionCriteriaResponse) SetStatusCode(v int32) *AddAssetSelectionCriteriaResponse {
	s.StatusCode = &v
	return s
}

func (s *AddAssetSelectionCriteriaResponse) SetBody(v *AddAssetSelectionCriteriaResponseBody) *AddAssetSelectionCriteriaResponse {
	s.Body = v
	return s
}

func (s *AddAssetSelectionCriteriaResponse) Validate() error {
	return dara.Validate(s)
}
