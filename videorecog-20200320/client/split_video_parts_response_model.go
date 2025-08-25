// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSplitVideoPartsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SplitVideoPartsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SplitVideoPartsResponse
	GetStatusCode() *int32
	SetBody(v *SplitVideoPartsResponseBody) *SplitVideoPartsResponse
	GetBody() *SplitVideoPartsResponseBody
}

type SplitVideoPartsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SplitVideoPartsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SplitVideoPartsResponse) String() string {
	return dara.Prettify(s)
}

func (s SplitVideoPartsResponse) GoString() string {
	return s.String()
}

func (s *SplitVideoPartsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SplitVideoPartsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SplitVideoPartsResponse) GetBody() *SplitVideoPartsResponseBody {
	return s.Body
}

func (s *SplitVideoPartsResponse) SetHeaders(v map[string]*string) *SplitVideoPartsResponse {
	s.Headers = v
	return s
}

func (s *SplitVideoPartsResponse) SetStatusCode(v int32) *SplitVideoPartsResponse {
	s.StatusCode = &v
	return s
}

func (s *SplitVideoPartsResponse) SetBody(v *SplitVideoPartsResponseBody) *SplitVideoPartsResponse {
	s.Body = v
	return s
}

func (s *SplitVideoPartsResponse) Validate() error {
	return dara.Validate(s)
}
