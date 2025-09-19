// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAssetSelectionSelectedTargetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAssetSelectionSelectedTargetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAssetSelectionSelectedTargetResponse
	GetStatusCode() *int32
	SetBody(v *ListAssetSelectionSelectedTargetResponseBody) *ListAssetSelectionSelectedTargetResponse
	GetBody() *ListAssetSelectionSelectedTargetResponseBody
}

type ListAssetSelectionSelectedTargetResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAssetSelectionSelectedTargetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAssetSelectionSelectedTargetResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAssetSelectionSelectedTargetResponse) GoString() string {
	return s.String()
}

func (s *ListAssetSelectionSelectedTargetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAssetSelectionSelectedTargetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAssetSelectionSelectedTargetResponse) GetBody() *ListAssetSelectionSelectedTargetResponseBody {
	return s.Body
}

func (s *ListAssetSelectionSelectedTargetResponse) SetHeaders(v map[string]*string) *ListAssetSelectionSelectedTargetResponse {
	s.Headers = v
	return s
}

func (s *ListAssetSelectionSelectedTargetResponse) SetStatusCode(v int32) *ListAssetSelectionSelectedTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAssetSelectionSelectedTargetResponse) SetBody(v *ListAssetSelectionSelectedTargetResponseBody) *ListAssetSelectionSelectedTargetResponse {
	s.Body = v
	return s
}

func (s *ListAssetSelectionSelectedTargetResponse) Validate() error {
	return dara.Validate(s)
}
