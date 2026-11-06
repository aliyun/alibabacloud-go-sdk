// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkItemRunSummariesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWorkItemRunSummariesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWorkItemRunSummariesResponse
	GetStatusCode() *int32
	SetBody(v *GetWorkItemRunSummariesResponseBody) *GetWorkItemRunSummariesResponse
	GetBody() *GetWorkItemRunSummariesResponseBody
}

type GetWorkItemRunSummariesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWorkItemRunSummariesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWorkItemRunSummariesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWorkItemRunSummariesResponse) GoString() string {
	return s.String()
}

func (s *GetWorkItemRunSummariesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWorkItemRunSummariesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWorkItemRunSummariesResponse) GetBody() *GetWorkItemRunSummariesResponseBody {
	return s.Body
}

func (s *GetWorkItemRunSummariesResponse) SetHeaders(v map[string]*string) *GetWorkItemRunSummariesResponse {
	s.Headers = v
	return s
}

func (s *GetWorkItemRunSummariesResponse) SetStatusCode(v int32) *GetWorkItemRunSummariesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkItemRunSummariesResponse) SetBody(v *GetWorkItemRunSummariesResponseBody) *GetWorkItemRunSummariesResponse {
	s.Body = v
	return s
}

func (s *GetWorkItemRunSummariesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
