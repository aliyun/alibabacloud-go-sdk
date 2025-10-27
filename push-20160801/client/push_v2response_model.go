// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PushV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PushV2Response
	GetStatusCode() *int32
	SetBody(v *PushV2ResponseBody) *PushV2Response
	GetBody() *PushV2ResponseBody
}

type PushV2Response struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushV2Response) String() string {
	return dara.Prettify(s)
}

func (s PushV2Response) GoString() string {
	return s.String()
}

func (s *PushV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PushV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PushV2Response) GetBody() *PushV2ResponseBody {
	return s.Body
}

func (s *PushV2Response) SetHeaders(v map[string]*string) *PushV2Response {
	s.Headers = v
	return s
}

func (s *PushV2Response) SetStatusCode(v int32) *PushV2Response {
	s.StatusCode = &v
	return s
}

func (s *PushV2Response) SetBody(v *PushV2ResponseBody) *PushV2Response {
	s.Body = v
	return s
}

func (s *PushV2Response) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
