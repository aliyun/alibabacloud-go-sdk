// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMergeRequestResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMergeRequestResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMergeRequestResponse
	GetStatusCode() *int32
	SetBody(v *GetMergeRequestResponseBody) *GetMergeRequestResponse
	GetBody() *GetMergeRequestResponseBody
}

type GetMergeRequestResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMergeRequestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMergeRequestResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMergeRequestResponse) GoString() string {
	return s.String()
}

func (s *GetMergeRequestResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMergeRequestResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMergeRequestResponse) GetBody() *GetMergeRequestResponseBody {
	return s.Body
}

func (s *GetMergeRequestResponse) SetHeaders(v map[string]*string) *GetMergeRequestResponse {
	s.Headers = v
	return s
}

func (s *GetMergeRequestResponse) SetStatusCode(v int32) *GetMergeRequestResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMergeRequestResponse) SetBody(v *GetMergeRequestResponseBody) *GetMergeRequestResponse {
	s.Body = v
	return s
}

func (s *GetMergeRequestResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
