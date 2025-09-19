// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSafConsoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSafConsoleResponseBody
	GetRequestId() *string
	SetBizData(v []*string) *DescribeSafConsoleResponseBody
	GetBizData() []*string
}

type DescribeSafConsoleResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned result.
	BizData []*string `json:"bizData,omitempty" xml:"bizData,omitempty" type:"Repeated"`
}

func (s DescribeSafConsoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSafConsoleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSafConsoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSafConsoleResponseBody) GetBizData() []*string {
	return s.BizData
}

func (s *DescribeSafConsoleResponseBody) SetRequestId(v string) *DescribeSafConsoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSafConsoleResponseBody) SetBizData(v []*string) *DescribeSafConsoleResponseBody {
	s.BizData = v
	return s
}

func (s *DescribeSafConsoleResponseBody) Validate() error {
	return dara.Validate(s)
}
