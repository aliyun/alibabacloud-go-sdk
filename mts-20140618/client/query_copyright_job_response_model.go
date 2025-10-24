// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCopyrightJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryCopyrightJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryCopyrightJobResponse
	GetStatusCode() *int32
	SetBody(v *QueryCopyrightJobResponseBody) *QueryCopyrightJobResponse
	GetBody() *QueryCopyrightJobResponseBody
}

type QueryCopyrightJobResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCopyrightJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCopyrightJobResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryCopyrightJobResponse) GoString() string {
	return s.String()
}

func (s *QueryCopyrightJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryCopyrightJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryCopyrightJobResponse) GetBody() *QueryCopyrightJobResponseBody {
	return s.Body
}

func (s *QueryCopyrightJobResponse) SetHeaders(v map[string]*string) *QueryCopyrightJobResponse {
	s.Headers = v
	return s
}

func (s *QueryCopyrightJobResponse) SetStatusCode(v int32) *QueryCopyrightJobResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCopyrightJobResponse) SetBody(v *QueryCopyrightJobResponseBody) *QueryCopyrightJobResponse {
	s.Body = v
	return s
}

func (s *QueryCopyrightJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
