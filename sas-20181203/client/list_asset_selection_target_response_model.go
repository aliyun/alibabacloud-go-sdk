// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAssetSelectionTargetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAssetSelectionTargetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAssetSelectionTargetResponse
	GetStatusCode() *int32
	SetBody(v *ListAssetSelectionTargetResponseBody) *ListAssetSelectionTargetResponse
	GetBody() *ListAssetSelectionTargetResponseBody
}

type ListAssetSelectionTargetResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAssetSelectionTargetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAssetSelectionTargetResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAssetSelectionTargetResponse) GoString() string {
	return s.String()
}

func (s *ListAssetSelectionTargetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAssetSelectionTargetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAssetSelectionTargetResponse) GetBody() *ListAssetSelectionTargetResponseBody {
	return s.Body
}

func (s *ListAssetSelectionTargetResponse) SetHeaders(v map[string]*string) *ListAssetSelectionTargetResponse {
	s.Headers = v
	return s
}

func (s *ListAssetSelectionTargetResponse) SetStatusCode(v int32) *ListAssetSelectionTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAssetSelectionTargetResponse) SetBody(v *ListAssetSelectionTargetResponseBody) *ListAssetSelectionTargetResponse {
	s.Body = v
	return s
}

func (s *ListAssetSelectionTargetResponse) Validate() error {
	return dara.Validate(s)
}
