// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddApListToApgroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddApListToApgroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddApListToApgroupResponse
	GetStatusCode() *int32
	SetBody(v *AddApListToApgroupResponseBody) *AddApListToApgroupResponse
	GetBody() *AddApListToApgroupResponseBody
}

type AddApListToApgroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddApListToApgroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddApListToApgroupResponse) String() string {
	return dara.Prettify(s)
}

func (s AddApListToApgroupResponse) GoString() string {
	return s.String()
}

func (s *AddApListToApgroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddApListToApgroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddApListToApgroupResponse) GetBody() *AddApListToApgroupResponseBody {
	return s.Body
}

func (s *AddApListToApgroupResponse) SetHeaders(v map[string]*string) *AddApListToApgroupResponse {
	s.Headers = v
	return s
}

func (s *AddApListToApgroupResponse) SetStatusCode(v int32) *AddApListToApgroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AddApListToApgroupResponse) SetBody(v *AddApListToApgroupResponseBody) *AddApListToApgroupResponse {
	s.Body = v
	return s
}

func (s *AddApListToApgroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
