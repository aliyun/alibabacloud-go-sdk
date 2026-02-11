// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendMseIncidentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SendMseIncidentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SendMseIncidentResponseBody
	GetSuccess() *bool
}

type SendMseIncidentResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendMseIncidentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendMseIncidentResponseBody) GoString() string {
	return s.String()
}

func (s *SendMseIncidentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendMseIncidentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SendMseIncidentResponseBody) SetRequestId(v string) *SendMseIncidentResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendMseIncidentResponseBody) SetSuccess(v bool) *SendMseIncidentResponseBody {
	s.Success = &v
	return s
}

func (s *SendMseIncidentResponseBody) Validate() error {
	return dara.Validate(s)
}
