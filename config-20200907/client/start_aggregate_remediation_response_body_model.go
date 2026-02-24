// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartAggregateRemediationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *StartAggregateRemediationResponseBody
	GetData() *bool
	SetRequestId(v string) *StartAggregateRemediationResponseBody
	GetRequestId() *string
}

type StartAggregateRemediationResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartAggregateRemediationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartAggregateRemediationResponseBody) GoString() string {
	return s.String()
}

func (s *StartAggregateRemediationResponseBody) GetData() *bool {
	return s.Data
}

func (s *StartAggregateRemediationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartAggregateRemediationResponseBody) SetData(v bool) *StartAggregateRemediationResponseBody {
	s.Data = &v
	return s
}

func (s *StartAggregateRemediationResponseBody) SetRequestId(v string) *StartAggregateRemediationResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartAggregateRemediationResponseBody) Validate() error {
	return dara.Validate(s)
}
