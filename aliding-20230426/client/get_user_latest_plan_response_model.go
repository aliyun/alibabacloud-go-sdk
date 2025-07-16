// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserLatestPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserLatestPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserLatestPlanResponse
	GetStatusCode() *int32
	SetBody(v *GetUserLatestPlanResponseBody) *GetUserLatestPlanResponse
	GetBody() *GetUserLatestPlanResponseBody
}

type GetUserLatestPlanResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserLatestPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserLatestPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserLatestPlanResponse) GoString() string {
	return s.String()
}

func (s *GetUserLatestPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserLatestPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserLatestPlanResponse) GetBody() *GetUserLatestPlanResponseBody {
	return s.Body
}

func (s *GetUserLatestPlanResponse) SetHeaders(v map[string]*string) *GetUserLatestPlanResponse {
	s.Headers = v
	return s
}

func (s *GetUserLatestPlanResponse) SetStatusCode(v int32) *GetUserLatestPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserLatestPlanResponse) SetBody(v *GetUserLatestPlanResponseBody) *GetUserLatestPlanResponse {
	s.Body = v
	return s
}

func (s *GetUserLatestPlanResponse) Validate() error {
	return dara.Validate(s)
}
