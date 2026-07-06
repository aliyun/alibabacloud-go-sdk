// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFaceRecordV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFaceRecordV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFaceRecordV2Response
	GetStatusCode() *int32
	SetBody(v *DeleteFaceRecordV2ResponseBody) *DeleteFaceRecordV2Response
	GetBody() *DeleteFaceRecordV2ResponseBody
}

type DeleteFaceRecordV2Response struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFaceRecordV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFaceRecordV2Response) String() string {
	return dara.Prettify(s)
}

func (s DeleteFaceRecordV2Response) GoString() string {
	return s.String()
}

func (s *DeleteFaceRecordV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFaceRecordV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFaceRecordV2Response) GetBody() *DeleteFaceRecordV2ResponseBody {
	return s.Body
}

func (s *DeleteFaceRecordV2Response) SetHeaders(v map[string]*string) *DeleteFaceRecordV2Response {
	s.Headers = v
	return s
}

func (s *DeleteFaceRecordV2Response) SetStatusCode(v int32) *DeleteFaceRecordV2Response {
	s.StatusCode = &v
	return s
}

func (s *DeleteFaceRecordV2Response) SetBody(v *DeleteFaceRecordV2ResponseBody) *DeleteFaceRecordV2Response {
	s.Body = v
	return s
}

func (s *DeleteFaceRecordV2Response) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
