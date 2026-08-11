// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiAppOverviewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAiAppOverviewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAiAppOverviewResponse
	GetStatusCode() *int32
	SetBody(v *GetAiAppOverviewResponseBody) *GetAiAppOverviewResponse
	GetBody() *GetAiAppOverviewResponseBody
}

type GetAiAppOverviewResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAiAppOverviewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAiAppOverviewResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAiAppOverviewResponse) GoString() string {
	return s.String()
}

func (s *GetAiAppOverviewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAiAppOverviewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAiAppOverviewResponse) GetBody() *GetAiAppOverviewResponseBody {
	return s.Body
}

func (s *GetAiAppOverviewResponse) SetHeaders(v map[string]*string) *GetAiAppOverviewResponse {
	s.Headers = v
	return s
}

func (s *GetAiAppOverviewResponse) SetStatusCode(v int32) *GetAiAppOverviewResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAiAppOverviewResponse) SetBody(v *GetAiAppOverviewResponseBody) *GetAiAppOverviewResponse {
	s.Body = v
	return s
}

func (s *GetAiAppOverviewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
