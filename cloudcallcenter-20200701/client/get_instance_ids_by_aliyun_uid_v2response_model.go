// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceIdsByAliyunUidV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceIdsByAliyunUidV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceIdsByAliyunUidV2Response
	GetStatusCode() *int32
	SetBody(v *GetInstanceIdsByAliyunUidV2ResponseBody) *GetInstanceIdsByAliyunUidV2Response
	GetBody() *GetInstanceIdsByAliyunUidV2ResponseBody
}

type GetInstanceIdsByAliyunUidV2Response struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceIdsByAliyunUidV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceIdsByAliyunUidV2Response) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceIdsByAliyunUidV2Response) GoString() string {
	return s.String()
}

func (s *GetInstanceIdsByAliyunUidV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceIdsByAliyunUidV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceIdsByAliyunUidV2Response) GetBody() *GetInstanceIdsByAliyunUidV2ResponseBody {
	return s.Body
}

func (s *GetInstanceIdsByAliyunUidV2Response) SetHeaders(v map[string]*string) *GetInstanceIdsByAliyunUidV2Response {
	s.Headers = v
	return s
}

func (s *GetInstanceIdsByAliyunUidV2Response) SetStatusCode(v int32) *GetInstanceIdsByAliyunUidV2Response {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceIdsByAliyunUidV2Response) SetBody(v *GetInstanceIdsByAliyunUidV2ResponseBody) *GetInstanceIdsByAliyunUidV2Response {
	s.Body = v
	return s
}

func (s *GetInstanceIdsByAliyunUidV2Response) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
