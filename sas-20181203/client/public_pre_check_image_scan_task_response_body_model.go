// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublicPreCheckImageScanTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *PublicPreCheckImageScanTaskResponseBodyData) *PublicPreCheckImageScanTaskResponseBody
	GetData() *PublicPreCheckImageScanTaskResponseBodyData
	SetRequestId(v string) *PublicPreCheckImageScanTaskResponseBody
	GetRequestId() *string
}

type PublicPreCheckImageScanTaskResponseBody struct {
	// The data returned if the call is successful.
	Data *PublicPreCheckImageScanTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// F9353221-40F4-5F98-B73C-2803DC804033
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PublicPreCheckImageScanTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PublicPreCheckImageScanTaskResponseBody) GoString() string {
	return s.String()
}

func (s *PublicPreCheckImageScanTaskResponseBody) GetData() *PublicPreCheckImageScanTaskResponseBodyData {
	return s.Data
}

func (s *PublicPreCheckImageScanTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PublicPreCheckImageScanTaskResponseBody) SetData(v *PublicPreCheckImageScanTaskResponseBodyData) *PublicPreCheckImageScanTaskResponseBody {
	s.Data = v
	return s
}

func (s *PublicPreCheckImageScanTaskResponseBody) SetRequestId(v string) *PublicPreCheckImageScanTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublicPreCheckImageScanTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PublicPreCheckImageScanTaskResponseBodyData struct {
	// The number of images to scan in the task.
	//
	// example:
	//
	// 6
	NeedAuthCount *int32 `json:"NeedAuthCount,omitempty" xml:"NeedAuthCount,omitempty"`
	// The quota for container image scan to be consumed by the task.
	//
	// example:
	//
	// 3
	ScanImageCount *int32 `json:"ScanImageCount,omitempty" xml:"ScanImageCount,omitempty"`
}

func (s PublicPreCheckImageScanTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s PublicPreCheckImageScanTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *PublicPreCheckImageScanTaskResponseBodyData) GetNeedAuthCount() *int32 {
	return s.NeedAuthCount
}

func (s *PublicPreCheckImageScanTaskResponseBodyData) GetScanImageCount() *int32 {
	return s.ScanImageCount
}

func (s *PublicPreCheckImageScanTaskResponseBodyData) SetNeedAuthCount(v int32) *PublicPreCheckImageScanTaskResponseBodyData {
	s.NeedAuthCount = &v
	return s
}

func (s *PublicPreCheckImageScanTaskResponseBodyData) SetScanImageCount(v int32) *PublicPreCheckImageScanTaskResponseBodyData {
	s.ScanImageCount = &v
	return s
}

func (s *PublicPreCheckImageScanTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
