// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHandleSimilarSecurityEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *HandleSimilarSecurityEventsResponseBody
	GetRequestId() *string
}

type HandleSimilarSecurityEventsResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// A3653911-33A6-5268-8B91-7690471F7AA1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s HandleSimilarSecurityEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s HandleSimilarSecurityEventsResponseBody) GoString() string {
	return s.String()
}

func (s *HandleSimilarSecurityEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *HandleSimilarSecurityEventsResponseBody) SetRequestId(v string) *HandleSimilarSecurityEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *HandleSimilarSecurityEventsResponseBody) Validate() error {
	return dara.Validate(s)
}
