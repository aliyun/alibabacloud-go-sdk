// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushInterventionDictionaryEntriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PushInterventionDictionaryEntriesResponseBody
	GetRequestId() *string
	SetResult(v []*string) *PushInterventionDictionaryEntriesResponseBody
	GetResult() []*string
}

type PushInterventionDictionaryEntriesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D77D0DAF-790D-F5F5-A9C0-133738165014
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned results.
	Result []*string `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s PushInterventionDictionaryEntriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PushInterventionDictionaryEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *PushInterventionDictionaryEntriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PushInterventionDictionaryEntriesResponseBody) GetResult() []*string {
	return s.Result
}

func (s *PushInterventionDictionaryEntriesResponseBody) SetRequestId(v string) *PushInterventionDictionaryEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushInterventionDictionaryEntriesResponseBody) SetResult(v []*string) *PushInterventionDictionaryEntriesResponseBody {
	s.Result = v
	return s
}

func (s *PushInterventionDictionaryEntriesResponseBody) Validate() error {
	return dara.Validate(s)
}
