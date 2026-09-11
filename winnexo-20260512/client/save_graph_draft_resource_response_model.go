// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveGraphDraftResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveGraphDraftResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveGraphDraftResourceResponse
	GetStatusCode() *int32
	SetBody(v *SaveGraphDraftResourceResponseBody) *SaveGraphDraftResourceResponse
	GetBody() *SaveGraphDraftResourceResponseBody
}

type SaveGraphDraftResourceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveGraphDraftResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveGraphDraftResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveGraphDraftResourceResponse) GoString() string {
	return s.String()
}

func (s *SaveGraphDraftResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveGraphDraftResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveGraphDraftResourceResponse) GetBody() *SaveGraphDraftResourceResponseBody {
	return s.Body
}

func (s *SaveGraphDraftResourceResponse) SetHeaders(v map[string]*string) *SaveGraphDraftResourceResponse {
	s.Headers = v
	return s
}

func (s *SaveGraphDraftResourceResponse) SetStatusCode(v int32) *SaveGraphDraftResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveGraphDraftResourceResponse) SetBody(v *SaveGraphDraftResourceResponseBody) *SaveGraphDraftResourceResponse {
	s.Body = v
	return s
}

func (s *SaveGraphDraftResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
