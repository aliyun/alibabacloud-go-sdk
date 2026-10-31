// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApprovalTaskDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetApprovalTaskDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetApprovalTaskDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetApprovalTaskDetailResponseBody) *GetApprovalTaskDetailResponse
	GetBody() *GetApprovalTaskDetailResponseBody
}

type GetApprovalTaskDetailResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApprovalTaskDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApprovalTaskDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetApprovalTaskDetailResponse) GoString() string {
	return s.String()
}

func (s *GetApprovalTaskDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetApprovalTaskDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetApprovalTaskDetailResponse) GetBody() *GetApprovalTaskDetailResponseBody {
	return s.Body
}

func (s *GetApprovalTaskDetailResponse) SetHeaders(v map[string]*string) *GetApprovalTaskDetailResponse {
	s.Headers = v
	return s
}

func (s *GetApprovalTaskDetailResponse) SetStatusCode(v int32) *GetApprovalTaskDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApprovalTaskDetailResponse) SetBody(v *GetApprovalTaskDetailResponseBody) *GetApprovalTaskDetailResponse {
	s.Body = v
	return s
}

func (s *GetApprovalTaskDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
