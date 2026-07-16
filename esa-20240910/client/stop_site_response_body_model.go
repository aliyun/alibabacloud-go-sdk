// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopSiteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopSiteResponseBody
	GetRequestId() *string
}

type StopSiteResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopSiteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopSiteResponseBody) GoString() string {
	return s.String()
}

func (s *StopSiteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopSiteResponseBody) SetRequestId(v string) *StopSiteResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopSiteResponseBody) Validate() error {
	return dara.Validate(s)
}
