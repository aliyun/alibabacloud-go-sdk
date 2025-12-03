// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPipelineJobHistorysResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPipelineJobHistorysResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPipelineJobHistorysResponse
	GetStatusCode() *int32
	SetBody(v *ListPipelineJobHistorysResponseBody) *ListPipelineJobHistorysResponse
	GetBody() *ListPipelineJobHistorysResponseBody
}

type ListPipelineJobHistorysResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPipelineJobHistorysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPipelineJobHistorysResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPipelineJobHistorysResponse) GoString() string {
	return s.String()
}

func (s *ListPipelineJobHistorysResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPipelineJobHistorysResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPipelineJobHistorysResponse) GetBody() *ListPipelineJobHistorysResponseBody {
	return s.Body
}

func (s *ListPipelineJobHistorysResponse) SetHeaders(v map[string]*string) *ListPipelineJobHistorysResponse {
	s.Headers = v
	return s
}

func (s *ListPipelineJobHistorysResponse) SetStatusCode(v int32) *ListPipelineJobHistorysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPipelineJobHistorysResponse) SetBody(v *ListPipelineJobHistorysResponseBody) *ListPipelineJobHistorysResponse {
	s.Body = v
	return s
}

func (s *ListPipelineJobHistorysResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
