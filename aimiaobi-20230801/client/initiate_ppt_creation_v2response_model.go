// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitiatePptCreationV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InitiatePptCreationV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InitiatePptCreationV2Response
	GetStatusCode() *int32
	SetBody(v *InitiatePptCreationV2ResponseBody) *InitiatePptCreationV2Response
	GetBody() *InitiatePptCreationV2ResponseBody
}

type InitiatePptCreationV2Response struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InitiatePptCreationV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InitiatePptCreationV2Response) String() string {
	return dara.Prettify(s)
}

func (s InitiatePptCreationV2Response) GoString() string {
	return s.String()
}

func (s *InitiatePptCreationV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InitiatePptCreationV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InitiatePptCreationV2Response) GetBody() *InitiatePptCreationV2ResponseBody {
	return s.Body
}

func (s *InitiatePptCreationV2Response) SetHeaders(v map[string]*string) *InitiatePptCreationV2Response {
	s.Headers = v
	return s
}

func (s *InitiatePptCreationV2Response) SetStatusCode(v int32) *InitiatePptCreationV2Response {
	s.StatusCode = &v
	return s
}

func (s *InitiatePptCreationV2Response) SetBody(v *InitiatePptCreationV2ResponseBody) *InitiatePptCreationV2Response {
	s.Body = v
	return s
}

func (s *InitiatePptCreationV2Response) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
