// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCopyrightExtractJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryCopyrightExtractJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryCopyrightExtractJobResponse
	GetStatusCode() *int32
	SetBody(v *QueryCopyrightExtractJobResponseBody) *QueryCopyrightExtractJobResponse
	GetBody() *QueryCopyrightExtractJobResponseBody
}

type QueryCopyrightExtractJobResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCopyrightExtractJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCopyrightExtractJobResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryCopyrightExtractJobResponse) GoString() string {
	return s.String()
}

func (s *QueryCopyrightExtractJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryCopyrightExtractJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryCopyrightExtractJobResponse) GetBody() *QueryCopyrightExtractJobResponseBody {
	return s.Body
}

func (s *QueryCopyrightExtractJobResponse) SetHeaders(v map[string]*string) *QueryCopyrightExtractJobResponse {
	s.Headers = v
	return s
}

func (s *QueryCopyrightExtractJobResponse) SetStatusCode(v int32) *QueryCopyrightExtractJobResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCopyrightExtractJobResponse) SetBody(v *QueryCopyrightExtractJobResponseBody) *QueryCopyrightExtractJobResponse {
	s.Body = v
	return s
}

func (s *QueryCopyrightExtractJobResponse) Validate() error {
	return dara.Validate(s)
}
