// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAutoResourceOptimizeRulesAsyncResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAutoResourceOptimizeRulesAsyncResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAutoResourceOptimizeRulesAsyncResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAutoResourceOptimizeRulesAsyncResponseBody) *UpdateAutoResourceOptimizeRulesAsyncResponse
	GetBody() *UpdateAutoResourceOptimizeRulesAsyncResponseBody
}

type UpdateAutoResourceOptimizeRulesAsyncResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAutoResourceOptimizeRulesAsyncResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAutoResourceOptimizeRulesAsyncResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAutoResourceOptimizeRulesAsyncResponse) GoString() string {
	return s.String()
}

func (s *UpdateAutoResourceOptimizeRulesAsyncResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAutoResourceOptimizeRulesAsyncResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAutoResourceOptimizeRulesAsyncResponse) GetBody() *UpdateAutoResourceOptimizeRulesAsyncResponseBody {
	return s.Body
}

func (s *UpdateAutoResourceOptimizeRulesAsyncResponse) SetHeaders(v map[string]*string) *UpdateAutoResourceOptimizeRulesAsyncResponse {
	s.Headers = v
	return s
}

func (s *UpdateAutoResourceOptimizeRulesAsyncResponse) SetStatusCode(v int32) *UpdateAutoResourceOptimizeRulesAsyncResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAutoResourceOptimizeRulesAsyncResponse) SetBody(v *UpdateAutoResourceOptimizeRulesAsyncResponseBody) *UpdateAutoResourceOptimizeRulesAsyncResponse {
	s.Body = v
	return s
}

func (s *UpdateAutoResourceOptimizeRulesAsyncResponse) Validate() error {
	return dara.Validate(s)
}
