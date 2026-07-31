// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadSemanticResultsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DownloadSemanticResultsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DownloadSemanticResultsResponse
	GetStatusCode() *int32
	SetBody(v *DownloadSemanticResultsResponseBody) *DownloadSemanticResultsResponse
	GetBody() *DownloadSemanticResultsResponseBody
}

type DownloadSemanticResultsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DownloadSemanticResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DownloadSemanticResultsResponse) String() string {
	return dara.Prettify(s)
}

func (s DownloadSemanticResultsResponse) GoString() string {
	return s.String()
}

func (s *DownloadSemanticResultsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DownloadSemanticResultsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DownloadSemanticResultsResponse) GetBody() *DownloadSemanticResultsResponseBody {
	return s.Body
}

func (s *DownloadSemanticResultsResponse) SetHeaders(v map[string]*string) *DownloadSemanticResultsResponse {
	s.Headers = v
	return s
}

func (s *DownloadSemanticResultsResponse) SetStatusCode(v int32) *DownloadSemanticResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *DownloadSemanticResultsResponse) SetBody(v *DownloadSemanticResultsResponseBody) *DownloadSemanticResultsResponse {
	s.Body = v
	return s
}

func (s *DownloadSemanticResultsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
