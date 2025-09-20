// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTriggersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListTriggersResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTriggersResponseBody
	GetRequestId() *string
	SetTriggers(v []*DataIngestion) *ListTriggersResponseBody
	GetTriggers() []*DataIngestion
}

type ListTriggersResponseBody struct {
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// If NextToken is empty, no next page exists.
	//
	// example:
	//
	// MTIzNDU2Nzg6aW1tdGVzdDpleGFtcGxlYnVja2V0OmRhdGFzZXQwMDE6b3NzOi8vZXhhbXBsZWJ1Y2tldC9zYW1wbGVvYmplY3QxLmpwZw==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F480BFAF-E778-5079-93AD-1E4631******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The triggers.
	Triggers []*DataIngestion `json:"Triggers,omitempty" xml:"Triggers,omitempty" type:"Repeated"`
}

func (s ListTriggersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTriggersResponseBody) GoString() string {
	return s.String()
}

func (s *ListTriggersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTriggersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTriggersResponseBody) GetTriggers() []*DataIngestion {
	return s.Triggers
}

func (s *ListTriggersResponseBody) SetNextToken(v string) *ListTriggersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTriggersResponseBody) SetRequestId(v string) *ListTriggersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTriggersResponseBody) SetTriggers(v []*DataIngestion) *ListTriggersResponseBody {
	s.Triggers = v
	return s
}

func (s *ListTriggersResponseBody) Validate() error {
	return dara.Validate(s)
}
