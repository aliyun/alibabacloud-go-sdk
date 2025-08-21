// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWebRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateWebRuleResponseBody
	GetRequestId() *string
}

type CreateWebRuleResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 9EC62E89-BD30-4FCD-9CB8-FA53865FF0D7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateWebRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWebRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWebRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWebRuleResponseBody) SetRequestId(v string) *CreateWebRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWebRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
