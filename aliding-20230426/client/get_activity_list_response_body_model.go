// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetActivityListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetActivityListResponseBody
	GetRequestId() *string
	SetResult(v []*GetActivityListResponseBodyResult) *GetActivityListResponseBody
	GetResult() []*GetActivityListResponseBodyResult
	SetVendorRequestId(v string) *GetActivityListResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *GetActivityListResponseBody
	GetVendorType() *string
}

type GetActivityListResponseBody struct {
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string                              `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*GetActivityListResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s GetActivityListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetActivityListResponseBody) GoString() string {
	return s.String()
}

func (s *GetActivityListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetActivityListResponseBody) GetResult() []*GetActivityListResponseBodyResult {
	return s.Result
}

func (s *GetActivityListResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *GetActivityListResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *GetActivityListResponseBody) SetRequestId(v string) *GetActivityListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetActivityListResponseBody) SetResult(v []*GetActivityListResponseBodyResult) *GetActivityListResponseBody {
	s.Result = v
	return s
}

func (s *GetActivityListResponseBody) SetVendorRequestId(v string) *GetActivityListResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *GetActivityListResponseBody) SetVendorType(v string) *GetActivityListResponseBody {
	s.VendorType = &v
	return s
}

func (s *GetActivityListResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetActivityListResponseBodyResult struct {
	// example:
	//
	// 0q8gsudxxx
	ActivityId *string `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	// example:
	//
	// activity123
	ActivityName *string `json:"ActivityName,omitempty" xml:"ActivityName,omitempty"`
	// example:
	//
	// activity123
	ActivityNameInEnglish *string `json:"ActivityNameInEnglish,omitempty" xml:"ActivityNameInEnglish,omitempty"`
}

func (s GetActivityListResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetActivityListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetActivityListResponseBodyResult) GetActivityId() *string {
	return s.ActivityId
}

func (s *GetActivityListResponseBodyResult) GetActivityName() *string {
	return s.ActivityName
}

func (s *GetActivityListResponseBodyResult) GetActivityNameInEnglish() *string {
	return s.ActivityNameInEnglish
}

func (s *GetActivityListResponseBodyResult) SetActivityId(v string) *GetActivityListResponseBodyResult {
	s.ActivityId = &v
	return s
}

func (s *GetActivityListResponseBodyResult) SetActivityName(v string) *GetActivityListResponseBodyResult {
	s.ActivityName = &v
	return s
}

func (s *GetActivityListResponseBodyResult) SetActivityNameInEnglish(v string) *GetActivityListResponseBodyResult {
	s.ActivityNameInEnglish = &v
	return s
}

func (s *GetActivityListResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
