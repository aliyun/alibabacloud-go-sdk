// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartCallV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartCallV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartCallV2Response
	GetStatusCode() *int32
	SetBody(v *StartCallV2ResponseBody) *StartCallV2Response
	GetBody() *StartCallV2ResponseBody
}

type StartCallV2Response struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartCallV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartCallV2Response) String() string {
	return dara.Prettify(s)
}

func (s StartCallV2Response) GoString() string {
	return s.String()
}

func (s *StartCallV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartCallV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartCallV2Response) GetBody() *StartCallV2ResponseBody {
	return s.Body
}

func (s *StartCallV2Response) SetHeaders(v map[string]*string) *StartCallV2Response {
	s.Headers = v
	return s
}

func (s *StartCallV2Response) SetStatusCode(v int32) *StartCallV2Response {
	s.StatusCode = &v
	return s
}

func (s *StartCallV2Response) SetBody(v *StartCallV2ResponseBody) *StartCallV2Response {
	s.Body = v
	return s
}

func (s *StartCallV2Response) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
