// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMassPushV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MassPushV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MassPushV2Response
	GetStatusCode() *int32
	SetBody(v *MassPushV2ResponseBody) *MassPushV2Response
	GetBody() *MassPushV2ResponseBody
}

type MassPushV2Response struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MassPushV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MassPushV2Response) String() string {
	return dara.Prettify(s)
}

func (s MassPushV2Response) GoString() string {
	return s.String()
}

func (s *MassPushV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MassPushV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MassPushV2Response) GetBody() *MassPushV2ResponseBody {
	return s.Body
}

func (s *MassPushV2Response) SetHeaders(v map[string]*string) *MassPushV2Response {
	s.Headers = v
	return s
}

func (s *MassPushV2Response) SetStatusCode(v int32) *MassPushV2Response {
	s.StatusCode = &v
	return s
}

func (s *MassPushV2Response) SetBody(v *MassPushV2ResponseBody) *MassPushV2Response {
	s.Body = v
	return s
}

func (s *MassPushV2Response) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
