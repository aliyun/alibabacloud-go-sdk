// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAxbBindFixedLineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAxbBindFixedLineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAxbBindFixedLineResponse
	GetStatusCode() *int32
	SetBody(v *QueryAxbBindFixedLineResponseBody) *QueryAxbBindFixedLineResponse
	GetBody() *QueryAxbBindFixedLineResponseBody
}

type QueryAxbBindFixedLineResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAxbBindFixedLineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAxbBindFixedLineResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAxbBindFixedLineResponse) GoString() string {
	return s.String()
}

func (s *QueryAxbBindFixedLineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAxbBindFixedLineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAxbBindFixedLineResponse) GetBody() *QueryAxbBindFixedLineResponseBody {
	return s.Body
}

func (s *QueryAxbBindFixedLineResponse) SetHeaders(v map[string]*string) *QueryAxbBindFixedLineResponse {
	s.Headers = v
	return s
}

func (s *QueryAxbBindFixedLineResponse) SetStatusCode(v int32) *QueryAxbBindFixedLineResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAxbBindFixedLineResponse) SetBody(v *QueryAxbBindFixedLineResponseBody) *QueryAxbBindFixedLineResponse {
	s.Body = v
	return s
}

func (s *QueryAxbBindFixedLineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
