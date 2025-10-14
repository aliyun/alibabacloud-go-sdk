// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAncillarySuggestResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AncillarySuggestResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AncillarySuggestResponse
	GetStatusCode() *int32
	SetBody(v *AncillarySuggestResponseBody) *AncillarySuggestResponse
	GetBody() *AncillarySuggestResponseBody
}

type AncillarySuggestResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AncillarySuggestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AncillarySuggestResponse) String() string {
	return dara.Prettify(s)
}

func (s AncillarySuggestResponse) GoString() string {
	return s.String()
}

func (s *AncillarySuggestResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AncillarySuggestResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AncillarySuggestResponse) GetBody() *AncillarySuggestResponseBody {
	return s.Body
}

func (s *AncillarySuggestResponse) SetHeaders(v map[string]*string) *AncillarySuggestResponse {
	s.Headers = v
	return s
}

func (s *AncillarySuggestResponse) SetStatusCode(v int32) *AncillarySuggestResponse {
	s.StatusCode = &v
	return s
}

func (s *AncillarySuggestResponse) SetBody(v *AncillarySuggestResponseBody) *AncillarySuggestResponse {
	s.Body = v
	return s
}

func (s *AncillarySuggestResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
