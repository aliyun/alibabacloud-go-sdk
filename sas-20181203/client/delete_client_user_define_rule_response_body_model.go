// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClientUserDefineRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteClientUserDefineRuleResponseBody
	GetRequestId() *string
}

type DeleteClientUserDefineRuleResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 4552b59b-18f2-4fad-b6a2-0d59b8f2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteClientUserDefineRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteClientUserDefineRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteClientUserDefineRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteClientUserDefineRuleResponseBody) SetRequestId(v string) *DeleteClientUserDefineRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteClientUserDefineRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
