// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStepResultOverviewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetStepResultOverviewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetStepResultOverviewResponse
	GetStatusCode() *int32
	SetBody(v *GetStepResultOverviewResponseBody) *GetStepResultOverviewResponse
	GetBody() *GetStepResultOverviewResponseBody
}

type GetStepResultOverviewResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetStepResultOverviewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetStepResultOverviewResponse) String() string {
	return dara.Prettify(s)
}

func (s GetStepResultOverviewResponse) GoString() string {
	return s.String()
}

func (s *GetStepResultOverviewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetStepResultOverviewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetStepResultOverviewResponse) GetBody() *GetStepResultOverviewResponseBody {
	return s.Body
}

func (s *GetStepResultOverviewResponse) SetHeaders(v map[string]*string) *GetStepResultOverviewResponse {
	s.Headers = v
	return s
}

func (s *GetStepResultOverviewResponse) SetStatusCode(v int32) *GetStepResultOverviewResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStepResultOverviewResponse) SetBody(v *GetStepResultOverviewResponseBody) *GetStepResultOverviewResponse {
	s.Body = v
	return s
}

func (s *GetStepResultOverviewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
