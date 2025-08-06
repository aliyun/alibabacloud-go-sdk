// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBucketRedundancyTransitionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBucketRedundancyTransitionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBucketRedundancyTransitionResponse
	GetStatusCode() *int32
	SetBody(v *ListBucketRedundancyTransitionResponseBody) *ListBucketRedundancyTransitionResponse
	GetBody() *ListBucketRedundancyTransitionResponseBody
}

type ListBucketRedundancyTransitionResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBucketRedundancyTransitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBucketRedundancyTransitionResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBucketRedundancyTransitionResponse) GoString() string {
	return s.String()
}

func (s *ListBucketRedundancyTransitionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBucketRedundancyTransitionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBucketRedundancyTransitionResponse) GetBody() *ListBucketRedundancyTransitionResponseBody {
	return s.Body
}

func (s *ListBucketRedundancyTransitionResponse) SetHeaders(v map[string]*string) *ListBucketRedundancyTransitionResponse {
	s.Headers = v
	return s
}

func (s *ListBucketRedundancyTransitionResponse) SetStatusCode(v int32) *ListBucketRedundancyTransitionResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBucketRedundancyTransitionResponse) SetBody(v *ListBucketRedundancyTransitionResponseBody) *ListBucketRedundancyTransitionResponse {
	s.Body = v
	return s
}

func (s *ListBucketRedundancyTransitionResponse) Validate() error {
	return dara.Validate(s)
}
