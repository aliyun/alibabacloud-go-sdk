// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloneSentinelRuleFromAhasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v map[string][]*string) *CloneSentinelRuleFromAhasResponseBody
	GetData() map[string][]*string
	SetRequestId(v string) *CloneSentinelRuleFromAhasResponseBody
	GetRequestId() *string
}

type CloneSentinelRuleFromAhasResponseBody struct {
	// The returned data.
	Data map[string][]*string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EE5C32A1-BC0E-4B79-817C-103E4EDF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloneSentinelRuleFromAhasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloneSentinelRuleFromAhasResponseBody) GoString() string {
	return s.String()
}

func (s *CloneSentinelRuleFromAhasResponseBody) GetData() map[string][]*string {
	return s.Data
}

func (s *CloneSentinelRuleFromAhasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloneSentinelRuleFromAhasResponseBody) SetData(v map[string][]*string) *CloneSentinelRuleFromAhasResponseBody {
	s.Data = v
	return s
}

func (s *CloneSentinelRuleFromAhasResponseBody) SetRequestId(v string) *CloneSentinelRuleFromAhasResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloneSentinelRuleFromAhasResponseBody) Validate() error {
	return dara.Validate(s)
}
