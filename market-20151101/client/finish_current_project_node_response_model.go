// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFinishCurrentProjectNodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FinishCurrentProjectNodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FinishCurrentProjectNodeResponse
	GetStatusCode() *int32
	SetBody(v *FinishCurrentProjectNodeResponseBody) *FinishCurrentProjectNodeResponse
	GetBody() *FinishCurrentProjectNodeResponseBody
}

type FinishCurrentProjectNodeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FinishCurrentProjectNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FinishCurrentProjectNodeResponse) String() string {
	return dara.Prettify(s)
}

func (s FinishCurrentProjectNodeResponse) GoString() string {
	return s.String()
}

func (s *FinishCurrentProjectNodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FinishCurrentProjectNodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FinishCurrentProjectNodeResponse) GetBody() *FinishCurrentProjectNodeResponseBody {
	return s.Body
}

func (s *FinishCurrentProjectNodeResponse) SetHeaders(v map[string]*string) *FinishCurrentProjectNodeResponse {
	s.Headers = v
	return s
}

func (s *FinishCurrentProjectNodeResponse) SetStatusCode(v int32) *FinishCurrentProjectNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *FinishCurrentProjectNodeResponse) SetBody(v *FinishCurrentProjectNodeResponseBody) *FinishCurrentProjectNodeResponse {
	s.Body = v
	return s
}

func (s *FinishCurrentProjectNodeResponse) Validate() error {
	return dara.Validate(s)
}
