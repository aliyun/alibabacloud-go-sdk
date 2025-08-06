// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitBucketRedundancyTransitionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitBucketRedundancyTransitionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitBucketRedundancyTransitionResponse
	GetStatusCode() *int32
	SetBody(v *SubmitBucketRedundancyTransitionResponseBody) *SubmitBucketRedundancyTransitionResponse
	GetBody() *SubmitBucketRedundancyTransitionResponseBody
}

type SubmitBucketRedundancyTransitionResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitBucketRedundancyTransitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitBucketRedundancyTransitionResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitBucketRedundancyTransitionResponse) GoString() string {
	return s.String()
}

func (s *SubmitBucketRedundancyTransitionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitBucketRedundancyTransitionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitBucketRedundancyTransitionResponse) GetBody() *SubmitBucketRedundancyTransitionResponseBody {
	return s.Body
}

func (s *SubmitBucketRedundancyTransitionResponse) SetHeaders(v map[string]*string) *SubmitBucketRedundancyTransitionResponse {
	s.Headers = v
	return s
}

func (s *SubmitBucketRedundancyTransitionResponse) SetStatusCode(v int32) *SubmitBucketRedundancyTransitionResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitBucketRedundancyTransitionResponse) SetBody(v *SubmitBucketRedundancyTransitionResponseBody) *SubmitBucketRedundancyTransitionResponse {
	s.Body = v
	return s
}

func (s *SubmitBucketRedundancyTransitionResponse) Validate() error {
	return dara.Validate(s)
}
