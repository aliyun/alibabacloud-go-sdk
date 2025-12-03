// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMergeRequestChangeTreeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMergeRequestChangeTreeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMergeRequestChangeTreeResponse
	GetStatusCode() *int32
	SetBody(v *GetMergeRequestChangeTreeResponseBody) *GetMergeRequestChangeTreeResponse
	GetBody() *GetMergeRequestChangeTreeResponseBody
}

type GetMergeRequestChangeTreeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMergeRequestChangeTreeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMergeRequestChangeTreeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMergeRequestChangeTreeResponse) GoString() string {
	return s.String()
}

func (s *GetMergeRequestChangeTreeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMergeRequestChangeTreeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMergeRequestChangeTreeResponse) GetBody() *GetMergeRequestChangeTreeResponseBody {
	return s.Body
}

func (s *GetMergeRequestChangeTreeResponse) SetHeaders(v map[string]*string) *GetMergeRequestChangeTreeResponse {
	s.Headers = v
	return s
}

func (s *GetMergeRequestChangeTreeResponse) SetStatusCode(v int32) *GetMergeRequestChangeTreeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMergeRequestChangeTreeResponse) SetBody(v *GetMergeRequestChangeTreeResponseBody) *GetMergeRequestChangeTreeResponse {
	s.Body = v
	return s
}

func (s *GetMergeRequestChangeTreeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
