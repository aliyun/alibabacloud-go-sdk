// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevertGraphDraftResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevertGraphDraftResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevertGraphDraftResourceResponse
	GetStatusCode() *int32
	SetBody(v *RevertGraphDraftResourceResponseBody) *RevertGraphDraftResourceResponse
	GetBody() *RevertGraphDraftResourceResponseBody
}

type RevertGraphDraftResourceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevertGraphDraftResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevertGraphDraftResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s RevertGraphDraftResourceResponse) GoString() string {
	return s.String()
}

func (s *RevertGraphDraftResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevertGraphDraftResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevertGraphDraftResourceResponse) GetBody() *RevertGraphDraftResourceResponseBody {
	return s.Body
}

func (s *RevertGraphDraftResourceResponse) SetHeaders(v map[string]*string) *RevertGraphDraftResourceResponse {
	s.Headers = v
	return s
}

func (s *RevertGraphDraftResourceResponse) SetStatusCode(v int32) *RevertGraphDraftResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *RevertGraphDraftResourceResponse) SetBody(v *RevertGraphDraftResourceResponseBody) *RevertGraphDraftResourceResponse {
	s.Body = v
	return s
}

func (s *RevertGraphDraftResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
