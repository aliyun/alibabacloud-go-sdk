// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDiskPerformanceLevelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDiskPerformanceLevelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDiskPerformanceLevelResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDiskPerformanceLevelResponseBody) *ModifyDiskPerformanceLevelResponse
	GetBody() *ModifyDiskPerformanceLevelResponseBody
}

type ModifyDiskPerformanceLevelResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDiskPerformanceLevelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDiskPerformanceLevelResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDiskPerformanceLevelResponse) GoString() string {
	return s.String()
}

func (s *ModifyDiskPerformanceLevelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDiskPerformanceLevelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDiskPerformanceLevelResponse) GetBody() *ModifyDiskPerformanceLevelResponseBody {
	return s.Body
}

func (s *ModifyDiskPerformanceLevelResponse) SetHeaders(v map[string]*string) *ModifyDiskPerformanceLevelResponse {
	s.Headers = v
	return s
}

func (s *ModifyDiskPerformanceLevelResponse) SetStatusCode(v int32) *ModifyDiskPerformanceLevelResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDiskPerformanceLevelResponse) SetBody(v *ModifyDiskPerformanceLevelResponseBody) *ModifyDiskPerformanceLevelResponse {
	s.Body = v
	return s
}

func (s *ModifyDiskPerformanceLevelResponse) Validate() error {
	return dara.Validate(s)
}
