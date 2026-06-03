// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCommonStatisticPreviewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryCommonStatisticPreviewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryCommonStatisticPreviewResponse
	GetStatusCode() *int32
	SetBody(v *QueryCommonStatisticPreviewResponseBody) *QueryCommonStatisticPreviewResponse
	GetBody() *QueryCommonStatisticPreviewResponseBody
}

type QueryCommonStatisticPreviewResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCommonStatisticPreviewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCommonStatisticPreviewResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryCommonStatisticPreviewResponse) GoString() string {
	return s.String()
}

func (s *QueryCommonStatisticPreviewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryCommonStatisticPreviewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryCommonStatisticPreviewResponse) GetBody() *QueryCommonStatisticPreviewResponseBody {
	return s.Body
}

func (s *QueryCommonStatisticPreviewResponse) SetHeaders(v map[string]*string) *QueryCommonStatisticPreviewResponse {
	s.Headers = v
	return s
}

func (s *QueryCommonStatisticPreviewResponse) SetStatusCode(v int32) *QueryCommonStatisticPreviewResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCommonStatisticPreviewResponse) SetBody(v *QueryCommonStatisticPreviewResponseBody) *QueryCommonStatisticPreviewResponse {
	s.Body = v
	return s
}

func (s *QueryCommonStatisticPreviewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
