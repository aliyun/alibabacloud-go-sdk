// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebuildDesktopsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRebuildResults(v []*RebuildDesktopsResponseBodyRebuildResults) *RebuildDesktopsResponseBody
	GetRebuildResults() []*RebuildDesktopsResponseBodyRebuildResults
	SetRequestId(v string) *RebuildDesktopsResponseBody
	GetRequestId() *string
}

type RebuildDesktopsResponseBody struct {
	// The recreation results.
	RebuildResults []*RebuildDesktopsResponseBodyRebuildResults `json:"RebuildResults,omitempty" xml:"RebuildResults,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RebuildDesktopsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RebuildDesktopsResponseBody) GoString() string {
	return s.String()
}

func (s *RebuildDesktopsResponseBody) GetRebuildResults() []*RebuildDesktopsResponseBodyRebuildResults {
	return s.RebuildResults
}

func (s *RebuildDesktopsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RebuildDesktopsResponseBody) SetRebuildResults(v []*RebuildDesktopsResponseBodyRebuildResults) *RebuildDesktopsResponseBody {
	s.RebuildResults = v
	return s
}

func (s *RebuildDesktopsResponseBody) SetRequestId(v string) *RebuildDesktopsResponseBody {
	s.RequestId = &v
	return s
}

func (s *RebuildDesktopsResponseBody) Validate() error {
	if s.RebuildResults != nil {
		for _, item := range s.RebuildResults {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RebuildDesktopsResponseBodyRebuildResults struct {
	// The recreation result code. If the request was successful, `success` is returned. If the request failed, an error message is returned.
	//
	// example:
	//
	// IncorrectDesktopStatus
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The IDs of the cloud computers.
	//
	// example:
	//
	// ecd-gx2x1dhsmucyy****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The error message. This parameter is invalid if the value of `Code` is `success`.
	//
	// example:
	//
	// The current status of the desktop does not support this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RebuildDesktopsResponseBodyRebuildResults) String() string {
	return dara.Prettify(s)
}

func (s RebuildDesktopsResponseBodyRebuildResults) GoString() string {
	return s.String()
}

func (s *RebuildDesktopsResponseBodyRebuildResults) GetCode() *string {
	return s.Code
}

func (s *RebuildDesktopsResponseBodyRebuildResults) GetDesktopId() *string {
	return s.DesktopId
}

func (s *RebuildDesktopsResponseBodyRebuildResults) GetMessage() *string {
	return s.Message
}

func (s *RebuildDesktopsResponseBodyRebuildResults) SetCode(v string) *RebuildDesktopsResponseBodyRebuildResults {
	s.Code = &v
	return s
}

func (s *RebuildDesktopsResponseBodyRebuildResults) SetDesktopId(v string) *RebuildDesktopsResponseBodyRebuildResults {
	s.DesktopId = &v
	return s
}

func (s *RebuildDesktopsResponseBodyRebuildResults) SetMessage(v string) *RebuildDesktopsResponseBodyRebuildResults {
	s.Message = &v
	return s
}

func (s *RebuildDesktopsResponseBodyRebuildResults) Validate() error {
	return dara.Validate(s)
}
