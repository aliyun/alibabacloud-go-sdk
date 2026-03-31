// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveDataSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveDataSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveDataSourceResponse
	GetStatusCode() *int32
	SetBody(v *RemoveDataSourceResponseBody) *RemoveDataSourceResponse
	GetBody() *RemoveDataSourceResponseBody
}

type RemoveDataSourceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveDataSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveDataSourceResponse) GoString() string {
	return s.String()
}

func (s *RemoveDataSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveDataSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveDataSourceResponse) GetBody() *RemoveDataSourceResponseBody {
	return s.Body
}

func (s *RemoveDataSourceResponse) SetHeaders(v map[string]*string) *RemoveDataSourceResponse {
	s.Headers = v
	return s
}

func (s *RemoveDataSourceResponse) SetStatusCode(v int32) *RemoveDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveDataSourceResponse) SetBody(v *RemoveDataSourceResponseBody) *RemoveDataSourceResponse {
	s.Body = v
	return s
}

func (s *RemoveDataSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
