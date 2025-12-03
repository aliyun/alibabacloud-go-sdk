// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLinkMergeRequestLabelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *LinkMergeRequestLabelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *LinkMergeRequestLabelResponse
	GetStatusCode() *int32
	SetBody(v *LinkMergeRequestLabelResponseBody) *LinkMergeRequestLabelResponse
	GetBody() *LinkMergeRequestLabelResponseBody
}

type LinkMergeRequestLabelResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LinkMergeRequestLabelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LinkMergeRequestLabelResponse) String() string {
	return dara.Prettify(s)
}

func (s LinkMergeRequestLabelResponse) GoString() string {
	return s.String()
}

func (s *LinkMergeRequestLabelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *LinkMergeRequestLabelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *LinkMergeRequestLabelResponse) GetBody() *LinkMergeRequestLabelResponseBody {
	return s.Body
}

func (s *LinkMergeRequestLabelResponse) SetHeaders(v map[string]*string) *LinkMergeRequestLabelResponse {
	s.Headers = v
	return s
}

func (s *LinkMergeRequestLabelResponse) SetStatusCode(v int32) *LinkMergeRequestLabelResponse {
	s.StatusCode = &v
	return s
}

func (s *LinkMergeRequestLabelResponse) SetBody(v *LinkMergeRequestLabelResponseBody) *LinkMergeRequestLabelResponse {
	s.Body = v
	return s
}

func (s *LinkMergeRequestLabelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
