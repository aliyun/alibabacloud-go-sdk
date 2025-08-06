// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelBucketRedundancyTransitionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelBucketRedundancyTransitionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelBucketRedundancyTransitionResponse
	GetStatusCode() *int32
	SetBody(v *CancelBucketRedundancyTransitionResponseBody) *CancelBucketRedundancyTransitionResponse
	GetBody() *CancelBucketRedundancyTransitionResponseBody
}

type CancelBucketRedundancyTransitionResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelBucketRedundancyTransitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelBucketRedundancyTransitionResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelBucketRedundancyTransitionResponse) GoString() string {
	return s.String()
}

func (s *CancelBucketRedundancyTransitionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelBucketRedundancyTransitionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelBucketRedundancyTransitionResponse) GetBody() *CancelBucketRedundancyTransitionResponseBody {
	return s.Body
}

func (s *CancelBucketRedundancyTransitionResponse) SetHeaders(v map[string]*string) *CancelBucketRedundancyTransitionResponse {
	s.Headers = v
	return s
}

func (s *CancelBucketRedundancyTransitionResponse) SetStatusCode(v int32) *CancelBucketRedundancyTransitionResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelBucketRedundancyTransitionResponse) SetBody(v *CancelBucketRedundancyTransitionResponseBody) *CancelBucketRedundancyTransitionResponse {
	s.Body = v
	return s
}

func (s *CancelBucketRedundancyTransitionResponse) Validate() error {
	return dara.Validate(s)
}
