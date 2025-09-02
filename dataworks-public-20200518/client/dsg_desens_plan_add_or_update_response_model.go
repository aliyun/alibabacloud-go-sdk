// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgDesensPlanAddOrUpdateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DsgDesensPlanAddOrUpdateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DsgDesensPlanAddOrUpdateResponse
	GetStatusCode() *int32
	SetBody(v *DsgDesensPlanAddOrUpdateResponseBody) *DsgDesensPlanAddOrUpdateResponse
	GetBody() *DsgDesensPlanAddOrUpdateResponseBody
}

type DsgDesensPlanAddOrUpdateResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DsgDesensPlanAddOrUpdateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DsgDesensPlanAddOrUpdateResponse) String() string {
	return dara.Prettify(s)
}

func (s DsgDesensPlanAddOrUpdateResponse) GoString() string {
	return s.String()
}

func (s *DsgDesensPlanAddOrUpdateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DsgDesensPlanAddOrUpdateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DsgDesensPlanAddOrUpdateResponse) GetBody() *DsgDesensPlanAddOrUpdateResponseBody {
	return s.Body
}

func (s *DsgDesensPlanAddOrUpdateResponse) SetHeaders(v map[string]*string) *DsgDesensPlanAddOrUpdateResponse {
	s.Headers = v
	return s
}

func (s *DsgDesensPlanAddOrUpdateResponse) SetStatusCode(v int32) *DsgDesensPlanAddOrUpdateResponse {
	s.StatusCode = &v
	return s
}

func (s *DsgDesensPlanAddOrUpdateResponse) SetBody(v *DsgDesensPlanAddOrUpdateResponseBody) *DsgDesensPlanAddOrUpdateResponse {
	s.Body = v
	return s
}

func (s *DsgDesensPlanAddOrUpdateResponse) Validate() error {
	return dara.Validate(s)
}
