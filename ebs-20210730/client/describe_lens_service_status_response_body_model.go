// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLensServiceStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeLensServiceStatusResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeLensServiceStatusResponseBody
	GetStatus() *string
}

type DescribeLensServiceStatusResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The state of CloudLens for EBS. Valid values:
	//
	// 	- Applying
	//
	// 	- UnAvailable
	//
	// 	- Available
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeLensServiceStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLensServiceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLensServiceStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLensServiceStatusResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeLensServiceStatusResponseBody) SetRequestId(v string) *DescribeLensServiceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLensServiceStatusResponseBody) SetStatus(v string) *DescribeLensServiceStatusResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeLensServiceStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
