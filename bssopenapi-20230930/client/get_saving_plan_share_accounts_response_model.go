// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSavingPlanShareAccountsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSavingPlanShareAccountsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSavingPlanShareAccountsResponse
	GetStatusCode() *int32
	SetBody(v *GetSavingPlanShareAccountsResponseBody) *GetSavingPlanShareAccountsResponse
	GetBody() *GetSavingPlanShareAccountsResponseBody
}

type GetSavingPlanShareAccountsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSavingPlanShareAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSavingPlanShareAccountsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSavingPlanShareAccountsResponse) GoString() string {
	return s.String()
}

func (s *GetSavingPlanShareAccountsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSavingPlanShareAccountsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSavingPlanShareAccountsResponse) GetBody() *GetSavingPlanShareAccountsResponseBody {
	return s.Body
}

func (s *GetSavingPlanShareAccountsResponse) SetHeaders(v map[string]*string) *GetSavingPlanShareAccountsResponse {
	s.Headers = v
	return s
}

func (s *GetSavingPlanShareAccountsResponse) SetStatusCode(v int32) *GetSavingPlanShareAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSavingPlanShareAccountsResponse) SetBody(v *GetSavingPlanShareAccountsResponseBody) *GetSavingPlanShareAccountsResponse {
	s.Body = v
	return s
}

func (s *GetSavingPlanShareAccountsResponse) Validate() error {
	return dara.Validate(s)
}
