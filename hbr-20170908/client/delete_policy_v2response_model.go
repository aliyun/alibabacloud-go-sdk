// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolicyV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePolicyV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePolicyV2Response
	GetStatusCode() *int32
	SetBody(v *DeletePolicyV2ResponseBody) *DeletePolicyV2Response
	GetBody() *DeletePolicyV2ResponseBody
}

type DeletePolicyV2Response struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePolicyV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePolicyV2Response) String() string {
	return dara.Prettify(s)
}

func (s DeletePolicyV2Response) GoString() string {
	return s.String()
}

func (s *DeletePolicyV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePolicyV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePolicyV2Response) GetBody() *DeletePolicyV2ResponseBody {
	return s.Body
}

func (s *DeletePolicyV2Response) SetHeaders(v map[string]*string) *DeletePolicyV2Response {
	s.Headers = v
	return s
}

func (s *DeletePolicyV2Response) SetStatusCode(v int32) *DeletePolicyV2Response {
	s.StatusCode = &v
	return s
}

func (s *DeletePolicyV2Response) SetBody(v *DeletePolicyV2ResponseBody) *DeletePolicyV2Response {
	s.Body = v
	return s
}

func (s *DeletePolicyV2Response) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
