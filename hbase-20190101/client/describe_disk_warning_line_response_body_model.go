// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiskWarningLineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDiskWarningLineResponseBody
	GetRequestId() *string
	SetWarningLine(v string) *DescribeDiskWarningLineResponseBody
	GetWarningLine() *string
}

type DescribeDiskWarningLineResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 08DF8283-D290-4107-931E-7913D6D3480D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The disk capacity alert threshold. For example, if the value is 75, an alert is triggered when disk usage exceeds 75%. If no value is returned, the user has not configured this parameter, and the system default value is 80%.
	//
	// example:
	//
	// 80
	WarningLine *string `json:"WarningLine,omitempty" xml:"WarningLine,omitempty"`
}

func (s DescribeDiskWarningLineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiskWarningLineResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiskWarningLineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDiskWarningLineResponseBody) GetWarningLine() *string {
	return s.WarningLine
}

func (s *DescribeDiskWarningLineResponseBody) SetRequestId(v string) *DescribeDiskWarningLineResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiskWarningLineResponseBody) SetWarningLine(v string) *DescribeDiskWarningLineResponseBody {
	s.WarningLine = &v
	return s
}

func (s *DescribeDiskWarningLineResponseBody) Validate() error {
	return dara.Validate(s)
}
