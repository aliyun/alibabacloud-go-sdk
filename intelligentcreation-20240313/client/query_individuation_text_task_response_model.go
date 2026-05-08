// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryIndividuationTextTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryIndividuationTextTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryIndividuationTextTaskResponse
	GetStatusCode() *int32
	SetBody(v *QueryIndividuationTextTaskResponseBody) *QueryIndividuationTextTaskResponse
	GetBody() *QueryIndividuationTextTaskResponseBody
}

type QueryIndividuationTextTaskResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryIndividuationTextTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryIndividuationTextTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryIndividuationTextTaskResponse) GoString() string {
	return s.String()
}

func (s *QueryIndividuationTextTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryIndividuationTextTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryIndividuationTextTaskResponse) GetBody() *QueryIndividuationTextTaskResponseBody {
	return s.Body
}

func (s *QueryIndividuationTextTaskResponse) SetHeaders(v map[string]*string) *QueryIndividuationTextTaskResponse {
	s.Headers = v
	return s
}

func (s *QueryIndividuationTextTaskResponse) SetStatusCode(v int32) *QueryIndividuationTextTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryIndividuationTextTaskResponse) SetBody(v *QueryIndividuationTextTaskResponseBody) *QueryIndividuationTextTaskResponse {
	s.Body = v
	return s
}

func (s *QueryIndividuationTextTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
