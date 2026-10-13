// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaProfilingResultDbOverviewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMetaProfilingResultDbOverviewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMetaProfilingResultDbOverviewResponse
	GetStatusCode() *int32
	SetBody(v *GetMetaProfilingResultDbOverviewResponseBody) *GetMetaProfilingResultDbOverviewResponse
	GetBody() *GetMetaProfilingResultDbOverviewResponseBody
}

type GetMetaProfilingResultDbOverviewResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMetaProfilingResultDbOverviewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMetaProfilingResultDbOverviewResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMetaProfilingResultDbOverviewResponse) GoString() string {
	return s.String()
}

func (s *GetMetaProfilingResultDbOverviewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMetaProfilingResultDbOverviewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMetaProfilingResultDbOverviewResponse) GetBody() *GetMetaProfilingResultDbOverviewResponseBody {
	return s.Body
}

func (s *GetMetaProfilingResultDbOverviewResponse) SetHeaders(v map[string]*string) *GetMetaProfilingResultDbOverviewResponse {
	s.Headers = v
	return s
}

func (s *GetMetaProfilingResultDbOverviewResponse) SetStatusCode(v int32) *GetMetaProfilingResultDbOverviewResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMetaProfilingResultDbOverviewResponse) SetBody(v *GetMetaProfilingResultDbOverviewResponseBody) *GetMetaProfilingResultDbOverviewResponse {
	s.Body = v
	return s
}

func (s *GetMetaProfilingResultDbOverviewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
