// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveGraphDraftBatchDefineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveGraphDraftBatchDefineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveGraphDraftBatchDefineResponse
	GetStatusCode() *int32
	SetBody(v *SaveGraphDraftBatchDefineResponseBody) *SaveGraphDraftBatchDefineResponse
	GetBody() *SaveGraphDraftBatchDefineResponseBody
}

type SaveGraphDraftBatchDefineResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveGraphDraftBatchDefineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveGraphDraftBatchDefineResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveGraphDraftBatchDefineResponse) GoString() string {
	return s.String()
}

func (s *SaveGraphDraftBatchDefineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveGraphDraftBatchDefineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveGraphDraftBatchDefineResponse) GetBody() *SaveGraphDraftBatchDefineResponseBody {
	return s.Body
}

func (s *SaveGraphDraftBatchDefineResponse) SetHeaders(v map[string]*string) *SaveGraphDraftBatchDefineResponse {
	s.Headers = v
	return s
}

func (s *SaveGraphDraftBatchDefineResponse) SetStatusCode(v int32) *SaveGraphDraftBatchDefineResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveGraphDraftBatchDefineResponse) SetBody(v *SaveGraphDraftBatchDefineResponseBody) *SaveGraphDraftBatchDefineResponse {
	s.Body = v
	return s
}

func (s *SaveGraphDraftBatchDefineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
