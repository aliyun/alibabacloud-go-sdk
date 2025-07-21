// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendTestByTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SendTestByTemplateResponseBody
	GetRequestId() *string
}

type SendTestByTemplateResponseBody struct {
	// Request ID
	//
	// example:
	//
	// 10A1AD70-E48E-476D-98D9-39BD92193837
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendTestByTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendTestByTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *SendTestByTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendTestByTemplateResponseBody) SetRequestId(v string) *SendTestByTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendTestByTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
