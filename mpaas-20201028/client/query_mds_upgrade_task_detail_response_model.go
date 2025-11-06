// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMdsUpgradeTaskDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMdsUpgradeTaskDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMdsUpgradeTaskDetailResponse
	GetStatusCode() *int32
	SetBody(v *QueryMdsUpgradeTaskDetailResponseBody) *QueryMdsUpgradeTaskDetailResponse
	GetBody() *QueryMdsUpgradeTaskDetailResponseBody
}

type QueryMdsUpgradeTaskDetailResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMdsUpgradeTaskDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMdsUpgradeTaskDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMdsUpgradeTaskDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryMdsUpgradeTaskDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMdsUpgradeTaskDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMdsUpgradeTaskDetailResponse) GetBody() *QueryMdsUpgradeTaskDetailResponseBody {
	return s.Body
}

func (s *QueryMdsUpgradeTaskDetailResponse) SetHeaders(v map[string]*string) *QueryMdsUpgradeTaskDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponse) SetStatusCode(v int32) *QueryMdsUpgradeTaskDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponse) SetBody(v *QueryMdsUpgradeTaskDetailResponseBody) *QueryMdsUpgradeTaskDetailResponse {
	s.Body = v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
