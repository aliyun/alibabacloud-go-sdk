// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWebCCRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateWebCCRuleResponseBody
	GetRequestId() *string
}

type CreateWebCCRuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0bcf28g5-d57c-11e7-9bs0-d89d6717dxbc
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateWebCCRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWebCCRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWebCCRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWebCCRuleResponseBody) SetRequestId(v string) *CreateWebCCRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWebCCRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
