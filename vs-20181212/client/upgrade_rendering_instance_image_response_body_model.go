// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeRenderingInstanceImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailedCount(v int64) *UpgradeRenderingInstanceImageResponseBody
	GetFailedCount() *int64
	SetFailedItems(v []*UpgradeRenderingInstanceImageResponseBodyFailedItems) *UpgradeRenderingInstanceImageResponseBody
	GetFailedItems() []*UpgradeRenderingInstanceImageResponseBodyFailedItems
	SetRequestId(v string) *UpgradeRenderingInstanceImageResponseBody
	GetRequestId() *string
	SetSuccessCount(v int64) *UpgradeRenderingInstanceImageResponseBody
	GetSuccessCount() *int64
	SetSuccessItems(v []*UpgradeRenderingInstanceImageResponseBodySuccessItems) *UpgradeRenderingInstanceImageResponseBody
	GetSuccessItems() []*UpgradeRenderingInstanceImageResponseBodySuccessItems
}

type UpgradeRenderingInstanceImageResponseBody struct {
	// The number of failed instances.
	//
	// example:
	//
	// 1
	FailedCount *int64 `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	// The information about failed instances.
	FailedItems []*UpgradeRenderingInstanceImageResponseBodyFailedItems `json:"FailedItems,omitempty" xml:"FailedItems,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of successful instances.
	//
	// example:
	//
	// 1
	SuccessCount *int64 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	// The information about successful instances.
	SuccessItems []*UpgradeRenderingInstanceImageResponseBodySuccessItems `json:"SuccessItems,omitempty" xml:"SuccessItems,omitempty" type:"Repeated"`
}

func (s UpgradeRenderingInstanceImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeRenderingInstanceImageResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeRenderingInstanceImageResponseBody) GetFailedCount() *int64 {
	return s.FailedCount
}

func (s *UpgradeRenderingInstanceImageResponseBody) GetFailedItems() []*UpgradeRenderingInstanceImageResponseBodyFailedItems {
	return s.FailedItems
}

func (s *UpgradeRenderingInstanceImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeRenderingInstanceImageResponseBody) GetSuccessCount() *int64 {
	return s.SuccessCount
}

func (s *UpgradeRenderingInstanceImageResponseBody) GetSuccessItems() []*UpgradeRenderingInstanceImageResponseBodySuccessItems {
	return s.SuccessItems
}

func (s *UpgradeRenderingInstanceImageResponseBody) SetFailedCount(v int64) *UpgradeRenderingInstanceImageResponseBody {
	s.FailedCount = &v
	return s
}

func (s *UpgradeRenderingInstanceImageResponseBody) SetFailedItems(v []*UpgradeRenderingInstanceImageResponseBodyFailedItems) *UpgradeRenderingInstanceImageResponseBody {
	s.FailedItems = v
	return s
}

func (s *UpgradeRenderingInstanceImageResponseBody) SetRequestId(v string) *UpgradeRenderingInstanceImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeRenderingInstanceImageResponseBody) SetSuccessCount(v int64) *UpgradeRenderingInstanceImageResponseBody {
	s.SuccessCount = &v
	return s
}

func (s *UpgradeRenderingInstanceImageResponseBody) SetSuccessItems(v []*UpgradeRenderingInstanceImageResponseBodySuccessItems) *UpgradeRenderingInstanceImageResponseBody {
	s.SuccessItems = v
	return s
}

func (s *UpgradeRenderingInstanceImageResponseBody) Validate() error {
	if s.FailedItems != nil {
		for _, item := range s.FailedItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SuccessItems != nil {
		for _, item := range s.SuccessItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpgradeRenderingInstanceImageResponseBodyFailedItems struct {
	// The error code of the failure.
	//
	// example:
	//
	// 200302
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message of the failure.
	//
	// example:
	//
	// Not Applied
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The cloud application service instance ID.
	//
	// example:
	//
	// render-072da95539d3402da90353b244191722
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
}

func (s UpgradeRenderingInstanceImageResponseBodyFailedItems) String() string {
	return dara.Prettify(s)
}

func (s UpgradeRenderingInstanceImageResponseBodyFailedItems) GoString() string {
	return s.String()
}

func (s *UpgradeRenderingInstanceImageResponseBodyFailedItems) GetErrCode() *string {
	return s.ErrCode
}

func (s *UpgradeRenderingInstanceImageResponseBodyFailedItems) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *UpgradeRenderingInstanceImageResponseBodyFailedItems) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *UpgradeRenderingInstanceImageResponseBodyFailedItems) SetErrCode(v string) *UpgradeRenderingInstanceImageResponseBodyFailedItems {
	s.ErrCode = &v
	return s
}

func (s *UpgradeRenderingInstanceImageResponseBodyFailedItems) SetErrMessage(v string) *UpgradeRenderingInstanceImageResponseBodyFailedItems {
	s.ErrMessage = &v
	return s
}

func (s *UpgradeRenderingInstanceImageResponseBodyFailedItems) SetRenderingInstanceId(v string) *UpgradeRenderingInstanceImageResponseBodyFailedItems {
	s.RenderingInstanceId = &v
	return s
}

func (s *UpgradeRenderingInstanceImageResponseBodyFailedItems) Validate() error {
	return dara.Validate(s)
}

type UpgradeRenderingInstanceImageResponseBodySuccessItems struct {
	// The cloud application service instance ID.
	//
	// example:
	//
	// render-1ada8cd82783407b99fa202826fc6447
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
}

func (s UpgradeRenderingInstanceImageResponseBodySuccessItems) String() string {
	return dara.Prettify(s)
}

func (s UpgradeRenderingInstanceImageResponseBodySuccessItems) GoString() string {
	return s.String()
}

func (s *UpgradeRenderingInstanceImageResponseBodySuccessItems) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *UpgradeRenderingInstanceImageResponseBodySuccessItems) SetRenderingInstanceId(v string) *UpgradeRenderingInstanceImageResponseBodySuccessItems {
	s.RenderingInstanceId = &v
	return s
}

func (s *UpgradeRenderingInstanceImageResponseBodySuccessItems) Validate() error {
	return dara.Validate(s)
}
