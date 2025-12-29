// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTagInfoBySelectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryTagInfoBySelectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryTagInfoBySelectionResponse
	GetStatusCode() *int32
	SetBody(v *QueryTagInfoBySelectionResponseBody) *QueryTagInfoBySelectionResponse
	GetBody() *QueryTagInfoBySelectionResponseBody
}

type QueryTagInfoBySelectionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTagInfoBySelectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTagInfoBySelectionResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryTagInfoBySelectionResponse) GoString() string {
	return s.String()
}

func (s *QueryTagInfoBySelectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryTagInfoBySelectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryTagInfoBySelectionResponse) GetBody() *QueryTagInfoBySelectionResponseBody {
	return s.Body
}

func (s *QueryTagInfoBySelectionResponse) SetHeaders(v map[string]*string) *QueryTagInfoBySelectionResponse {
	s.Headers = v
	return s
}

func (s *QueryTagInfoBySelectionResponse) SetStatusCode(v int32) *QueryTagInfoBySelectionResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTagInfoBySelectionResponse) SetBody(v *QueryTagInfoBySelectionResponseBody) *QueryTagInfoBySelectionResponse {
	s.Body = v
	return s
}

func (s *QueryTagInfoBySelectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
