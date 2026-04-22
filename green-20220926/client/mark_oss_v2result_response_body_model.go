// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMarkOssV2ResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailureRequestIds(v []*string) *MarkOssV2ResultResponseBody
	GetFailureRequestIds() []*string
	SetRequestId(v string) *MarkOssV2ResultResponseBody
	GetRequestId() *string
	SetSuccessRequestIds(v []*string) *MarkOssV2ResultResponseBody
	GetSuccessRequestIds() []*string
}

type MarkOssV2ResultResponseBody struct {
	FailureRequestIds []*string `json:"FailureRequestIds,omitempty" xml:"FailureRequestIds,omitempty" type:"Repeated"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId         *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuccessRequestIds []*string `json:"SuccessRequestIds,omitempty" xml:"SuccessRequestIds,omitempty" type:"Repeated"`
}

func (s MarkOssV2ResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MarkOssV2ResultResponseBody) GoString() string {
	return s.String()
}

func (s *MarkOssV2ResultResponseBody) GetFailureRequestIds() []*string {
	return s.FailureRequestIds
}

func (s *MarkOssV2ResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MarkOssV2ResultResponseBody) GetSuccessRequestIds() []*string {
	return s.SuccessRequestIds
}

func (s *MarkOssV2ResultResponseBody) SetFailureRequestIds(v []*string) *MarkOssV2ResultResponseBody {
	s.FailureRequestIds = v
	return s
}

func (s *MarkOssV2ResultResponseBody) SetRequestId(v string) *MarkOssV2ResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *MarkOssV2ResultResponseBody) SetSuccessRequestIds(v []*string) *MarkOssV2ResultResponseBody {
	s.SuccessRequestIds = v
	return s
}

func (s *MarkOssV2ResultResponseBody) Validate() error {
	return dara.Validate(s)
}
