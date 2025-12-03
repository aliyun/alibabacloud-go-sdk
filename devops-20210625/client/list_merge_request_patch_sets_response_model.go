// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMergeRequestPatchSetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMergeRequestPatchSetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMergeRequestPatchSetsResponse
	GetStatusCode() *int32
	SetBody(v *ListMergeRequestPatchSetsResponseBody) *ListMergeRequestPatchSetsResponse
	GetBody() *ListMergeRequestPatchSetsResponseBody
}

type ListMergeRequestPatchSetsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMergeRequestPatchSetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMergeRequestPatchSetsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMergeRequestPatchSetsResponse) GoString() string {
	return s.String()
}

func (s *ListMergeRequestPatchSetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMergeRequestPatchSetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMergeRequestPatchSetsResponse) GetBody() *ListMergeRequestPatchSetsResponseBody {
	return s.Body
}

func (s *ListMergeRequestPatchSetsResponse) SetHeaders(v map[string]*string) *ListMergeRequestPatchSetsResponse {
	s.Headers = v
	return s
}

func (s *ListMergeRequestPatchSetsResponse) SetStatusCode(v int32) *ListMergeRequestPatchSetsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMergeRequestPatchSetsResponse) SetBody(v *ListMergeRequestPatchSetsResponseBody) *ListMergeRequestPatchSetsResponse {
	s.Body = v
	return s
}

func (s *ListMergeRequestPatchSetsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
