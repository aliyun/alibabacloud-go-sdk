// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfirmVirusEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ConfirmVirusEventsResponseBody
	GetRequestId() *string
	SetResult(v bool) *ConfirmVirusEventsResponseBody
	GetResult() *bool
}

type ConfirmVirusEventsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5DFD6277-CC36-57F7-ACE6-F5952XXXXXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ConfirmVirusEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfirmVirusEventsResponseBody) GoString() string {
	return s.String()
}

func (s *ConfirmVirusEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfirmVirusEventsResponseBody) GetResult() *bool {
	return s.Result
}

func (s *ConfirmVirusEventsResponseBody) SetRequestId(v string) *ConfirmVirusEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfirmVirusEventsResponseBody) SetResult(v bool) *ConfirmVirusEventsResponseBody {
	s.Result = &v
	return s
}

func (s *ConfirmVirusEventsResponseBody) Validate() error {
	return dara.Validate(s)
}
