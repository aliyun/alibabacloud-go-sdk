// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPriceV1Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryPriceV1Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryPriceV1Response
	GetStatusCode() *int32
	SetBody(v *QueryPriceV1ResponseBody) *QueryPriceV1Response
	GetBody() *QueryPriceV1ResponseBody
}

type QueryPriceV1Response struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPriceV1ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPriceV1Response) String() string {
	return dara.Prettify(s)
}

func (s QueryPriceV1Response) GoString() string {
	return s.String()
}

func (s *QueryPriceV1Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryPriceV1Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryPriceV1Response) GetBody() *QueryPriceV1ResponseBody {
	return s.Body
}

func (s *QueryPriceV1Response) SetHeaders(v map[string]*string) *QueryPriceV1Response {
	s.Headers = v
	return s
}

func (s *QueryPriceV1Response) SetStatusCode(v int32) *QueryPriceV1Response {
	s.StatusCode = &v
	return s
}

func (s *QueryPriceV1Response) SetBody(v *QueryPriceV1ResponseBody) *QueryPriceV1Response {
	s.Body = v
	return s
}

func (s *QueryPriceV1Response) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
