// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGraphDraftAssembledResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetGraphDraftAssembledResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetGraphDraftAssembledResponse
	GetStatusCode() *int32
	SetBody(v *GetGraphDraftAssembledResponseBody) *GetGraphDraftAssembledResponse
	GetBody() *GetGraphDraftAssembledResponseBody
}

type GetGraphDraftAssembledResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGraphDraftAssembledResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGraphDraftAssembledResponse) String() string {
	return dara.Prettify(s)
}

func (s GetGraphDraftAssembledResponse) GoString() string {
	return s.String()
}

func (s *GetGraphDraftAssembledResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetGraphDraftAssembledResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetGraphDraftAssembledResponse) GetBody() *GetGraphDraftAssembledResponseBody {
	return s.Body
}

func (s *GetGraphDraftAssembledResponse) SetHeaders(v map[string]*string) *GetGraphDraftAssembledResponse {
	s.Headers = v
	return s
}

func (s *GetGraphDraftAssembledResponse) SetStatusCode(v int32) *GetGraphDraftAssembledResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGraphDraftAssembledResponse) SetBody(v *GetGraphDraftAssembledResponseBody) *GetGraphDraftAssembledResponse {
	s.Body = v
	return s
}

func (s *GetGraphDraftAssembledResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
