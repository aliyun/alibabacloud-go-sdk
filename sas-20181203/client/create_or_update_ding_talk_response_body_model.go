// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrUpdateDingTalkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateOrUpdateDingTalkResponseBody
	GetRequestId() *string
}

type CreateOrUpdateDingTalkResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 76975B7A-34DC-5CB6-9538-91700D4F112E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateOrUpdateDingTalkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateDingTalkResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateDingTalkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOrUpdateDingTalkResponseBody) SetRequestId(v string) *CreateOrUpdateDingTalkResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOrUpdateDingTalkResponseBody) Validate() error {
	return dara.Validate(s)
}
